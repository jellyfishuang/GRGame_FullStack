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
    exampleChiOption: cc.Node;
    exampleCancel: cc.Node;

    chiOptions: cc.Node[] = [];

    init(){
        if(this.exampleChiOption == null){
            this.exampleChiOption = this.node.getChildByName("example_chow_option");
        }
        if(this.exampleCancel == null){
            this.exampleCancel = this.node.getChildByName("example_cancel");
        }
    }

    showChowOptions(chiOptions:number[][]){
        this.init();
        this.node.removeAllChildren();
        for (let i = 0 ; i < chiOptions.length ; i++) {
            let newNode = cc.instantiate(this.exampleChiOption);
            newNode.getComponent(cc.Button).clickEvents[0].customEventData = i.toString();
            this.editChowOption(newNode, chiOptions[i]);
            this.node.addChild(newNode);
        }
        this.node.addChild(this.exampleCancel);
        this.show();
    }

    editChowOption(chiOption:cc.Node, chiCardIds:number[]){
        let chiOptionNode = chiOption.getChildByName("chi_option");
        let chiCards = chiOptionNode.children;
        for (let i = 0 ; i < chiCards.length ; i++) {
            chiCards[i].getComponent(Unit).set(chiCardIds[i]);
        }
    }

    start () {
    }

    show(){
        this.node.active = true;
    }

    hide(){
        this.node.active = false;
    }
}
