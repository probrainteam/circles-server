import express, { NextFunction, request } from 'express';
const app = express();
const port = process.env.PORT || 4000;
const bodyParser = require('body-parser');
var header = `<h1>Probrain server</h1>`
var body = `<button onclick = "window.location = window.location.href+'test/visit';">Test</button> 
<button onclick = "window.location = window.location.href+'test/ssh';">Get ssh</button> 
<button onclick = "window.location = window.location.href+'api/v1/user';">Register</button> 
<button>Select * (X)</button> <button>Filtering (X)</button> <button>Sort (X)</button> 
<button>Send Mail (X)</button> <button>Permit join (X)</button> <button>Deny join (X)</button> 
<button>Virtual account (X)</button>` 
var copy = `<p style='position:fixed; bottom:0;  width:100%;'>&copy; 2022 Probrain dev</p>`;
var count = 0;
var ip;
var Logger = function (req: express.Request, res: express.Response, next: NextFunction) {
  console.log(`[LOG] ${(new Date()).toLocaleString()} :  ${req.ip.slice(7,)}\nConnected at \x1b[1;3;33;94m${req.get('host')+req.url}\x1b[0m`);
  next();
};

app.use(bodyParser.json())
app.use(Logger);
app.get('/', (req: express.Request, res: express.Response) => {
  res.send(header+body+copy);
  count++;
});

app.get('/test/visit', (req: express.Request, res: express.Response) => {
  ip = req.headers['x-forwarded-for'] ||  req.connection.remoteAddress;
  ip = ip?.slice(7,);
  res.send(`<h1>Hi ${ip} <br>
  Todays :  ${count}</h1>`+copy);
});
app.get('/test/ssh', (req: express.Request, res: express.Response) => {
  const homedir = require('os').homedir()+'/.ssh';
  var fs = require('fs');
  var list;
  fs.readdir(homedir, function(error: any, filelist: any){
    list = filelist;
    res.send(`<h3>SSH List : ${list}</h3>`+copy)
  })
});
app.get('/test/ssh/1', (req: express.Request, res: express.Response) => {
  res.json([{id:1}])
});
app.post('/api/v1/user', (req: express.Request, res:express.Response) => {
  console.log('user post api exec');
  var email = req.body['email']
  console.log(email);
  res.json([{id:1, username:email}]);
  // res.send(JSON.stringify(req.body));
});
app.listen(port, () => {
  console.log(`Club management server listening on port ${port}!`);
});