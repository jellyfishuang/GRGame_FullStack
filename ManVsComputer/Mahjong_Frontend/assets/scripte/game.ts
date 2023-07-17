// Learn TypeScript:
//  - https://docs.cocos.com/creator/manual/en/scripting/typescript.html
// Learn Attribute:
//  - https://docs.cocos.com/creator/manual/en/scripting/reference/attributes.html
// Learn life-cycle callbacks:
//  - https://docs.cocos.com/creator/manual/en/scripting/life-cycle-callbacks.html
import Effect from "./effect"
import { EffectId } from "./effect"
import InteractionMenu from "./InteractionMenu"
import SelectChowMenu from "./SelectChowMenu"
import SelectKongMenu from "./SelectKongMenu"
import SelectWithoutMenu from "./SelectWithoutMenu"
import ChangeThreeTileMenu from "./ChangeThreeTileMenu"
import WithoutMenu from "./WithoutMenu"
import GameOverMenu from  "./GameOverMenu"
import SettlementMenu from "./SettlementMenu"
import DataResolver from "./SocketDataHandler/ServerDataResolver"
import DataCreator from "./SocketDataHandler/ClientDataCreator"
import TurnDirectionMenu from "./TurnDirectionMenu"
import Player from "./Player"
import { GameState, ActionId, ServerData, PlayerData, GameRule, DoActionData} from "./SocketDataHandler/DataSchema";
import { UiUtil, AppUtil} from "./AppUtil";

const {ccclass, property} = cc._decorator;

@ccclass
export default class NewClass extends cc.Component {

    @property(cc.SpriteAtlas)
    imgAtlas: cc.SpriteAtlas = null;

    @property(cc.Node)
    openGameButton: cc.Node=null;

    /**
     * user this when index is local, 0 is always real player
     */
    @property(Player)
    players: Player[] = [];

    @property(cc.Node)
    preGameMenu: cc.Node = null;

    @property(InteractionMenu)
    interactionMenu: InteractionMenu = null;

    @property(SelectChowMenu)
    selectChowMenu: SelectChowMenu = null;

    @property(SelectKongMenu)
    selectKongMenu: SelectKongMenu = null;

    @property(SelectWithoutMenu)
    selectWithoutMenu: SelectWithoutMenu = null;

    @property(ChangeThreeTileMenu)
    changeThreeTileMenu: ChangeThreeTileMenu = null;

    @property(WithoutMenu)
    withoutMenu: WithoutMenu = null;

    @property(GameOverMenu)
    gameOverMenu: GameOverMenu = null;

    @property(cc.Sprite)
    gameDirection: cc.Sprite=null;

    @property(TurnDirectionMenu)
    turnDirectionMenu: TurnDirectionMenu = null;

    @property(cc.Label)
    cardCount: cc.Label=null;

    @property(SettlementMenu)
    settlementMenu: SettlementMenu = null;

    @property(Effect)
    effectUI: Effect=null;

    soundVolume : number = 0.5;

    webSocket : WebSocket = null;

    /**
     * Set global for ui to send message back to server
     */
    userPlayerData: PlayerData = null;
    userplayerServerIndex: number = -1;
    /**
     * The mapping of server player index and player, use this when index is from server
     */
    playerMap: Map<number, Player> = new Map();

    isGameOver: boolean = false;

    isSelfHu: boolean = false;

    g_playerNow: Player= null;
    
    start() {
    }
    
    startGame(){
        this.webSocket.send(DataCreator.createDealTilesData());
        this.openGameButton.active = false;
        this.preGameMenu.active = true;
    }

    onLoad () {
        let webSocket = window["webSocket"];
        this.webSocket = webSocket;
        this.players[0].webSocket = webSocket;
        this.configWebSocket();
        this.setPlayerId();
        cc.debug.setDisplayStats(false);
    }

    setPlayerId(){
        this.players[0].playerId = window["playerId"];
        for(let i = 1 ; i < 4 ; i++){
            this.players[i].playerId = "Bot" + i;
        }
    }

    configWebSocket(){
        this.webSocket.onmessage = (event) => {
            var data = JSON.parse(event.data);
            let serverData = DataResolver.getServerData(data);
            console.log("server data = ", serverData);
            if(serverData.isGameInit()){
                this.initGame(serverData);
                return;
            }
            
            if(serverData.isPreGameStage()){
                this.preGameStage(serverData);
                return;
            }
            this.isGameOver = serverData.isGameOver();

            this.demoPlays(serverData);

            this.updateGameBoard(serverData);

            if(this.isGameOver){
                this.gameOverMenu.show();
                return;
            }

            this.userPlayerPhase(serverData);
        }
    }

    initGame(serverData: ServerData){
        this.initGameUi(serverData.GameRule);
        this.initPlayerMap(serverData.RealPlayer);
        this.setGameDirection(serverData.PrevailingWind);
        this.setPlayerHands(serverData.Players);
    }

    initGameUi(gameRule: GameRule){
        this.interactionMenu.initButtons(UiUtil.initAvailableInteractions(gameRule));
    }

    initPlayerMap(userPlayerIndex: number){
        console.log("playerMap initail");
        for (let i = 0; i < 4; i++) {
            let serverIndex = this.getRealPlayerIndex(userPlayerIndex, i);
            console.log("serverIndex = %d , clientIndex = %d", serverIndex, i);
            let player = this.players[i];
            player.playerServerIndex = serverIndex;
            this.playerMap.set(serverIndex, player);
        }
        this.userplayerServerIndex = userPlayerIndex;
    }

    getRealPlayerIndex(realUserPlayerIndex: number, frontEndPlayerIndex: number): number{
        let boundary: number = 4;
        let a: number = frontEndPlayerIndex + realUserPlayerIndex;
        if(boundary <= a) return a - boundary;
        return a;
    }

    setGameDirection(dealerWind: number){
        this.gameDirection.spriteFrame = this.getWindSprite(dealerWind);
    }

    getWindSprite(dealerWind: number){
        if(dealerWind == 27) return this.imgAtlas.getSpriteFrame("MahjongVal_101_@2x");// 東
        if(dealerWind == 28) return this.imgAtlas.getSpriteFrame("MahjongVal_103_@2x");// 南
        if(dealerWind == 29) return this.imgAtlas.getSpriteFrame("MahjongVal_105_@2x");// 西
        if(dealerWind == 30) return this.imgAtlas.getSpriteFrame("MahjongVal_107_@2x");// 北
        return null;
    }

    setPlayerHands(players: PlayerData[]){
        for (let i = 0; i < players.length; i++) {
            this.updateHand(i, players[i]);
        }
    }

    updateHand(playerIndex: number, playerData: PlayerData){
        let player = this.playerMap.get(playerIndex);
        player.setCardZone(playerData.HandTiles);
        player.drawZone.remove();
    }

    preGameStage(serverData: ServerData){
        switch(serverData.GameState){
            case GameState.StateWithoutTiles : this.updateWithouts(serverData.Players); break;
            case GameState.StateChangeTiles : this.updateChangeThreeTiles(serverData.Players); break;
        }
    }

    demoPlays(serverData: ServerData){
        console.log("demoPlays");
        let playerIndex = serverData.PlayerNow;
        let playerNowData = serverData.Players[playerIndex];
        let playerNow = this.playerMap.get(playerIndex);
        this.g_playerNow = playerNow;

        this.isSelfHu = serverData.DoAction.Action == ActionId[ActionId.SelfHu];
        let isHu = serverData.DoAction.Action == ActionId[ActionId.Hu] || serverData.DoAction.Action == ActionId[ActionId.SelfHu];
        this.updateDrawTile(playerNow, playerNowData.OnDrawTile, isHu);
        this.updateFlowers(playerNow, playerNowData.FlowerTiles);
        switch(serverData.GameState){
            case GameState.StateDiscard : this.discard(playerNowData, playerNow, serverData.OntableTile); break;
            case GameState.StateAction : this.demoActions(serverData); break;
        }
    }

    async updateDrawTile(player: Player, drawTileId: number, isHu: boolean){
        if(drawTileId == -1) return;
        await AppUtil.sleep(500);
        this.draw(player, drawTileId);
        await AppUtil.sleep(500);
        if(!isHu && player.playerIndex != 0){
            this.webSocket.send(DataCreator.createAIDiscardTileData(this.players[0].playerId, player.playerServerIndex));
        }
    }

    draw(player: Player, card: number){
        console.log("Player: %d draw: %d", player.playerIndex, card);
        player.draw(card);
    }

    updateFlowers(player: Player, flowers: number[]){
        if(flowers == null || flowers.length == 0) {
            return;
        }
        player.setFlowers(flowers);
    }

    updateWithouts(playerDatas: PlayerData[]){
        for(let i = 0 ; i < 4 ; i++){
            let player  = this.playerMap.get(i);
            let playerData = playerDatas[i];
            this.selectWithout(player, playerData.WithoutTile);
        }
    }

    updateChangeThreeTiles(playerDatas: PlayerData[]){
        for(let i = 0 ; i < 4 ; i++){
            let player  = this.playerMap.get(i);
            let playerData = playerDatas[i];
            this.changeThreeTiles(playerData, player);
        }
    }

    discard(playerData: PlayerData, player: Player, card: number){
        console.log("Player: %d discard: %d", player.playerIndex, card);
        for(let i = 0 ; i < 4 ; i++){
            this.players[i].dropZone.unHighlightLastCard();
        }
        player.refreshCardZone(playerData.DiscardTiles, playerData.HandTiles);
    }

    demoActions(serverData: ServerData){
        let doActionData = serverData.DoAction;
        let doActionPlayer = this.playerMap.get(doActionData.Player);
        let doActionPlayerData = serverData.Players[doActionData.Player];

        let previousPlayerIndex = doActionData.PriviousPlayer;
        let previousPlayerData = serverData.Players[previousPlayerIndex];
        let previousPlayer = this.playerMap.get(previousPlayerIndex);
        if(doActionData.Tiles == null) return;

        this.updatePreviousPlayer(previousPlayerData, previousPlayer);
        console.log("Player: %d demoAction: %s", doActionPlayer.playerIndex, doActionData.Action);
        switch(doActionData.Action){
            case ActionId[ActionId.Buhua] : this.buhua(); break;
            case ActionId[ActionId.Chow] : this.chow(doActionPlayerData, doActionPlayer, doActionData.Tiles); break;
            case ActionId[ActionId.Pong] : this.pong(doActionPlayerData, doActionPlayer, doActionData.Tiles[0]); break;
            case ActionId[ActionId.ExposedKong] : this.kong(doActionPlayerData, doActionPlayer, doActionData.Tiles[0]); break;
            case ActionId[ActionId.ConcealedKong] : this.concealedKong(doActionPlayerData, doActionPlayer, doActionData.Tiles[0]); break;
            case ActionId[ActionId.AddKong] : this.addKong(doActionPlayerData, doActionPlayer, doActionData.Tiles[0]); break;
            case ActionId[ActionId.Hu] : this.hu(this.players,serverData.Players); break;
            case ActionId[ActionId.SelfHu] : this.selfDrawn(this.players,serverData.Players); break;
        }
    }

    updatePreviousPlayer(playerData: PlayerData, player: Player){
        player.dropZone.setCards(playerData.DiscardTiles, false);
    }

    buhua(){

    }

    changeThreeTiles(playerData: PlayerData, player: Player){
        player.setCardZone(playerData.HandTiles);
    }

    selectWithout(player: Player, withoutId: number){
        this.withoutMenu.updateMenu(player.playerIndex, withoutId);
    }

    chow(playerData: PlayerData, player: Player, cards: number[]){
        this.effectUI.playEffect(player.playerIndex, EffectId.Chi);
        player.chow(cards, playerData.HandTiles);
    }

    pong(playerData: PlayerData, player: Player, card: number){
        this.effectUI.playEffect(player.playerIndex, EffectId.Peng);
        player.pong(card, playerData.HandTiles);
    }

    kong(playerData: PlayerData, player: Player, card: number){
        this.effectUI.playEffect(player.playerIndex, EffectId.Gang);
        player.kong(card, playerData.HandTiles);
    }

    concealedKong(playerData: PlayerData, player: Player, card: number){
        this.effectUI.playEffect(player.playerIndex, EffectId.Gang);
        player.kong(card, playerData.HandTiles);
    }

    addKong(playerData: PlayerData, player: Player, card: number){
        this.effectUI.playEffect(player.playerIndex, EffectId.Gang);
        player.addKong(card, playerData.HandTiles);
    }

    hu(players: Player[], playerDatas: PlayerData[]){
        this.effectUI.playEffect(0, EffectId.Hu);
        this.settlementMenu.settles(players, playerDatas);
        this.settlementMenu.show();
    }

    selfDrawn(players: Player[], playerDatas: PlayerData[]){
        this.effectUI.playEffect(0, EffectId.Zm);
        this.settlementMenu.settles(players, playerDatas);
        this.settlementMenu.show();
    }

    ting(playerIndex: number){
        this.effectUI.playEffect(playerIndex, EffectId.Ting);
    }

    ming(playerIndex: number){
        this.effectUI.playEffect(playerIndex, EffectId.Ming);
    }

    crit(playerIndex: number){
        this.effectUI.playEffect(playerIndex, EffectId.Crit);
    }

    updateGameBoard(serverData: ServerData){
        this.updateCardCount(serverData.TilesSea.length);
        this.updateturnDirection(this.playerMap.get(serverData.PlayerNow).playerIndex);
        this.updatePoint(serverData.Players);
    }

    updateCardCount(tileSeaCount: number){
        this.cardCount.string = String(tileSeaCount);
    }

    updateturnDirection(playerIndex: number){
        this.turnDirectionMenu.turnByIndex(playerIndex);
    }

    updatePoint(playerDatas: PlayerData[]){
        for(let i = 0 ; i < playerDatas.length ; i++){
            let player = this.playerMap.get(i);
            let playerData = playerDatas[i];
            player.point = playerData.Point;
        }
    }

    userPlayerPhase(serverData: ServerData){
        let isUserPlayerTurn = serverData.isUserPlayerThrow();
        this.updateUserLock(isUserPlayerTurn);
        this.userPlayerData = serverData.Players[this.userplayerServerIndex];
        this.showInteractionMenu(this.getInteractions(serverData), isUserPlayerTurn, serverData.isDiscardTile());
    }

    showInteractionMenu(availableActions: number[], isUserPlayerTurn: boolean, isSomeoneDiscardTile: boolean){
        if (availableActions.length == 0 && !isUserPlayerTurn && isSomeoneDiscardTile){
            this.webSocket.send(DataCreator.createNothingData(this.userplayerServerIndex));
            return;
        }
        console.log("showInteractions: ", availableActions);
        this.interactionMenu.showByIndex(availableActions, this.isCancelable(availableActions, isUserPlayerTurn));
    }

    isCancelable(buttoneIds: number[], isUserPlayerTurn: boolean){
        return !isUserPlayerTurn && this.isInteractionCancelable(buttoneIds);
    }

    isInteractionCancelable(buttoneIds: number[]){
        return !buttoneIds.includes(ActionId.Buhua) || 
        !buttoneIds.includes(ActionId.SelectWithout);
    }

    getInteractions(serverData: ServerData): number[]{
        let actions = []
        let gameState = serverData.GameState;
        let userPlayerData = serverData.Players[serverData.RealPlayer];
        if(gameState == GameState.StateBuHua) actions.push(ActionId.Buhua);
        if(userPlayerData.CanChow) actions.push(ActionId.Chow);
        if(userPlayerData.CanPon) actions.push(ActionId.Pong);
        if(userPlayerData.CanKong) actions.push(ActionId.ExposedKong);
        if(userPlayerData.CanConcealedKong) actions.push(ActionId.ConcealedKong);
        if(userPlayerData.CanAddKong) actions.push(ActionId.AddKong);
        if(userPlayerData.CanHu) actions.push(ActionId.Hu);
        //if(userPlayerData.) actions.push(ActionId.Ting);
        if(userPlayerData.CanSelfHu) actions.push(ActionId.SelfHu);
        //if(userPlayerData.) actions.push(ActionId.Crit);
        //if(userPlayerData.) actions.push(ActionId.Ming);
        return actions;
    }

    updateUserLock(isUserPlayerTurn: boolean){
        if(isUserPlayerTurn){
            this.players[0].wakeUp();
        } else{
            this.players[0].sleep();
        }
    }

    createWithoutList(...indexs: number[]){
        return indexs;
    }

    volumeSet(slider:cc.Slider){
        this.soundVolume = slider.progress;
    }

    preGameButtonOpenGame(){
        this.webSocket.send(DataCreator.createOpenGameData());
        this.preGameMenu.active = false;
    }

    preGameButtonSelectWithout(){
        this.preGameMenu.active = false;
        this.selectWithoutMenu.show();
    }
    
    userSelectWithout(event: Event, CustomEventData){
        this.webSocket.send(DataCreator.createSelectWithoutData(Number(CustomEventData), this.userplayerServerIndex));
        this.selectWithoutMenu.hide();
        this.preGameMenu.active = true;
    }

   preGameButtonChangeThreeTile(){
        this.players[0].pickMultipleCardLimit = 3;
        this.updateUserLock(true);
        this.preGameMenu.active = false;
        this.changeThreeTileMenu.show();
    }

    decideChangeThreeTile(){
        this.updateUserLock(false);
        let tiles = this.players[0].getMultiPickedCards();
        this.webSocket.send(DataCreator.createChangeThreeTilesData(tiles, this.playerMap.entries().next().value[0]));
        this.changeThreeTileMenu.hide();
        this.preGameMenu.active = true;
    }
    
    interactionButtonCancel(){
        this.webSocket.send(DataCreator.createNothingData(this.userplayerServerIndex));
        this.interactionMenu.hide();
    }

    interactionButtonChow(){
        this.interactionMenu.hide();
        this.selectChowMenu.showChowOptions(this.userPlayerData.CanChowSet);
    }

    chowOption(event: Event, CustomEventData){
        this.webSocket.send(DataCreator.createChowData(this.userPlayerData.CanChowSet[Number(CustomEventData)], this.userplayerServerIndex));
        this.selectChowMenu.hide();
    }

    interactionButtonPong(){
        this.webSocket.send(DataCreator.createPongData(this.userPlayerData.CanPonSet, this.userplayerServerIndex));
        this.interactionMenu.hide();
    }

    interactionButtonKong(){
        this.interactionMenu.hide();
        this.selectKongMenu.showKongOptions(this.userPlayerData.CanKongSet, false);
    }

    kongOption(event: Event, CustomEventData){
        this.webSocket.send(DataCreator.createKongData(this.userPlayerData.CanKongSet[Number(CustomEventData)], this.userplayerServerIndex));
        this.selectKongMenu.hide();
    }

    interactionButtonConcealedKong(){
        this.interactionMenu.hide();
        this.selectKongMenu.showKongOptions(this.userPlayerData.CanConcealedKongSet, true);
    }

    concealedKongOption(event: Event, CustomEventData){
        this.webSocket.send(DataCreator.createConcealedKongData(this.userPlayerData.CanConcealedKongSet[Number(CustomEventData)], this.userplayerServerIndex));
        this.selectKongMenu.hide();
    }

    interactionButtonAddKong(){
        this.webSocket.send(DataCreator.createAddKongData(this.userPlayerData.CanAddKongSet, this.userplayerServerIndex));
        this.interactionMenu.hide();
    }

    interactionButtonTin(){
    }

    interactionButtonHu(){
        this.webSocket.send(DataCreator.createHuData(this.userPlayerData.CanHuTile, this.userplayerServerIndex));
        this.interactionMenu.hide();
    }

    interactionButtonSelfHu(){
        this.webSocket.send(DataCreator.createSelfHuData(this.userPlayerData.CanSelfHuTile, this.userplayerServerIndex));
        this.interactionMenu.hide();
    }

    interactionButtonCrit(){
    }

    interactionButtonMing(){
    }

    interactionButtonAddFlower(){
    }

    selectChiMenuButtonCancel(){
        this.selectChowMenu.hide();
        this.interactionMenu.show();
    }

    selectKongMenuButtonCancel(){
        this.selectKongMenu.hide();
        this.interactionMenu.show();
    }

    closeSettelmentLayer(){
        if(!this.isGameOver){
            if(this.g_playerNow.playerIndex != 0){
                //AI
                if(this.isSelfHu){
                    this.webSocket.send(DataCreator.createAIDiscardTileData(this.players[0].playerId, this.g_playerNow.playerServerIndex));
                    this.g_playerNow.throwCard();
                } else{
                    this.webSocket.send(DataCreator.createNothingData(this.userplayerServerIndex));
                }
            } else {
                // Real player
                if(this.isSelfHu){
                    this.g_playerNow.throwCard();
                }
                this.webSocket.send(DataCreator.createNothingData(this.userplayerServerIndex));
            }
        }
        this.settlementMenu.hide();
    }
    
    test(){
        this.players[0].setCardZone([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 39])
        this.players[0].draw(40)
        this.players[0].setFlowers([34, 35, 36, 37, 38, 39, 40, 41]);

        this.players[1].setCardZone([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 39])
        this.players[1].draw(40)
        this.players[1].setFlowers([34, 35, 36, 37, 38, 39, 40, 41]);

        this.players[2].setCardZone([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 39])
        this.players[2].draw(40)
        this.players[2].setFlowers([34, 35, 36, 37, 38, 39, 40, 41]);

        this.players[3].setCardZone([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 39])
        this.players[3].draw(40)
        this.players[3].setFlowers([34, 35, 36, 37, 38, 39, 40, 41]);
    }

    test3(){
        this.players[0].throwCard();
    }
}