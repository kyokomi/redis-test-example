package sample

import (
	"testing"

	redistest "github.com/soh335/go-test-redisserver"
	redis "gopkg.in/redis.v2"
)

func TestSampleService(t *testing.T) {

	server, err := redistest.NewServer(true, nil)
	if err != nil {
		t.Error(err)
	}
	defer server.Stop()

	host := server.Config["unixsocket"]

	opt := redis.Options{}
	opt.Addr = host
	opt.Network = "unix"

	client := redis.NewClient(&opt)

	Service.Initialize(client)

	result := Service.Do()
	if result != "bar" {
		t.Errorf("error sampleService.Do result: %s", result)
	}
}
