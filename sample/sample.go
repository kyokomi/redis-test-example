package sample

import (
	"fmt"
	"log"

	"github.com/kyokomi/redis-test-example/foo"
	redis "gopkg.in/redis.v2"
)

var Service SampleService

func init() {

	f := foo.NewFoo()
	fmt.Println("sample init", f.Name)
	Service = SampleService{}
}

type SampleService struct {
	redisClient *redis.Client
}

func (s *SampleService) Initialize(redisClient *redis.Client) {
	s.redisClient = redisClient
}

func (s *SampleService) Do() string {
	var result string

	if res, err := s.redisClient.Set("foo2", "bar").Result(); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(res)
	}

	if res, err := s.redisClient.Get("foo2").Result(); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(res)

		result = res
	}

	return result
}

func (s *SampleService) Close() {
	s.redisClient.Close()
}
