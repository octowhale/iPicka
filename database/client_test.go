package database

import (
	"fmt"
	"testing"
)

func Test_Redis(t *testing.T) {

	client, _ := New(Config{DatabaseType: "redis",
		Host:     "172.18.8.88",
		Port:     "53697",
		DBName:   "11",
		Password: ""})

	// ExampleClient(client)
	k := "naruto_md5_v6"
	v := "https://cdn.tangx.in/v6/naruto_md5_v6.jpg"

	client.Set(k, v)

	val, _ := client.Get(k)
	fmt.Println("val= ", val)
}
