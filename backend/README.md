# Restaurateur API

## Documentation

API documentation can be seen in Swagger UI on `/docs/index.html`
The API can be tested directly on that page by clicking on *try out*
and inputting query parameters (cookie authentication cannot be tested there).

Generate new documentation:

`$ swag init -g internal/api/api.go -o docs`

## Environment variables 

**DB_DSN**

Example:

`$ export DB_DSN='user=postgres dbname=postgres password=postgres sslmode=disable'`

**ORIGIN_ALLOWED**

Example:

`$ export ORIGIN_ALLOWED=http://localhost:3000`

## Database set-up

postgres >= 13

Before the server is started, the database schema is checked and tables are installed
if they are missing. If restaurants are not present in the database, they are downloaded before the 
server is started.

**Re-downloading data**

If you want to download fresh data about all restaurants and reinstall the restaurants schema,
you can use the following command:

`$ ./backend --download`

## Starting the server

`$ cd cmd/backend && go build `

`$ ./backend`

To configure listen port use the `-p` or `--port` flag
