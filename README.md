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

to turn down:

```bash
docker compose down
```

3. run jmeter test in cli

```bash
chmod +x ./test/jmeter.sh
./test/jmeter.sh
```

or, alternatively

```bash
bash ./test/jmeter.sh
```

### Development

1. install dependencies

```bash
go mod tidy
```

2. port docker environment's terminal

```bash
docker compose run --service-ports web bash
```

3. run jmeter in cli

```bash
jmeter -n -t ./test/test.jmx -l log.jtl
```

according to the [documentation](https://jmeter.apache.org/usermanual/get-started.html#non_gui).  
to enhance the jvm heap size, we set variable at front, as describe in the [doc](https://jmeter.apache.org/usermanual/get-started.html#running):

```bash
JVM_ARGS="-Xms3072m -Xmx3072m" jmeter -n -t ./test/test.jmx -l log.jtl
```

4. remove local postgre volume

```bash
docker volume ls
docker volume rm goapitest_postres-db
```