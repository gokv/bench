
bench:
	docker rm -f postgres_store ;\
		docker rm -f redis_store ;\
		docker run -d --name postgres_store -p 5432:5432 postgres:10 &&\
		docker run -d --name redis_store -p 6379:6379 redis:4-alpine
		sleep 3 &&\
		docker exec postgres_store psql -U postgres -c 'create database store;' &&\
		go test -bench .

cleanup:
	docker rm -f postgres_store ; docker rm -f redis_store
