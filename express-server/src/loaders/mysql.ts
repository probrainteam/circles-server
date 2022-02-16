import { InitiateMysqlEnviroment } from "../models/InitiateMysqlEnvironment";

export default async (): Promise<boolean> => {
    const initiateor = new InitiateMysqlEnviroment();
    
    return await initiateor.initialize();
}
