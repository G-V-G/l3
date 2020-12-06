# l3

**Run sql script:**

`psql -U <username> -d <dbname> -f db/script.sql ...`

**Or (in psql):**

`\i db/script.sql`

**Run server:**

`cd server && go run ./cmd/server -h <host> -p <port>`

**Run client (test requests):**

`node client/example`

**Result:**

Adding a new forum (no args) 

Error: { status: 400, body: { message: 'Forum name is not provided' } } 
 
=========================================================

Adding a new forum (no name) 

Error: { status: 400, body: { message: 'Forum name is not provided' } } 
 
=========================================================

Adding a new forum (no topic) 

Error: {
  status: 400,
  body: { message: 'Topic keyword name is not provided' }
} 
 
=========================================================

Adding a new forum (empty name) 

Error: { status: 400, body: { message: 'Forum name is not provided' } } 
 
=========================================================

Adding a new forum (empty topic) 

Error: {
  status: 400,
  body: { message: 'Topic keyword name is not provided' }
} 
 
=========================================================

Adding a new forum (name exists) 

Error: {
  status: 400,
  body: { message: 'Forum with this name or topic already exists' }
} 
 
=========================================================

Adding a new forum (topic exists) 

Error: {
  status: 400,
  body: { message: 'Forum with this name or topic already exists' }
} 
 
=========================================================

Adding a new forum 

Result:
{ status: 201 }

=========================================================

Adding a new user (no args) 

Error: { status: 400, body: { message: 'Username is not provided' } } 
 
=========================================================

Adding a new user (no name) 

Error: { status: 400, body: { message: 'Username is not provided' } } 
 
=========================================================

Adding a new user (no interests) 

Error: { status: 400, body: { message: 'Interests are not provided' } } 
 
=========================================================

Adding a new user (empty string interests) 

Error: { status: 400, body: { message: 'Interest cannot be empty' } } 
 
=========================================================

Adding a new user (empty name) 

Error: { status: 400, body: { message: 'Username is not provided' } } 
 
=========================================================

Adding a new user (existing) 

Error: {
  status: 400,
  body: { message: 'User with this name already exists' }
} 
 
=========================================================

Adding a new user 

Result:
{ status: 201 }

=========================================================

Getting user (not registered) 

Error: { status: 400, body: { message: 'No such user' } } 
 
=========================================================

Getting user (empty name) 

Error: { status: 400, body: { message: 'User name is not provided' } } 
 
=========================================================

Getting user (null name) 

Error: { status: 400, body: { message: 'User name is not provided' } } 
 
=========================================================

Getting user 

Result:
{
  status: 200,
  body: [ { name: 'Barbara', interests: [ 'golang' ] } ]
}

=========================================================

Getting forum (not registered) 

Error: { status: 400, body: { message: 'No such forum' } } 
 
=========================================================

Getting forum (empty name) 

Error: { status: 400, body: { message: 'Forum name is not provided' } } 
 
=========================================================

Getting forum (null name) 

Error: { status: 400, body: { message: 'Forum name is not provided' } } 
 
=========================================================

Getting forum 

Result:
{
  status: 200,
  body: [ { name: 'Gophers', Topic: 'golang', users: [ 'Barbara' ] } ]
}

=========================================================

Getting all users 

Result:
{
  status: 200,
  body: [
    { name: 'Bob', interests: [ 'Jojo References', 'Games' ] },
    { name: 'Nick', interests: [ 'Games' ] },
    { name: 'Simon', interests: [ 'Books' ] },
    { name: 'Barbara', interests: [ 'golang' ] }
  ]
}

=========================================================

Getting all forums 

Result:
{
  status: 200,
  body: [
    {
      name: 'Jojo References',
      Topic: 'jojo bizzare adventure',
      users: [ 'Bob' ]
    },
    { name: 'Movies fan', Topic: 'Movies', users: [] },
    { name: 'Book enjoyer', Topic: 'Books', users: [ 'Simon' ] },
    { name: 'Gaming', Topic: 'Games', users: [ 'Bob', 'Nick' ] },
    { name: 'Gophers', Topic: 'golang', users: [ 'Barbara' ] }
  ]
}

=========================================================
