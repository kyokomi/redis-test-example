package main

import (
	redis "gopkg.in/redis.v2"
	"log"
	"github.com/kyokomi/redis-test-example/sample"
	"fmt"
)

func init() {
	log.SetFlags(log.Llongfile)

	fmt.Println("main init")

	client := newRedisClient("localhost:6379", "tcp", 0)
	sample.Service.Initialize(client)
}

func main() {
	defer sample.Service.Close()

	sample.Service.Do()
}

// network: tcp or unix
func newRedisClient(addr, network string, db int64) *redis.Client {
	opt := redis.Options{}
	opt.Addr = addr
	opt.Network = network
	opt.DB = db

	client := redis.NewClient(&opt)

	return client
}
