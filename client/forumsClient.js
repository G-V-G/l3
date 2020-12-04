const http = require('http');

const connection = {
  hostname: '192.168.1.3',
  port: '5000'
};

// response sender

const makeRquest = (options, responseJSON = null) => new Promise((res, rej) => {
  const { method } = options;
  const responseStringified = JSON.stringify(responseJSON);
  const resData = {};
  const responseOptions = { ...connection, ...options };
  if (method !== 'POST')
    responseOptions.headers = {
      "Content-Type": "application/json",
      "Content-Length": Buffer.byteLength(responseStringified)
    };
  const servRespose = http.request(responseOptions, response => {
    response
      .setEncoding('utf-8')
      .on('data', data =>  (
        resData.status === 200 &&
        resData.status !== 201 ?
          resData.body = JSON.parse(data) :
          resData.message = data.trim()
      ))
      .on('close', () => {
        
        return(
        resData.status === 200 || 
        resData.status === 201 ?
          res(resData) : rej(resData)
      )})
  });
  servRespose
    .on('response', r => (resData.status = r.statusCode))
    .on('error', e => rej(e));
  if (method === 'POST')
    servRespose.end(responseStringified);
  else
    servRespose.end();
});

// client requests

const addForum = (name, topic) => makeRquest(
  { method: 'POST', path: '/forums' },
  { name, topic }
);

const addUser = (name, interests) => makeRquest(
  { method: 'POST', path: '/users' },
  { name, interests }
);

const getForum = name => makeRquest(
  { method: 'POST', path: '/forum' },
  { name }
);

const getUser = name => makeRquest(
  { method: 'POST', path: '/user' },
  { name }
);

const getForums = () => makeRquest({ method: 'GET', path: '/forums' });
const getUsers = () => makeRquest({ method: 'GET', path: '/users' });

module.exports = {
  addForum,
  getForum,
  getForums,
  addUser,
  getUser,
  getUsers
}
