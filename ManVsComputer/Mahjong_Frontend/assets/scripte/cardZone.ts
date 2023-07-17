// Learn TypeScript:
//  - https://docs.cocos.com/creator/manual/en/scripting/typescript.html
// Learn Attribute:
//  - https://docs.cocos.com/creator/manual/en/scripting/reference/attributes.html
// Learn life-cycle callbacks:
//  - https://docs.cocos.com/creator/manual/en/scripting/life-cycle-callbacks.html
import Unit from "./unit";
import Player from "./Player";
import { CardUtil } from "./AppUtil";

const {ccclass, property} = cc._decorator;

@ccclass
export default class CardZone extends cc.Component {
    units: Unit[] = [];
    unitsIdList: number[] = [];

    start () {
    }

    setCards(cardIds : number[]) {
        this.unitsIdList = cardIds;
        this.refreshCards();
    }

    refreshCards(){
        this.node.removeAllChildren();
        let playerId = this.node.parent.parent.getComponent(Player).playerIndex;
        let prefab = this.node.parent.parent.getComponent(Player).prefab;

        for (let i = 0; i < this.unitsIdList.length; i++) {
            var newUnit = CardUtil.newCard(prefab, this.unitsIdList[i], i, playerId, false);
            this.node.addChild(newUnit);
        }
        this.units = this.getComponentsInChildren(Unit);
        CardUtil.sortCards(this.unitsIdList, this.units);
        this.node.children.reverse();
    }

    addToHand(cardId:number){
        this.unitsIdList.push(cardId);
        this.refreshCards();
    }

    throwFromHand(handIndex:number): number{
        let newCards = this.unitsIdList;
        let cardId = this.unitsIdList[handIndex];
        newCards.forEach( (item, idx) => {
            if(idx === handIndex) newCards.splice(idx,1);
          });
        this.refreshCards();
        return cardId;
    }

    searchCardIndex(cardId: number): number[]{
        let reult = [];
        for (let i = 0; i < this.unitsIdList.length; i++) {
            if(cardId == this.unitsIdList[i]){
                reult.push(i);
            }
        }
        return reult;
    }

    pick(index:number){
        this.units[index].jump();
    }

    unpick(index:number){
        this.units[index].down();
    }

    unPickAll(){
        for (let i = 0; i < this.units.length; i++) {
            this.units[i].down();
        }
    }
}
