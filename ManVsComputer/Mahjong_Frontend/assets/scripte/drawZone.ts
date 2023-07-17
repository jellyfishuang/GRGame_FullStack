// Learn TypeScript:
//  - https://docs.cocos.com/creator/manual/en/scripting/typescript.html
// Learn Attribute:
//  - https://docs.cocos.com/creator/manual/en/scripting/reference/attributes.html
// Learn life-cycle callbacks:
//  - https://docs.cocos.com/creator/manual/en/scripting/life-cycle-callbacks.html
import Unit from "./unit";

const {ccclass, property} = cc._decorator;

@ccclass
export default class drawZone extends cc.Component {
    drawUnit: Unit = null;
    
    start () {
        this.drawUnit = this.node.getChildByName("drawCard").getComponent(Unit);
    }

    drawNewCard(cardId: number){
        this.drawUnit.set(cardId);
    }

    remove(): number{
        let currentId = this.drawUnit.ID;
        this.drawNewCard(99);
        this.drawUnit.down();
        return currentId;
    }

    pick(){
        this.drawUnit.jump();
    }

    unpick(){
        this.drawUnit.down();
    }
}
