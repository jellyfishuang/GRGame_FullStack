// Learn TypeScript:
//  - https://docs.cocos.com/creator/manual/en/scripting/typescript.html
// Learn Attribute:
//  - https://docs.cocos.com/creator/manual/en/scripting/reference/attributes.html
// Learn life-cycle callbacks:
//  - https://docs.cocos.com/creator/manual/en/scripting/life-cycle-callbacks.html

const {ccclass, property} = cc._decorator;

@ccclass
export default class NewClass extends cc.Component {

    @property(cc.Node)
    directions: cc.Node[] = [];

    start () {
    }

    turnByIndex(index:number){
        for (let i = 0 ; i < this.directions.length ; i++) {
            if(index == i){
                this.directions[index].active = true;
            } else {
                this.directions[i].active = false;
            }
        }
    }
}
