export enum CommandId {
    Connect,
    Createroom,
    DealTiles,
    ChangeTiles,
    WithoutTiles,
    OpenGame,
    DiscardTile,
    Action,
    AIDiscardTile,
    AITest
}
export enum ActionId {
    Chow,
    Pong,
    ExposedKong,
    ConcealedKong,
    AddKong,
    Hu,
    SelfHu,
    Nothing,
    DiscardTile,
    Buhua,
    ChangeThreeTiles,
    SelectWithout,
    Ming,
    Ting,
    Crit
}

export enum GameState{
    StateCreateRoom   = 0, // 創建房間,設定玩法,發牌
	StateDealTiles    = 1, // 發牌
	StateChangeTiles  = 2, // 換牌
	StateWithoutTiles = 3, // 定缺
	StateBuHua        = 4, // 補花
	StateOpenGame     = 5, // 開門(發第一張牌給莊家)
	StateDraw         = 6, // 摸牌
	StateAction       = 7, // 動作
	StateDiscard      = 8, // 丟牌
	StateGameOver     = 9 // 遊戲結束
}

export class Command{
    Command: string;
    Data: string;
}

export class ClientData{
    UserId: string;
    GameRule: string;
    Player: number;
    DiscardTile: number;
    Action: string;
    Tiles: any;
    Rounds: number;
}

export class ServerData{
    RoomId: number;
    TilesSea: number[];
    UnOpenPool: number[];
    PrevailingWind: number;
    PlayerNow: number;
    RealPlayer: number;
    Action: number;
    ContinueCount: number;
    RoundCount: number;
    Gameover: boolean;
    OntableTile: number;
    DoAction: DoActionData;
    Players: PlayerData[];
    NothingDo: boolean;
    GameRule: GameRule;
    GameState: number;

    isGameInit(){
        return this.GameState == GameState.StateDealTiles;
    }

    isPreGameStage(){
        return this.GameState == GameState.StateChangeTiles||
        this.GameState == GameState.StateWithoutTiles;
    }

    isGameOver(){
        return this.GameState == GameState.StateGameOver;
    }

    isUserPlayerThrow(){
        return this.PlayerNow == this.RealPlayer;
    }

    isDiscardTile(){
        return this.GameState == GameState.StateDiscard;
    }
}

export class GameRule{
    JokerTurnIntos: number[];
    LimitTai: number;
    CanEat: boolean;
    CanPong: boolean;
    CanKong: boolean;
    ChangeTileSameColor: boolean;
    ChangeTileCount: number;
    KongAfterHuUseSimpleRule: boolean;
    Debug: boolean;
    CanUseJokerAsEye: boolean;
    CanJokersingleGon: boolean;
    CanJokerGon: boolean;
    CanMing: boolean;
    LogAICsv: boolean;
    JokerNumbers: number;
    GameMode: number;
    CanCrit: boolean;
    LogAIParameter: boolean;
    GuoShouHu: boolean;
    FanCount: number[];
    XorTable: number[];
    IsbuHua: boolean;
}

export class DoActionData{
    Action: string;
    PriviousPlayer: number;
    Player: number;
    Tiles: number[];
}

export class PlayerData{
    DealerWind: number;
    HandTiles: number[];
    OnDrawTile: number;
    DiscardTiles: number[];
    ShowTiles: number[];
    MeldTiles: number[];
    FlowerTiles: number[];
    Point: number;
    CanChow: boolean;
    CanChowSet: number[][];
    CanPon: boolean;
    CanPonSet: number[];
    CanKong: boolean;
    CanKongSet: number[][];
    CanConcealedKong: boolean;
    CanConcealedKongSet: number[][];
    CanAddKong: boolean;
    CanAddKongSet: number[];
    CanHu: boolean;
    CanHuTile: number;
    CanHuFaanList: number[];
    CanHuFaanListStr: string[];
    CanReadyHand: boolean;
    CanSelfHu: boolean;
    CanSelfHuTile: number;
    WithoutTile: number;
}
