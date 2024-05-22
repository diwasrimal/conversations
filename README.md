# Conversations

Chat application made with Go and React. Work in Progress...

## Build and run

1. Clone the repository
```console
git clone https://github.com/diwasrimal/conversations.git
cd conversations
```

2. Build backend and initialize postgresql database
```
cd backend
make
```
or 

```console
cd backend
psql createdb chatdb
psql -d chatdb -f ./db/create_tables.sql
go build -o app .
```

2. Run the backend connecting database
```console
DATABASE_URL="postgres://user:password@host/chatdb" ./app
```

3. Run frontend
```console
cd frontend
npm install
npm run dev
```
