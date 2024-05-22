BINARY = app
DB = chatdb

all: build db

build:
	@go build -o $(BINARY) .

test:
	@go test -v ./...

db:
	@createdb $(DB)
	@psql -d $(DB) -f ./db/create_tables.sql 

dbreset:
	@psql -d $(DB) -f ./db/delete_tables.sql
	@psql -d $(DB) -f ./db/create_tables.sql
	@psql -d $(DB) -f ./db/reset_tables.sql

clean:
	@go clean
	rm -f $(BINARY)
