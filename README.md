# Sample Go Application

This is a sample go application based on DDD.


## Start app on local without Docker

Required: PostgreSQL is running on your local machine.

```sh
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_PORT=5432
export DB_NAME=go-app-sample
make start-api
```

## Start app on local using Docker

Requirement: Docker is running on your local machine.

```sh
make run
```
