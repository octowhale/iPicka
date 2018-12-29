package main

import (
	"encoding/json"
	"io/ioutil"
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

}

func init() {
	l := logger.LogConfig{
		ENV:   util.GetEnvDefault("ENV", "online"),
		Level: util.GetEnvDefault("LOG_LEVEL", "info"),
	}
	l.SetEnv(l.ENV)
	l.SetLevel(l.Level)
}
