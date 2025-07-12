# Drunk Dwarves API

Drunk Dwarves API is written in [Go language](https://go.dev/). It is used for a Drunk Dwarves website.

## Setup

### Docker

To install a DB to your machnine you can use Docker:  
```docker compose up```

### Run locally

To install all the required dependencies to run the application run:  
```go mod download```

To run the migrations, run:  
```goose up```

### Live Reload

To enable Live Reload you can use [`air`](https://github.com/air-verse/air). To install air run:
```go install github.com/air-verse/air@latest```

To run the app with live reloads just run: 
```air```

## Environment

To set up the environment, `.env` file is used. Example:
```
PORT=8080
DB_URL=postgres://*user name*:*user password*@*host_url*:*host_port*/*db_name*?sslmode=disable

GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://*user name*:*user password*@*host_url*:*host_port*/*db_name*?sslmode=disable
GOOSE_MIGRATION_DIR=sql/schema

API_PATH=./cmd/api/main.go
```

## Commands

### Generating Feature Group

To generate a migration you need to have [`goimports`](https://pkg.go.dev/golang.org/x/tools/cmd/goimports). To install goose run:  
```go install golang.org/x/tools/cmd/goimports@latest```

To generate a feature required for correct jet work files first compile the `gen-feat` script using:  
```go build -o gen-feat ./cmd/gen/feat/main.go```

By executing this script, required feature directory and `dto`, `handlers`, `services` and `router` files will be generated. Generated feature will automatically be added to the general router

### Migrations

To generate a migration you can use [`goose`](https://github.com/pressly/goose). To install goose run:  
```go install github.com/pressly/goose/v3/cmd/goose@latest```

To create a migration run:  
```goose create migration_name go```

To apply all pending migrations run:  
```goose up```

To apply only the next migration run:  
```goose up-by-one```

To revert a single the migration run:  
```goose down```

To check the migration status run:  
```goose status```

### Query Builder

To access database [`jet`](https://github.com/go-jet/jet) is used in this project. To install jet run:  
```go install github.com/go-jet/jet/v2/cmd/jet@latest```

To generate all required for correct jet work files first compile the `prep-db` script using:  
```go build -o prep-db ./cmd/prep/db/main.go```

By executing this script, required jet files will be generated

### Documentation

To create the documentation run you need to install [`swag`](https://github.com/swaggo/swag). To install swag run:  
```go install github.com/swaggo/swag/cmd/swag@latest```  

To generate the documentation from code comments first compile the `prep-docs` script using:  
```go build -o prep-docs ./cmd/prep/docs/main.go```

By executing this script, api documentation will be generated. To access the documentation visit `/api/swagger`. For local development use:
```http://localhost:8080/api/swagger```

## Structure
