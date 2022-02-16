const fs = require('fs');
const path = require('path');
import { AbstractMysqlConnector } from "./AbstractMysqlConnector";
import { getDbDatabase } from "../conf/conf"

class InitiateMysqlEnviroment extends AbstractMysqlConnector{
    readonly _database : string;

    constructor(){
        super();
        this._database = getDbDatabase(process.argv[2]);
    }
    // @override
    public async initialize() : Promise<boolean> {
        const connection = await this.connect();
        let result = false;

        try{  // Table 생성 Query 날림
            const [rows, fields] = await connection.query("SHOW TABLES;")
            
            // Table 유무는 현재 table row개수가 0이라면 없다고 판정.
            // @TODO :: 해당 로직을 구체적으로 변경
            if (Object.keys(rows).length === 0){
                console.warn("> There is no tables")
                console.warn("> Create tables... ")
                // Table 생성 query문을 '../conf/tableConfig.sql'로 부터 읽어와 execute
                const tableCreateQuery : string = fs.readFileSync(path.join(__dirname, '../conf/tableConfig.sql')).toString();
                await connection.query(tableCreateQuery);
                console.warn("> Success !")
            }
            console.log("MYSQL Intialized ✅");
            result = true;
        } catch(error: any){ // Error -> Table 생성 쿼리 중 오류 발생
            console.log("MYSQL Failed to Intialized ❌");
            if(process.argv[2] === "dev") // 개발환경인 경우 error terminal에 출력
                console.error(error);
            else{
                // @TODO :: logger file에 해당 내용 저장
            }
        } finally{
            // Query를 보내던 안보내던 connection은 destroy가 되어야함
            // destroy 되지 않으면 connection이 살아있어 db thread leak이 발생
            // mysql에서 SHOW STATUS LIKE 'Threads_connected';쿼리를 통해 현재 thread connection을
            // 볼 수 있음
            connection.destroy();
            return result;
        }
        
    }
    // @override
    public async connect(): Promise<any> {
        this.connection = await this.mysql.createConnection({
            host: this._host, 
            port: this._port,

            user: this._user,
            password: this._password,
            multipleStatements:true
        })
        try{
            await this.connection.query(`USE ${this._database}`)
        } catch(error:any){
            console.warn(`> There is no database ${this._database}`)
            console.warn("> Create database... ")
            await this.connection.query(`CREATE DATABASE ${this._database};`)
            await this.connection.query("SHOW DATABASES;")
            console.log("> Success !")
            await this.connection.query(`USE ${this._database}`)
        } finally{
            return this.getConnection();
        }
    }
}

export {InitiateMysqlEnviroment}