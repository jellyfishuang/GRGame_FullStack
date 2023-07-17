import Unit from "./unit";
import { GameRule, ActionId} from "./SocketDataHandler/DataSchema";

export class CardUtil{
    static newCard(prefab:cc.Prefab, cardId:number, idx:number, playerId:number, isStacked:boolean):cc.Node{
        var newUnit = cc.instantiate(prefab);
            newUnit.getComponent(Unit).index = idx;
            newUnit.getComponent(Unit).player = playerId;
            newUnit.getComponent(Unit).isStacked = isStacked;
            newUnit.getComponent(Unit).set(cardId);
            return newUnit;
    }

    static sortCards(cardIdList:number[], cardList:Unit[]){
        this.sortCardIdList(cardIdList);
        for (let i = 0; i < cardIdList.length; i++) {
            cardList[i].set(cardIdList[i]);
        }
    }

    static sortCardIdList(cardIdList:number[]){
        cardIdList.sort(function(a, b){return a - b;});
    }
}

export class UiUtil{
    static initAvailableInteractions(gameRule: GameRule){
        let availableActions = [];
        if(gameRule.CanEat) availableActions.push(ActionId.Chow);
        if(gameRule.CanPong) availableActions.push(ActionId.Pong);
        if(gameRule.CanKong) availableActions.push(ActionId.ExposedKong);
        if(gameRule.CanKong) availableActions.push(ActionId.ConcealedKong);
        if(gameRule.CanKong) availableActions.push(ActionId.AddKong);
        if(gameRule.CanMing) availableActions.push(ActionId.Ming);
        if(gameRule.CanCrit) availableActions.push(ActionId.Crit);
        availableActions.push(ActionId.Ting);
        availableActions.push(ActionId.Hu);
        availableActions.push(ActionId.SelfHu);
        availableActions.push(ActionId.Buhua);
        return availableActions;
    }
}

export class AppUtil{
    static async sleep(time : number) : Promise<void>{
        return new Promise<void>((res,rej)=>{
            setTimeout(res,time);
        });
    }
}