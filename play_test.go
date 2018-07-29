package main

import (
	"database/sql"
	"os"
	"testing"

	"github.com/gokv/mem"
	"github.com/gokv/postgres"
	"github.com/gokv/redis"
	_ "github.com/lib/pq"
)

func BenchmarkMem(b *testing.B) {
	s := mem.New()
	defer s.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Do(s)
	}
}

func BenchmarkPostgres(b *testing.B) {
	var host string
	if host = os.Getenv("POSTGRES_HOST"); host == "" {
		host = "localhost"
	}

	db, err := sql.Open("postgres", "host="+host+" user=postgres dbname=store sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	s, err := postgres.New(db, "test_table")
	if err != nil {
		panic(err)
	}
	defer s.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Do(s)
	}
}

func BenchmarkRedis(b *testing.B) {
	var addr string
	if addr = os.Getenv("REDIS_ADDR"); addr == "" {
		addr = "localhost:6379"
	}

	s := redis.New(addr, os.Getenv("REDIS_PASS"))
	defer s.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Do(s)
	}
}
