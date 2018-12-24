package redis

import (
	"fmt"
	"testing"
)

func Test_Redis(t *testing.T) {

	client := Config{
		RedisHost:     "172.18.8.88",
		RedisPort:     "53697",
		RedisPassword: "",
		RedisDB:       11,
	}

	// ExampleClient(client)
	k := "naruto_md5"
	v := "https://cdn.tangx.in/v6/naruto.jpg"

	client.Set(k, v)

	val, _ := client.Get(k)
	fmt.Println("val= ", val)
}
