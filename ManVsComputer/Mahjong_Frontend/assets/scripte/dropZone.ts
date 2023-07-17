// Learn TypeScript:
//  - https://docs.cocos.com/creator/manual/en/scripting/typescript.html
// Learn Attribute:
//  - https://docs.cocos.com/creator/manual/en/scripting/reference/attributes.html
// Learn life-cycle callbacks:
//  - https://docs.cocos.com/creator/manual/en/scripting/life-cycle-callbacks.html
import Unit from "./unit";
const {ccclass, property} = cc._decorator;

@ccclass
export default class NewClass extends cc.Component {

    @property(cc.Prefab)
    prefab: cc.Prefab = null;

    start () {
    }
    
    unHighlightLastCard(){
        let cardCount = this.node.childrenCount;
        if(cardCount > 0){
            this.node.children[cardCount - 1].getComponent(Unit).white();
        }
    }

    setCards(cardIds: number[], isBlueTheLast: boolean){
        this.node.removeAllChildren();
        let cardLength = cardIds.length;
        for(let i = 0 ; i < cardLength ; i++){
            this.addToDrop(cardIds[i]);
        }
        if(isBlueTheLast) this.node.children[cardLength - 1].getComponent(Unit).blue();
    }

    addToDrop(cardId:number){
        var newUnit = cc.instantiate(this.prefab);
        newUnit.getComponent(Unit).set(cardId);
        newUnit.getComponent(Unit).isStacked = true;
        this.node.addChild(newUnit);
        return newUnit;
    }
}
