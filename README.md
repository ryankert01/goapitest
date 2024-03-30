# goapitest

### Run

1. set up a `.env` file for db credentials, for example:

```
DB_USER=ryankert
DB_PASSWORD=ryankert
DB_NAME=ryankert
```

2. use docker compose to bring up db and web service

```bash
docker compose up
```

### Development

install dependencies

```bash
go mod tidy
```

port docker environment's terminal

```bash
docker compose run --service-ports web bash
```