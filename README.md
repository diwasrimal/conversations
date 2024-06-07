# Conversations

Chat application made with Go and React.


## Build and Run

The frontend build files are placed inside `backend/dist` and are served by
the backend server.

```console
$ git clone https://github.com/diwasrimal/conversations.git
$ cd conversations/frontend
$ npm install
$ npm run build
$ cp -r ./dist ../backend/
$ cd ../backend
$ psql createdb chatdb
$ psql -d chatdb -f ./db/create_tables.sql
$ go build -o app .
$ POSTGRES_URL="postgres://user:password@host/chatdb" ./app
```
