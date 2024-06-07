# Conversations

Chat application made with Go, React and PostgreSQL.


## Build and Run

1. Development build

```console
$ git clone https://github.com/diwasrimal/conversations.git
$
$ cd conversations/backend
$ createdb chatdb
$ psql -d chatdb -f ./db/create_tables.sql
$ go build -o app .
$ POSTGRES_URL="postgres://user:password@host/chatdb" MODE="dev" ./app
$
$ cd ../frontend
$ npm install
$ npm run dev
```

2. Production build

In production build, the frontend build files are placed inside `backend/dist`
and are served by the backend server.

```console
$ git clone https://github.com/diwasrimal/conversations.git
$
$ cd conversations/frontend
$ npm install
$ npm run build
$ cp -r ./dist ../backend/
$
$ cd ../backend
$ createdb chatdb
$ psql -d chatdb -f ./db/create_tables.sql
$ go build -o app .
$ POSTGRES_URL="postgres://user:password@host/chatdb" MODE="prod" ./app
```
