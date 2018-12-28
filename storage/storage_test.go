package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func Test_Storage(t *testing.T) {

	configFile := path.Join(os.Getenv("HOME"), ".ipic.json")
	fmt.Println(configFile)

	j := Storage{}

	b, _ := ioutil.ReadFile(configFile)
	json.Unmarshal(b, &j)

	fmt.Println(j)

	client, _ := New(&j)

	client.Put("v7/", "/data/tmp/naruto.jpg")

}
