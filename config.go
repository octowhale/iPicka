package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/octowhale/iPicka/logger"
	"github.com/octowhale/iPicka/util"

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
	storageClient, err := storage.New(config.Storage)
	if err != nil {
		log.Fatal("%+v", err)
	}
	storageClient.Ping()

	backendClient, err := backend.New(config.Backend)
	if err != nil {
		log.Fatal("%+v", err)
	}
	backendClient.Ping()

}

func init() {
	l := logger.LogConfig{
		ENV:   util.GetEnvDefault("ENV", "online"),
		Level: util.GetEnvDefault("LOG_LEVEL", "info"),
	}
	l.SetEnv(l.ENV)
	l.SetLevel(l.Level)
}
