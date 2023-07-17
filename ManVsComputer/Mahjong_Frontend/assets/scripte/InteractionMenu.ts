// Learn TypeScript:
//  - https://docs.cocos.com/creator/manual/en/scripting/typescript.html
// Learn Attribute:
//  - https://docs.cocos.com/creator/manual/en/scripting/reference/attributes.html
// Learn life-cycle callbacks:
//  - https://docs.cocos.com/creator/manual/en/scripting/life-cycle-callbacks.html

import { ActionId } from "./SocketDataHandler/DataSchema";

const {ccclass, property} = cc._decorator;

@ccclass
export default class InteractionMenu extends cc.Component {
    exampleOption: cc.Node;
    exampleCancel: cc.Node;
    buttonCache: Map<number, cc.Node> = new Map<number, cc.Node>();

    init(){
        if(this.exampleOption == null){
            this.exampleOption = this.node.getChildByName("example_option");
        }
        if(this.exampleCancel == null){
            this.exampleCancel = this.node.getChildByName("example_cancel");
        }
    }

    initButtons(buttonIds:number[]){
        this.init();
        for (let i = 0 ; i < buttonIds.length ; i++) {
            let buttonId = buttonIds[i];
            this.buttonCache.set(buttonId, this.createButtonById(buttonId));
        }
        console.log("available buttons = ", this.buttonCache);
    }

    createButtonById(buttonId:number): cc.Node{
        switch(buttonId){
            case ActionId.Chow : return this.createChowButton();
            case ActionId.Pong : return this.createPongButton();
            case ActionId.ExposedKong : return this.createKongButton();
            case ActionId.ConcealedKong : return this.createConcealedKongButton();
            case ActionId.AddKong : return this.createAddKongButton();
            case ActionId.Ting : return this.createTinButton();
            case ActionId.Hu : return this.createHuButton();
            case ActionId.SelfHu : return this.createSelfHuButton();
            case ActionId.Ming : return this.createMingButton();
            case ActionId.Crit : return this.createCritButton();
        }
        return null;
    }

    editButtonLabel(button:cc.Node, labelString:string){
        let label = button.getChildByName("Background").getChildByName("Label").getComponent(cc.Label);
        label.string = labelString;
    }

    editButtonClickEvent(button:cc.Node, functionName:string){
        button.getComponent(cc.Button).clickEvents[0].handler = functionName;
    }

    createChowButton(): cc.Node{
        return this.createButton("吃", "interactionButtonChow");
    }

    createPongButton(): cc.Node{
        return this.createButton("碰", "interactionButtonPong");
    }

    createKongButton(): cc.Node{
        return this.createButton("槓", "interactionButtonKong");
    }

    createConcealedKongButton(): cc.Node{
        return this.createButton("暗槓", "interactionButtonConcealedKong");
    }

    createAddKongButton(): cc.Node{
        return this.createButton("加槓", "interactionButtonAddKong");
    }

    createTinButton(): cc.Node{
        return this.createButton("聽牌", "interactionButtonTin");
    }

    createHuButton(): cc.Node{
        return this.createButton("胡", "interactionButtonHu");
    }

    createSelfHuButton(): cc.Node{
        return this.createButton("自摸", "interactionButtonSelfHu");
    }

    createCritButton(): cc.Node{
        return this.createButton("爆擊", "interactionButtonCrit");
    }

    createMingButton(): cc.Node{
        return this.createButton("明牌", "interactionButtonMing");
    }

    createAddFlowerButton(): cc.Node{
        return this.createButton("補花", "interactionButtonAddFlower");
    }

    createButton(labelString:string, functionName:string): cc.Node{
        let newNode = cc.instantiate(this.exampleOption);
        this.editButtonLabel(newNode, labelString);
        this.editButtonClickEvent(newNode, functionName);
        this.node.addChild(newNode);
        return newNode;
    }
    
    showByIndex(buttoneIds:number[], isCancelable: boolean){
        if(buttoneIds.length == 0) return;
        this.node.removeAllChildren();
        for (let i = 0 ; i < buttoneIds.length ; i++) {
            this.node.addChild(this.buttonCache.get(buttoneIds[i]));
        }
        if(isCancelable) this.node.addChild(this.exampleCancel);// example cancel is good enough and always be last one
        this.show();
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
