package backend

import (
	"errors"

	"github.com/octowhale/iPicka/backend/mysql"
	"github.com/octowhale/iPicka/backend/redis"
)

type BackendClient interface {
	Set(k, v string) (bool, error)
	Get(k string) (string, error)
	Ping()
}

func New(config *Config) (BackendClient, error) {

	switch config.Driver {
	case "redis":
		return redis.NewRedisAgent(config.Host, config.Port, config.Password, config.Dbname), nil

	case "mysql":
		return mysql.NewMysqlAgent(config.Host, config.Port, config.User, config.Password, config.Dbname)
	}
	return nil, errors.New("Invalid Backend Driver")
}
