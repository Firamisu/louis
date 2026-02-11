# Louis App


## Prerequisites
- [Goose](https://github.com/pressly/goose) - migrartions management 
- [SQLC](https://docs.sqlc.dev/en/latest/) - sql to go codegen


## Setting up the database
Start the Postgres db container
```bash
docker compose up -d
```


## Running the API
```bash
make run
```

## Adding a new table
1. Create a new migration file under `internal/adapters/postgres/migrations`
```bash
goose -s create <name> sql
```

2. Run the migrations
```bash
goose up
```

3. Generate the SQLC code
```bash
sqlc generate
```
