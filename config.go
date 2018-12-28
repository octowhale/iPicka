package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/octowhale/iPicka/backend"
	"github.com/octowhale/iPicka/storage"
)

type Config struct {
	Backend *backend.Config `json:"backend,omitempty"`
	Storage *storage.Config `json:"storage,omitempty"`
}

var config *Config

func init() {
	data, _ := ioutil.ReadFile(path.Join(os.Getenv("HOME"), "/", ".ipic.json"))
	json.Unmarshal(data, &config)

	// logrus.Infoln(config.Storage)
	// storageClient, _ := storage.New(config.Storage)
	// backendClient, _ := backend.New(config.Backend)

	// err := storageClient.Put("v7/2019-naruto.jpg", "/data/tmp/naruto.jpg")
	// fmt.Println(err)
	// backendClient.Set("naruto", "http://cdn.tangx.in/v7/2019-naruto.jpg")
	// s, _ := backendClient.Get("naruto")
	// fmt.Println(s)
}
