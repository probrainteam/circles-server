import mysqlLoader from '../loaders/mysql'
import express, { NextFunction, Request, Response } from "express";
/*
const basicQuery = async (query:string) : Promise<object> =>{
    const mysqlConnection = await mysqlLoader();
    const  [rows, fields] = await mysqlConnection.query(query);
    console.log(rows[0])
    mysqlConnection.destroy();
    // @TODO :: throw err logic
    return rows;
}

const example = async () => { // Would be delete ...
    const mysqlConnection = await mysqlLoader();
    const  [rows, fields] = await mysqlConnection.query("SHOW STATUS LIKE 'Threads_connected';");
    console.log(rows[0])
    mysqlConnection.destroy();
}

const createExample = async (query:string) : Promise<boolean>=> {
    const mysqlConnection = await mysqlLoader();
    const  [rows, fields] = await mysqlConnection.query(`CREATE ... ${query}`);
    console.log(rows[0])
    mysqlConnection.destroy();
    // @TODO :: change logic using basicQuery
    return true;
}

export default {basicQuery, example, createExample}*/