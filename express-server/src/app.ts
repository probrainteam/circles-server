const express = require('express');
import { getDomainUri, getDomainPort, getDbUri, getDbPort } from './conf/conf'
import loaders from './loaders'

async function startServer() {
    const app = express();
    await loaders({ expressApp: app });
    
    const MODE: string = process.argv[2]; // main or dev
    const PORT: string = getDomainPort(MODE) // 포트
    const domain: string = `${getDomainUri(MODE)}:${PORT}`; // uri:port
    const db: string = `${getDbUri(MODE)}:${getDbPort(MODE)}`; // uri:port

    console.warn(`
    ---------------------------------------------
        Start Server with Condition :: ${MODE}
        Using below options ...\n
        Domain : ${domain}
        db : ${db}
    ---------------------------------------------
    `);

    app.listen(PORT,() =>{
        console.log(`
        ################################################
        🛡️  Server listening on port: ${PORT}🛡️
        ################################################
      `);
    });
  }
  
  startServer();

export {startServer}
