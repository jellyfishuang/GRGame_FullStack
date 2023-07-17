// Learn TypeScript:
//  - https://docs.cocos.com/creator/manual/en/scripting/typescript.html
// Learn Attribute:
//  - https://docs.cocos.com/creator/manual/en/scripting/reference/attributes.html
// Learn life-cycle callbacks:
//  - https://docs.cocos.com/creator/manual/en/scripting/life-cycle-callbacks.html

import DropItem from "./dropItem";

const {ccclass, property} = cc._decorator;

@ccclass
export default class NewClass extends cc.Component {

    @property(cc.Node)
    dropNode: cc.Node = null;

    @property(cc.Label)
    showNode: cc.Label = null;

    @property(cc.Prefab)
    prefab: cc.Prefab = null;

    // LIFE-CYCLE CALLBACKS:

    // onLoad () {}

    start () {
        var texts:string[]  = ["極限五個字","下一位","789","聽牌","明牌","槓紅中","123","456","789","9999","0000"];   
        this.set(texts);
    }

    set(texts: string[]){
        this.dropNode.removeAllChildren();
        for (var i = 0; i < texts.length; i++) {
            var newUnit = cc.instantiate(this.prefab);
            newUnit.getComponent(DropItem).set(texts[i]);
            this.dropNode.addChild(newUnit);
        }
    }

    setShowNode(text:string){
        this.showNode.string = text
        this.node.active =false;
    }

    drop(){
        this.node.active =true;
    }

    // update (dt) {}
}
