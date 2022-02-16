// Static variables
import * as configure from "./configure.json";
const DEV_DB_URI: string = configure.develop.database.uri;
const DEV_DB_PORT: string = configure.develop.database.port;
const DEV_DB_HOST: string = configure.develop.database.host;
const DEV_DB_PASSWORD: string = configure.develop.database.password;
const DEV_DB_DATABASE: string = configure.develop.database.database;
const DEV_DB_USER: string = configure.develop.database.user;
const DEV_URI: string = configure.develop.uri;
const DEV_PORT: string = configure.develop.port;

const MAIN_DB_URI: string = configure.main.database.uri;
const MAIN_DB_PORT: string = configure.main.database.port;
const MAIN_URI: string = configure.main.uri;
const MAIN_PORT: string = configure.main.port;
const MAIN_DB_HOST: string = configure.main.database.host;
const MAIN_DB_PASSWORD: string = configure.main.database.password;
const MAIN_DB_DATABASE: string = configure.main.database.database;
const MAIN_DB_USER: string = configure.main.database.user;

function getDbUser(condition: string): string{
    if(condition === "dev")
        return DEV_DB_USER;
    else
        return MAIN_DB_USER;
}
function getDbHost(condition: string): string{
    if(condition === "dev")
        return DEV_DB_HOST;
    else
        return MAIN_DB_HOST;
}

function getDbPassword(condition: string): string{
    if(condition === "dev")
        return DEV_DB_PASSWORD;
    else
        return MAIN_DB_PASSWORD;
}

function getDbDatabase(condition: string): string{
    if(condition === "dev")
        return DEV_DB_DATABASE;
    else
        return MAIN_DB_DATABASE;
}
function getDomainUri(condition: string): string{
    if(condition === "dev")
        return DEV_URI;
    else
        return MAIN_URI;
}
function getDomainPort(condition: string): string{
    if(condition === "dev")
        return DEV_PORT;
    else
        return MAIN_PORT;
}
function getDbUri(condition: string): string{
    if(condition === "dev")
        return DEV_DB_URI;
    else
        return MAIN_DB_URI;
}

function getDbPort(condition: string): string{
    if(condition === "dev")
        return DEV_DB_PORT;
    else
        return MAIN_DB_PORT;
}
export { getDbUser, getDbHost, getDbPassword, getDbDatabase, getDomainUri, getDomainPort, getDbUri, getDbPort}