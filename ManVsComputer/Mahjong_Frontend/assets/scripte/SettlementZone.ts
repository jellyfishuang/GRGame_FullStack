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
    playerId: cc.Label = null;

    @property(cc.Label)
    point: cc.Label = null;

    @property(cc.Label)
    fanns: cc.Label = null;

    setContent(playerId: string, point: string, fanns: string){
        this.playerId.string = playerId;
        this.point.string = point;
        this.fanns.string = fanns;
    }

    start () {
    }
}
