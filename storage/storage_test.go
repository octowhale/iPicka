package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_Storage(t *testing.T) {

	// configFile := path.Join(os.Getenv("HOME"), ".ipic.json")
	configFile := "config.json"
	fmt.Println(configFile)

	j := Config{}

	b, _ := ioutil.ReadFile(configFile)
	json.Unmarshal(b, &j)

	fmt.Println(j)

	client, _ := New(&j)

	client.Put("v7/naruto-test.jpg", "/data/tmp/naruto.jpg")

}
