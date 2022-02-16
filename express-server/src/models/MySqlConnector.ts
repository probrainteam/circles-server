import { AbstractMysqlConnector } from "./AbstractMysqlConnector";
import { getDbDatabase } from "../conf/conf"

class MysqlConnector extends AbstractMysqlConnector{
    readonly _database : string;

    constructor(){
        super();
        this._database = getDbDatabase(process.argv[2]);
    }
    // @override
    public async connect(): Promise<object> {
        this.connection = await this.mysql.createConnection({
            host: this._host, 
            port: this._port,
            user: this._user,
            password: this._password,
            database: this._database,
            multipleStatements: false
        })
        return this.getConnection();
    }
}

export { MysqlConnector }