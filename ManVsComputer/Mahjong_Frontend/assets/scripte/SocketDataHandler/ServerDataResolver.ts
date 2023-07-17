import { ServerData} from "./DataSchema";
export default class DataResolver{

    static getServerData(jsonData: any): ServerData {
        let serverData = new ServerData();
        let serverDataJson = jsonData.Data;
        if(serverDataJson == null) return serverData;
        return Object.assign(serverData, serverDataJson);
    }
}