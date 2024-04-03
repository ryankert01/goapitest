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

### Design

#### Database Design 1: Postgres Materialized Views

主要除存資料使用postgres，並且由於isAlive的Ads每秒可能有不同，因此每秒要清掉存在RAM的cache，且Requests的可能性為`100(age)*2(gender)*249(countries)*3(platform) ~ 1e5` 約為10000rps目標的10倍，因此放棄使用cache(eg. Redis)。  

**Materialized View**:  
Materialized View can maintain a precomputed list of records within the desired time range.

**Non-blocking Materialized View**:  
PostgreSQL allows a non-blocking refresh option (REFRESH MATERIALIZED VIEW CONCURRENTLY your_view). This allows you to query the materialized view even while it is being refreshed. The view will continue to show the old data until the refresh is complete. 

因此我使用**Non-blocking Materialized View**，並且每秒進行`REFRESH`，確保Ad is alive，也希望能達到最高效率。

#### Database Design 2: Update per minute

Base on我對Dcard的觀察，用戶主要集中在兩個國家(`TW`,`JP`)，因此Requests的可能性約為`100(age)*2(gender)*2(countries)*3(platform) ~ 1e3` 小於10000十倍，而且廣告投放常理來說是以小時為單位，因此可以使用Redis進行Caching，並且1hr清除Cache一次。

