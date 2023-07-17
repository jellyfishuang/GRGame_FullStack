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
export default class NewClass extends cc.Component {
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
        let playerId = this.node.parent.getComponent(Player).playerIndex;
        let prefab = this.node.parent.getComponent(Player).prefab;

        for (let i = 0; i < this.unitsIdList.length; i++) {
            var newUnit = CardUtil.newCard(prefab, this.unitsIdList[i], i, playerId, false);
            this.node.addChild(newUnit);
        }
        this.units = this.getComponentsInChildren(Unit);
        CardUtil.sortCards(this.unitsIdList, this.units);
        this.node.children.reverse();
    }

    addToFlower(cardId:number){
        this.unitsIdList.push(cardId);
        this.refreshCards();
    }
}
