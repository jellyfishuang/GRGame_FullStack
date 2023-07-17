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
export default class stackZone extends cc.Component {
    chowList: number[] = [];
    pongList: number[] = [];
    kongList: number[] = [];
    concealKongList: number[] = [];
    units: Unit[] = [];
    unitsIdList: number[] = [];

    start () {
    }

    addChow(cardIds: number[]){
        console.log("addChow");
        this.chowList.push(...cardIds);
        CardUtil.sortCardIdList(this.chowList);
        this.refreshStackZone();
    }

    addPong(cardId : number){
        console.log("addPong");
        this.pongList.push(cardId);
        CardUtil.sortCardIdList(this.pongList);
        this.refreshStackZone();
    }

    removePong(cardId : number){
        const index = this.pongList.indexOf(cardId, 0);
        this.pongList.splice(index, 1);
        CardUtil.sortCardIdList(this.pongList);
        this.refreshStackZone();
    }

    addKong(cardId : number){
        console.log("addKong");
        this.kongList.push(cardId);
        CardUtil.sortCardIdList(this.kongList);
        this.refreshStackZone();
    }

    addConcealKong(cardId : number){
        console.log("addConcealKong");
        this.concealKongList.push(cardId);
        CardUtil.sortCardIdList(this.concealKongList);
        this.refreshStackZone();
    }

    refreshStackZone(){
        this.node.removeAllChildren();
        let playerId = this.node.parent.parent.getComponent(Player).playerIndex;
        let prefab = this.node.parent.parent.getComponent(Player).prefab;
        this.createChowCards(prefab, playerId);
        this.createPongCards(prefab, playerId);
        this.createKongCards(prefab, playerId);
    }

    createChowCards(prefab:cc.Prefab, playerId:number){
        for(let i = 0; i < this.chowList.length; i++){
            var newUnit = CardUtil.newCard(prefab, this.chowList[i], i, playerId, true);
            newUnit.getComponent(Unit).jump();
            this.node.addChild(newUnit);
        }
    }

    createPongCards(prefab:cc.Prefab, playerId:number){
        let idx = 0;
        for(let i = 0; i < this.pongList.length; i++){
            for (let j = 0; j < 3; j++) {
                var newUnit = CardUtil.newCard(prefab, this.pongList[i], idx, playerId, true);
                newUnit.getComponent(Unit).jump();
                this.node.addChild(newUnit);
                idx++;
            }
        }
    }
    
    createKongCards(prefab:cc.Prefab, playerId:number){
        let idx = 0;
        for(let i = 0; i < this.kongList.length; i++){
            for (let j = 0; j < 3; j++) {
                var newUnit = CardUtil.newCard(prefab, this.kongList[i], idx, playerId, true);
                newUnit.getComponent(Unit).jump();
                if(j == 1){
                    newUnit.getComponent(Unit).stack();
                }
                this.node.addChild(newUnit);
                idx++;
            }
        }
    }
}
