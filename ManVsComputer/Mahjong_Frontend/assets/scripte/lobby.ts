// Learn TypeScript:
//  - https://docs.cocos.com/creator/manual/en/scripting/typescript.html
// Learn Attribute:
//  - https://docs.cocos.com/creator/manual/en/scripting/reference/attributes.html
// Learn life-cycle callbacks:
//  - https://docs.cocos.com/creator/manual/en/scripting/life-cycle-callbacks.html

import DataCreator from "./SocketDataHandler/ClientDataCreator"
import DataResolver from "./SocketDataHandler/ServerDataResolver"
import { AppUtil } from "./AppUtil";

const {ccclass, property} = cc._decorator;

@ccclass
export default class Lobby extends cc.Component {

    @property(cc.EditBox)
    idTextBox: cc.EditBox=null;

    @property(cc.EditBox)
    ipTextBox: cc.EditBox=null;

    @property(cc.EditBox)
    portTextBox: cc.EditBox=null;

    @property(cc.Label)
    connectInfoLabel: cc.Label=null;

    @property(cc.Label)
    roomIdLabel: cc.Label=null;

    @property(cc.Button)
    connectButton: cc.Button=null;

    @property(cc.Button)
    startButton: cc.Button=null;

    @property(cc.ToggleContainer)
    gameModToggles: cc.ToggleContainer=null;

    @property(cc.EditBox)
    aiTestRounds: cc.EditBox=null;

    webSocket : WebSocket = null;

    dataResolver : DataResolver = null;

    gameMode : number = 0;

    //defaultAddress : string = 'ws://localhost:1234';
    defaultAddress : string = 'ws://140.118.157.29:1234';

    getSocketAddress(ip:string, port:string):string{
        if(ip !== "" && port !== ""){
            return 'ws://' + ip + ':' + port;
        }
        return this.defaultAddress;
    }

    chooseGameMode(){
        let gameModes = this.gameModToggles.toggleItems;
        for(let i = 0 ; i < gameModes.length ; i++){
            if(gameModes[i].isChecked){
                this.gameMode = i;
                break;
            }
        }
        console.log("choose game mode : " + this.gameMode);
    }

    getGameModeData(){
        switch(this.gameMode){
            case 0: return "8Joker";
            case 1: return "46Joker";
            case 2: return "bloodBattle";
            case 3: return "TaiwanMJ";
        }
    }

    creteWebSocket(ip:string, port:string):WebSocket{
        let socket = new WebSocket(this.getSocketAddress(ip, port));
        socket.addEventListener("open",()=>{console.log("socket connected")})
        socket.onmessage = (event) => {
            let serverData = DataResolver.getServerData(event.data);
            console.log("server data = ", serverData);
            let roomId = serverData.RoomId;
            this.roomIdLabel.string = String(roomId);
            window["RoomID"] = roomId;
        }
        return socket;
    }

    async connectToServer(){
        this.webSocket = this.creteWebSocket(this.ipTextBox.string, this.portTextBox.string);
        this.connectInfoLabel.string = "Connecting...";
        let counter = 0;
        let tryLimit = 5;
        while(!this.isSocketConneted() && counter < tryLimit){
            await AppUtil.sleep(1000);
            counter++;
        }
        console.log("counter = ", counter);
        if(counter == tryLimit){
            this.connectInfoLabel.string = "Connection fail, please try again";
            return;
        }

        this.connectInfoLabel.string = "Connection success!";
        setTimeout(()=>{this.webSocket.send(DataCreator.createConnectionData())}, 100);
        setTimeout(()=>{this.webSocket.send(DataCreator.createCreateRoomData(this.getPlayerId(), this.getGameModeData()))}, 100);
    }

    isSocketConneted(): boolean{
        if(this.webSocket == null) return false;
        let status = this.webSocket.readyState;
        console.log("socket status = ", WebSocket[status]);
        return status == WebSocket.OPEN;
    }

    gameStart(){
        if(!this.isSocketConneted()){
            this.connectInfoLabel.string = "Not connected to server, please try connect first!";
            return;
        }
        
        let playerId = this.getPlayerId();
        window["playerId"] = playerId;
        window["webSocket"] = this.webSocket;
        cc.director.loadScene("game");
    }

    getPlayerId(){
        if(this.idTextBox.string !== ""){
            return this.idTextBox.string;
        }
        return "玩家";
    }

    aiTestButton(){
        this.webSocket.send(DataCreator.createAITestData(Number(this.aiTestRounds.string), this.getGameModeData()));
    }

    start () {
    }
}
