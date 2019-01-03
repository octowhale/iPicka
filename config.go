package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/octowhale/iPicka/logger"
	"github.com/octowhale/iPicka/util"
	"github.com/sirupsen/logrus"

	"github.com/octowhale/iPicka/backend"
	"github.com/octowhale/iPicka/storage"
)

func init() {
	l := logger.LogConfig{
		ENV:   util.GetEnvDefault("ENV", "online"),
		Level: util.GetEnvDefault("LOG_LEVEL", "info"),
	}
	l.SetEnv(l.ENV)
	l.SetLevel(l.Level)
}

type Config struct {
	Backend *backend.Config `json:"backend,omitempty"`
	Storage *storage.Config `json:"storage,omitempty"`
}

var config *Config

var storageClient storage.StorageClient
var backendClient backend.BackendClient

func init() {

	data, _ := ioutil.ReadFile(path.Join(os.Getenv("HOME"), "/", ".ipic.json"))
	json.Unmarshal(data, &config)

	logrus.Debugf("main init: init storage Client")
	// storageClient, err := storage.New(config.Storage)
	var err error
	storageClient, err = storage.New(config.Storage)
	if err != nil {
		logrus.Fatalf("%v", err)
	}
	storageClient.Ping()
	logrus.Debugf("main init: init storage Client End")

	logrus.Debugf("main init: init backend Client")
	backendClient, err = backend.New(config.Backend)
	if err != nil {
		logrus.Fatalf("%v", err)
	}
	backendClient.Ping()
	logrus.Debugf("main init: init backend Client End")

}
