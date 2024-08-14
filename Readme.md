# Go (Fiber + Gorm + Postgres)

This project is for POC Golang for backend development
Currently has only user registeration service

### Setup

1. Duplicate config/example.env to config.env
2. Fill Postgres config as you want (PG_XXX=) except PG_HOST=localhost

```
PG_HOST=localhost
PG_PORT=5432
PG_USER=db_user_username
PG_PASS=db_user_password
PG_DB=database_name
```

### Run

1. Run Docker Postgres `docker-compose --env-file ./config/.env  up -d`
2. Run `go run ./app` to run API service
3. Run `go run ./test/...` to run all unit tests
