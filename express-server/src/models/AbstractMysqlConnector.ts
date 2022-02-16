import { getDbHost, getDbPassword, getDbPort, getDbUser } from "../conf/conf"

abstract class AbstractMysqlConnector {
    readonly _host: string;
    readonly _user: string;
    readonly _password: string;
    readonly _port: string;
    protected connection: any;
    protected mysql = require('promise-mysql2');

    constructor(){
        // @TODO .env 기반으로 변경
        this._port = getDbPort(process.argv[2]);
        this._host = getDbHost(process.argv[2]); // ex) localhost
        this._user = getDbUser(process.argv[2]); // ex) root
        this._password = getDbPassword(process.argv[2]); // ex) qwer1234
    }
    public async connect(): Promise<any> {
        this.connection = await this.mysql.createConnection({
            host: this._host, 
            port: this._port,
            user: this._user,
            password: this._password,
            multipleStatements: false
        })
        return this.getConnection();
    }
    public getConnection(): any {
        return this.connection;
    }
}
export {AbstractMysqlConnector};