// Learn TypeScript:
//  - https://docs.cocos.com/creator/manual/en/scripting/typescript.html
// Learn Attribute:
//  - https://docs.cocos.com/creator/manual/en/scripting/reference/attributes.html
// Learn life-cycle callbacks:
//  - https://docs.cocos.com/creator/manual/en/scripting/life-cycle-callbacks.html

const {ccclass, property} = cc._decorator;
import Player from "./Player";
import game from "./game";

@ccclass
export default class Unit extends cc.Component {

    @property(cc.SpriteAtlas)
    Mahjongs: cc.SpriteAtlas = null;

    @property
    player: number = 0;

    @property
    index: number = 0;

    // Stack means the card is in drop zone or stack zone
    isStacked: boolean = false;

    ID: number = 0;
    air: boolean = false;
    isPicked: boolean = false;
    isMerged: boolean = false;

    start() {
        let self = this;
        //建立監聽事件 滑鼠點擊時控制this.node.position
        if(self.player == 0){
            self.node.on(cc.Node.EventType.MOUSE_DOWN, function(){
                let player = self.node.parent.parent.parent.getComponent(Player);
                if(player == null) return;
                if(player.isInteractable && !self.isStacked){
                    self.node.parent.parent.parent.getComponent(Player).pick(self.index);
                }
            })
        }
    }

    jump(){
        this.isPicked = true;
        switch ( this.player ) { 
            case 3 : { this.node.x += 20; break ; } 
            case 2 : { this.node.y -= 20; break ; } 
            case 1 : { this.node.x -= 20; break ; } 
            case 0 : { this.node.y  = 20; break ; } 
        }
    }
    
    down(){
        this.isPicked = false;
        this.node.y = 0;
    }

    set(ID: number) {
        if(this.node.getComponent(cc.Sprite).spriteFrame == this.Mahjongs.getSpriteFrame("mingmah_00")){return;}
        let FileName:string;
        this.ID = ID;

        //ID直接對照檔名 萬0~8 條9~17 筒18~26 字27~33
        if(this.player == 0 && !this.isStacked){
            FileName = "handmah_" + ID;
        }
        else{
            FileName = "mingmah_" + ID;        
        }
        this.node.getComponent(cc.Sprite).spriteFrame = this.Mahjongs.getSpriteFrame(FileName);
    }

    //複製當前unit，加入子節點後疊至上層
    stack(){
        // console.log(this.node.scale)
        //複製this 加入子節點
        var newNode = new cc.Node();
        newNode.addComponent(cc.Sprite);
        //複製後的unit大小不一致，問題不明，暫時以手動調整
        switch(this.player){
            case 0:
                newNode.y += 30;
                break;
            case 1:
                newNode.y += 15;
                //newNode.setContentSize(70,70)
                newNode.scale = 0.58;
                break;
            case 2:
                newNode.y = 25;
                newNode.setScale(0.85,0.75);
                break;
            case 3:
                newNode.y += 15;
                newNode.scale = 0.58;
                break;
        }
        //檔名問題，先將player調整為3
        this.player =3;
        this.set(this.ID);
        newNode.getComponent(cc.Sprite).spriteFrame = this.node.getComponent(cc.Sprite).spriteFrame;
        this.node.addChild(newNode);
        // console.log(newNode.scale)
    }

    white(){
        this.node.color = new cc.Color(255,255,255,255);
    }

    blue(){
        this.node.color = new cc.Color(160,230,255,255);
    }

    back(){
        this.node.getComponent(cc.Sprite).spriteFrame = this.Mahjongs.getSpriteFrame("mingmah_00");
    }
}