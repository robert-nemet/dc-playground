# Docker Compose Playground

## Echo App

Echo app is simple Go API app several endpoints and DB layer.

API:

| name | path | method | payload | what it does |
|------|------|--------|---------|--------------|
| echo | /echo | POST | `{"msg": "Hello World" }` | Prints out message it receives |
| counter set | /counter | POST | none | increase counter in db for a day |
| counter get | /counter | GET | none | retrieves a counter value | 

Environment variables:

| name | what |
|------|------|
| APP_PORT | port on which app work |
| DB_HOST | Db host name |
| DB_PORT=5432 | db port |
| DB_USER | db user |
| DB_PASSWORD | db user password |
| DB_NAME | db name |
| DB_TYPE | db type, default is pg for the postgres |

App work with the postgres or mariadb.

Example:

```toml
APP_PORT=9999
DB_HOST="pg"
DB_PORT=5432
DB_USER="root"
DB_PASSWORD="password"
DB_NAME="docker"
DB_TYPE="PG"
```

## Issues

MariaDB requires to delete ./maria folder for each change to the docker-compose file, if changes affects mariadb(like passwords, etc.)

## Observations

Start setup with `make start_observer`. On `http://localhost:3000` you will see Grafana. Login with `admin` and `grafana`. Prometheus is added as data source.
To start load test run `make restart COMPOSE_FILE=compose-extended.yml APP=vegeta`. Duration is defined in `compose-extended.yml` file.

Setup is done on OSX. It should work on Linux as well.
