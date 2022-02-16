const figlet = require('figlet');
import expressLoader from './express';
import mysqlLoader from './mysql';

export default async ({ expressApp } : {expressApp: any}) => {
    console.log(figlet.textSync('Circles - server', {
        horizontalLayout: 'default',
        verticalLayout: 'default',
        whitespaceBreak: true
    }));

    console.warn("MYSQL in Intialize sequence ...")
    await mysqlLoader();
 
    console.warn("Express in Intialize sequence ...")
    await expressLoader({ app: expressApp });
    console.log('Express Intialized ✅');
}