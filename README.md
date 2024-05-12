# Conversations

Chat application made with Go and React. Work in Progress...

## Build

1. Clone the repository
```console
git clone https://github.com/diwasrimal/conversations.git
cd conversations
```

2. Run the backend with your postgresql database
```console
cd backend
go build -o app .
DATABASE_URL="postgres://user:@host/chatdb" ./app
```

3. Run frontend
```console
cd frontend
npm install
npm run dev
```