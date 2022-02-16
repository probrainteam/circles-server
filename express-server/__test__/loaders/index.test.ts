const express = require('express');
const app = express();
import request from "supertest"
import loaders from "../../src/loaders"
describe("Routing test", ()=>{
    test("Get example/test", async ()=>{
        await loaders({ expressApp: app });
        const response = await request(app).get("/example/test")
        expect(response.statusCode).toBe(200)
    })

    test("Get test/example", async ()=>{
        await loaders({ expressApp: app });
        const response = await request(app).get("/test/example")
        expect(response.statusCode).toBe(404)
    })
})
    
    