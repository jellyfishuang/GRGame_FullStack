import { GameState, ActionId, ClientData, Command, CommandId} from "./DataSchema";
export default class DataCreator{
    static createConnectionData(): string {
        let command = new Command();
        command.Command = CommandId[CommandId.Connect];
        return JSON.stringify(command);
    }

    static createCreateRoomData(userId: string, gameRule: string): string {
        let data = new ClientData();
        data.UserId = userId;
        data.GameRule = gameRule;

        let command = new Command();
        command.Command = CommandId[CommandId.Createroom];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }

    static createDealTilesData(): string {
        let command = new Command();
        command.Command = CommandId[CommandId.DealTiles];
        return JSON.stringify(command);
    }

    static createOpenGameData(): string {
        let command = new Command();
        command.Command = CommandId[CommandId.OpenGame];
        return JSON.stringify(command);
    }

    static createSelectWithoutData(withoutId:number, playerIndex:number): string {
        let data = new ClientData();
        data.Player = playerIndex;
        data.Tiles = withoutId;

        let command = new Command();
        command.Command = CommandId[CommandId.WithoutTiles];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }

    static createChangeThreeTilesData(tiles: number[], playerIndex:number): string {
        let data = new ClientData();
        data.Player = playerIndex;
        data.Tiles = tiles;

        let command = new Command();
        command.Command = CommandId[CommandId.ChangeTiles];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }

    static createChowData(tiles: number[], playerIndex:number){
        let data = new ClientData();
        data.Action = ActionId[ActionId.Chow];
        data.Player = playerIndex;
        data.Tiles = tiles;

        let command = new Command();
        command.Command = CommandId[CommandId.Action];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }

    static createPongData(tiles: number[], playerIndex:number){
        let data = new ClientData();
        data.Action = ActionId[ActionId.Pong];
        data.Player = playerIndex;
        data.Tiles = tiles;

        let command = new Command();
        command.Command = CommandId[CommandId.Action];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }

    static createKongData(tiles: number[], playerIndex:number){
        let data = new ClientData();
        data.Action = ActionId[ActionId.ExposedKong];
        data.Player = playerIndex;
        data.Tiles = tiles;

        let command = new Command();
        command.Command = CommandId[CommandId.Action];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }

    static createConcealedKongData(tiles: number[], playerIndex:number){
        let data = new ClientData();
        data.Action = ActionId[ActionId.ConcealedKong];
        data.Player = playerIndex;
        data.Tiles = tiles;

        let command = new Command();
        command.Command = CommandId[CommandId.Action];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }

    static createAddKongData(tiles: number[], playerIndex:number){
        let data = new ClientData();
        data.Action = ActionId[ActionId.AddKong];
        data.Player = playerIndex;
        data.Tiles = tiles;

        let command = new Command();
        command.Command = CommandId[CommandId.Action];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }

    static createDiscardTileData(playerId: string, playerIndex: number, dicardTile: number){
        let data = new ClientData();
        data.UserId = playerId;
        data.Player = playerIndex;
        data.DiscardTile = dicardTile;

        let command = new Command();
        command.Command = CommandId[CommandId.DiscardTile];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }

    static createNothingData(playerIndex:number){
        let data = new ClientData();
        data.Action = ActionId[ActionId.Nothing];
        data.Player = playerIndex;

        let command = new Command();
        command.Command = CommandId[CommandId.Action];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }

    static createAIDiscardTileData(playerId: string, playerIndex: number){
        let data = new ClientData();
        data.UserId = playerId;
        data.Player = playerIndex;

        let command = new Command();
        command.Command = CommandId[CommandId.AIDiscardTile];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }

    static createAITestData(rounds: number, gameRule: string){
        let data = new ClientData();
        data.Rounds = rounds;
        data.GameRule = gameRule;

        let command = new Command();
        command.Command = CommandId[CommandId.AITest];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }

    static createHuData(tile: number, playerIndex:number){
        let data = new ClientData();
        data.Action = ActionId[ActionId.Hu];
        data.Player = playerIndex;
        data.Tiles = tile;

        let command = new Command();
        command.Command = CommandId[CommandId.Action];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }

    static createSelfHuData(tile: number, playerIndex:number){
        let data = new ClientData();
        data.Action = ActionId[ActionId.SelfHu];
        data.Player = playerIndex;
        data.Tiles = tile;

        let command = new Command();
        command.Command = CommandId[CommandId.Action];
        command.Data = JSON.stringify(data)
        return JSON.stringify(command);
    }
}