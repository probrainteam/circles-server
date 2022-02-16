import { InitiateMysqlEnviroment } from "../models/InitiateMysqlEnvironment";

export default async(): Promise<any>=>{
    const test = new InitiateMysqlEnviroment();
    test.connect();
    return;
}