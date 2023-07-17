// Learn TypeScript:
//  - https://docs.cocos.com/creator/manual/en/scripting/typescript.html
// Learn Attribute:
//  - https://docs.cocos.com/creator/manual/en/scripting/reference/attributes.html
// Learn life-cycle callbacks:
//  - https://docs.cocos.com/creator/manual/en/scripting/life-cycle-callbacks.html

import SettlementZone from "./SettlementZone";
import Player from "./Player";
import { PlayerData } from "./SocketDataHandler/DataSchema";

const {ccclass, property} = cc._decorator;

@ccclass
export default class NewClass extends cc.Component {

    @property(SettlementZone)
    settlementZones: SettlementZone[] = [];

    settles(players: Player[], playerDatas: PlayerData[]){
        for(let i = 0; i < this.settlementZones.length ; i++){
            let player = players[i];
            let playerData = playerDatas[player.playerServerIndex];
            let fannStr = playerData.CanHuFaanListStr == null ? "" : playerData.CanHuFaanListStr.join(",");
            this.settlementZones[i].setContent(player.playerId, player.point.toString(), fannStr);
        }
    }

    show(){
        this.node.active = true;
    }

    hide(){
        this.node.active = false;
    }

    start () {
    }
}
