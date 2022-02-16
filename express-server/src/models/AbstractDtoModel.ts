import { MysqlConnector}  from "./MySqlConnector"

abstract class AbstractDtoModel {
    readonly connector;

    constructor(){
        this.connector = new MysqlConnector().connect();
    }
    getConnector(): any {
        return this.connector;
    }

}

export { AbstractDtoModel }