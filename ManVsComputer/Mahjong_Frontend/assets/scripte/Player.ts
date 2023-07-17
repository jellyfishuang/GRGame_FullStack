// Learn TypeScript:
//  - https://docs.cocos.com/creator/manual/en/scripting/typescript.html
// Learn Attribute:
//  - https://docs.cocos.com/creator/manual/en/scripting/reference/attributes.html
// Learn life-cycle callbacks:
//  - https://docs.cocos.com/creator/manual/en/scripting/life-cycle-callbacks.html

import cardZone from "./cardZone";
import drawZone from "./drawZone";
import stackZone from "./stackZone";
import dropZone from "./dropZone";
import flowerZone from "./flowerZone";
import DataCreator from "./SocketDataHandler/ClientDataCreator";
import game from "./game";

const {ccclass, property} = cc._decorator;

@ccclass
export default class Player extends cc.Component {
    DRAW_CARD_INDEX: number = 999;

    @property(cc.Prefab)
    prefab: cc.Prefab = null;

    @property(cardZone)
    cardZone: cardZone = null;

    @property(drawZone)
    drawZone: drawZone = null;

    @property(dropZone)
    dropZone: dropZone = null;

    @property(stackZone)
    stackZone: stackZone = null;

    @property(flowerZone)
    flowerZone: flowerZone = null;

    @property
    playerIndex: number = 0;

    @property
    isInteractable : boolean = false;

    point: number = 0;
    playerServerIndex: number = -1;
    playerId: string = null;
    pickUnitIndex: number = -1;
    without:number = 0;
    webSocket : WebSocket = null;
    pickMultipleCardLimit: number = -1;
    pickMultiCardIndex: number[] = [];


    setCardZone(cardIds : number[]){
        this.cardZone.setCards(cardIds);
    }

    setFlowers(cardIds : number[]){
        this.flowerZone.setCards(cardIds);
    }

    pick(handIndex:number){
        if(handIndex == -1){
            return;
        }

        if(this.isPickMultipleCards()){
            this.handleMutiPick(handIndex);
            return;
        }

        if(this.isCardPicked(handIndex)){
            this.handlePickedCard(handIndex);
            return;
        }

        if(handIndex == this.DRAW_CARD_INDEX){
            this.pickDrawCard();
            return;
        }
        this.pickCardZone(handIndex);
    }

    handleMutiPick(handIndex:number){
        // remove if exist, otherwise add to list
        const index = this.pickMultiCardIndex.indexOf(handIndex, 0);
        if (index > -1) {
            this.pickMultiCardIndex.splice(index, 1);
        } else{
            this.pickMultiCardIndex.push(handIndex)
        }

        // more than limitation then shift
        if(this.pickMultiCardIndex.length> this.pickMultipleCardLimit){
            this.pickMultiCardIndex.shift();
        }

        // reset picked card by pickMultiCardIndex
        this.cardZone.unPickAll();
        for(let i = 0; i < this.pickMultiCardIndex.length ; i++){
            this.cardZone.pick(this.pickMultiCardIndex[i]);
        }
        let currentGame = this.cardZone.node.parent.parent.parent.getComponent(game);
        currentGame.changeThreeTileMenu.decideButton.interactable = this.isReachMaxMultiCardLimit();
    }

    handlePickedCard(handIndex:number){
        this.cardZone.node.parent.parent.parent.getComponent(game).interactionMenu.hide();
        this.sleep();
        this.pickUnitIndex = -1;
        this.webSocket.send(DataCreator.createDiscardTileData(this.playerId, this.playerServerIndex, this.findPickedCardId(handIndex)));
        this.webSocket.send(DataCreator.createNothingData(this.playerServerIndex));
    }

    getMultiPickedCards(): number[]{
        let result : number[] = [];
        for(let i = 0; i < this.pickMultiCardIndex.length ; i++){
            result.push(this.cardZone.unitsIdList[this.pickMultiCardIndex[i]]);
        }
        return result;
    }

    pickDrawCard(){
        this.drawZone.pick();
        if(this.pickUnitIndex != this.DRAW_CARD_INDEX){
            this.cardZone.unPickAll();
        }
    }

    pickCardZone(handIndex:number){
        this.cardZone.unPickAll();
        this.cardZone.pick(handIndex);
        this.drawZone.unpick();
        this.pickUnitIndex = handIndex;
    }

    isCardPicked(handIndex:number): boolean{
        if(handIndex == this.DRAW_CARD_INDEX) return this.drawZone.drawUnit.isPicked;
        return this.cardZone.units[handIndex].isPicked;
    }

    findPickedCardId(handIndex:number){
        if(handIndex == this.DRAW_CARD_INDEX) return this.drawZone.drawUnit.ID;
        return this.cardZone.unitsIdList[handIndex];
    }

    isPickMultipleCards(): boolean{
        return this.pickMultipleCardLimit != -1;
    }

    isReachMaxMultiCardLimit(){
        return this.pickMultiCardIndex.length == this.pickMultipleCardLimit;
    }

    throwCard(){
        if(this.drawZone.drawUnit == null || this.drawZone.drawUnit.ID == 99){
            return;
        }
        let cardId = this.drawZone.remove();
        this.dropZone.addToDrop(cardId);
    }

    refreshCardZone(discardCardIds: number[], newHands: number[]){
        this.drawZone.remove(); // draw card will always be removed
        this.cardZone.setCards(newHands);
        this.dropZone.setCards(discardCardIds, true);
    }

    draw(cardId : number){
        this.drawZone.drawNewCard(cardId);
    }

    chow(cardIds: number[], newHands: number[]){
        this.stackZone.addChow(cardIds);
        this.cardZone.setCards(newHands);
    }

    pong(cardId : number, newHands: number[]){
        this.stackZone.addPong(cardId);
        this.cardZone.setCards(newHands);
    }

    kong(cardId : number, newHands: number[]){
        this.stackZone.addKong(cardId);
        this.cardZone.setCards(newHands);
    }

    addKong(cardId : number, newHands: number[]){
        this.stackZone.removePong(cardId);
        this.kong(cardId, newHands);
    }

    wakeUp(){
        this.isInteractable = true;
    }

    sleep(){
        this.pickUnitIndex = -1;
        this.pickMultipleCardLimit = -1;
        this.cardZone.unPickAll();
        this.isInteractable = false;
    }
}
