package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func Test_Config(t *testing.T) {
	c := path.Join(os.Getenv("HOME"), "/", ".ipic.json")

	// util.IsFileExist(c)

	var j *Config

	data, _ := ioutil.ReadFile(c)

	json.Unmarshal(data, &j)

	fmt.Println(j)
	fmt.Println(j.Storage.Driver)
	fmt.Println(j.Backend.Driver)
}
