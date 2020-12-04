const forumsClient = require('./forumsClient');

const testData = [
  ['addForum', 'Adding a new forum (no args)', []],
  ['addForum', 'Adding a new forum (no name)', [undefined, 'golang']],
  ['addForum', 'Adding a new forum (no topic)', ['Gophers', undefined]],
  ['addForum', 'Adding a new forum', ['Gophers', 'golang']],
  ['addUser', 'Adding a new user (no args)', []],
  ['addUser', 'Adding a new user (no name)', [undefined, ['golang']]],
  ['addUser', 'Adding a new user (no interests)', ['Barbara', undefined]],
  ['addUser', 'Adding a new user', ['Barbara', ['golang']]],
  ['getUser', 'Getting user (not registered)', ['Bill']],
  ['getUser', 'Getting user', ['Barbara']],
  ['getForum', 'Getting forum (not registered)', ['Java Hell']],
  ['getForum', 'Getting forum', ['Gophers']],
  ['getUsers', 'Getting all users', []],
  ['getForums', 'Getting all forums', []]
];

const sendTestResponses = async () => {
  const separator = '\n=========================================================\n';
  for (const [ method, comment, args ] of testData)
    try {
      console.log(comment, '\n');
      const responseFn = forumsClient[method];
      const res = await responseFn(...args);
      console.log('Result:\n')
      console.dir(res, { depth: null })
      console.log(separator)
    } catch (e) {
      console.error('Error:', e, '\n', separator);
    }
};

sendTestResponses();

// .then/.catch example
// forumsClient.getUser('Bob')
//   .then(r => console.dir(r, { depth: null }))
//   .catch(e => console.error(e));
