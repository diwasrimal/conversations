BINARY = app
DB = chatdb

all: build

build:
	@go build -o $(BINARY) .

test:
	@go test -v ./...

clean:
	@go clean
	rm -f $(BINARY)

dbreset:
	@psql -d $(DB) -f ./db/create_tables.sql
	@psql -d $(DB) -f ./db/reset_tables.sql