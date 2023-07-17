// Learn TypeScript:
//  - https://docs.cocos.com/creator/manual/en/scripting/typescript.html
// Learn Attribute:
//  - https://docs.cocos.com/creator/manual/en/scripting/reference/attributes.html
// Learn life-cycle callbacks:
//  - https://docs.cocos.com/creator/manual/en/scripting/life-cycle-callbacks.html

const {ccclass, property} = cc._decorator;

@ccclass
export default class NewClass extends cc.Component {

    @property(cc.Label)
    labels: cc.Label[] = [];

    start () {
    }

    updateMenu(playerIndex: number, cardId: number){
        this.labels[playerIndex].string = this.getWithoutValue(cardId);
    }

    getWithoutValue(withoutId:number){
        console.log(withoutId)
        switch(withoutId){
            case 0:
                return "萬";
            case 1:
                return "條";
            case 2:
                return "筒";
        }
    }
}
