// Learn TypeScript:
//  - https://docs.cocos.com/creator/manual/en/scripting/typescript.html
// Learn Attribute:
//  - https://docs.cocos.com/creator/manual/en/scripting/reference/attributes.html
// Learn life-cycle callbacks:
//  - https://docs.cocos.com/creator/manual/en/scripting/life-cycle-callbacks.html

const {ccclass, property} = cc._decorator;

export enum EffectId {
    Chi,
    Peng,
    Gang,
    Hu,
    Zm,
    Ting,
    Ming,
    Crit
}

@ccclass
export default class NewClass extends cc.Component {

    @property(cc.SpriteAtlas)
    img_atlas: cc.SpriteAtlas = null;

    @property(cc.Node)
    player: cc.Node[] = [];

    start () {
        
    }

    playEffect(player:number,action:number){

        var playerEffect=this.player[player];
        playerEffect.opacity=255
        playerEffect.getComponent(cc.Sprite).spriteFrame = this.img_atlas.getSpriteFrame(EffectId[action]);
        this.schedule(function(){
            playerEffect.scale+=0.2;
        },0.1,3);
        this.schedule(function(){
            playerEffect.opacity-=50;
        },0.1,10,0.1)
        playerEffect.scale=1;
    }
}
