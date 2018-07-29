# gokv bench

This package contains a benchmark comparing the different Store implementations. Honestly, I'm not sure that these benchmarks really mean anything. I think that the most interesting part is just how easy it is with `gokv` to switch implementations by just injecting the store.

## Run

The benchmark below was run on a MacBook Pro (3,1 GHz Intel Core i5) with macOS.

The in-memory implementation is obviously not comparable to the other two.

```shell
❯ make bench
docker rm -f postgres_store ;\
                docker rm -f redis_store ;\
                docker run -d --name postgres_store -p 5432:5432 postgres:10 &&\
                docker run -d --name redis_store -p 6379:6379 redis:4-alpine
postgres_store
redis_store
e1b25c018927424a266d8c44229dc6102f17628b1980b8756869ec44d9758075
1a861e6e367201b5f92f74111df4eb9efb14796afa7e2673c95e7a66be13570d
sleep 3 &&\
                docker exec postgres_store psql -U postgres -c 'create database store;' &&\
                go test -bench .
CREATE DATABASE
goos: darwin
goarch: amd64
pkg: github.com/gokv/play
BenchmarkMem-4           1000000              1117 ns/op
BenchmarkPostgres-4         1000           1164375 ns/op
BenchmarkRedis-4            2000            584328 ns/op
PASS
ok      github.com/gokv/play    3.790s
```
