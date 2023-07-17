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
    exampleKongOption: cc.Node;
    exampleCancel: cc.Node;

    chiOptions: cc.Node[] = [];

    init(){
        if(this.exampleKongOption == null){
            this.exampleKongOption = this.node.getChildByName("example_kong_option");
        }
        if(this.exampleCancel == null){
            this.exampleCancel = this.node.getChildByName("example_cancel");
        }
    }

    showKongOptions(kongCards:number[][], concealed: boolean){
        this.init();
        this.node.removeAllChildren();
        for (let i = 0 ; i < kongCards.length ; i++) {
            let newNode = cc.instantiate(this.exampleKongOption);
            this.bindKongEvent(newNode, concealed, String(i));
            this.editKongOption(newNode, kongCards[i]);
            this.node.addChild(newNode);
        }
        this.node.addChild(this.exampleCancel);
        this.show();
    }

    getKongFunctionName(concealed: boolean){
        return concealed ? "concealedKongOption": "kongOption";
    }

    bindKongEvent(button:cc.Node, concealed: boolean, eventData: string){
        button.getComponent(cc.Button).clickEvents[0].handler = this.getKongFunctionName(concealed);
        button.getComponent(cc.Button).clickEvents[0].customEventData = eventData;
    }

    editKongOption(chiOption:cc.Node, kongCardIds:number[]){
        let kongOptionNode = chiOption.getChildByName("kong_option");
        let kongCards = kongOptionNode.children;
        for (let i = 0 ; i < kongCards.length ; i++) {
            kongCards[i].getComponent(Unit).set(kongCardIds[i]);
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
