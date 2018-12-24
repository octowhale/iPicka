package database

import (
	"errors"

	"github.com/octowhale/iPicka/database/mysql"
	"github.com/octowhale/iPicka/database/redis"
)

type DatabaseClient interface {
	Get(k string) (string, error)
	Set(k, v string) error
}

func New(config Config) (DatabaseClient, error) {
	switch config.DatabaseType {
	case "redis":
		client := redis.NewRedis(config.Host, config.Port, config.Password, config.DBName)
		// client := redis.Config{RedisHost: config.Host, RedisPort: config.Port, RedisPassword: config.Password, RedisDB: config.DBName}
		return client, nil
	case "mysql":
		client, err := mysql.NewMysqlAgent(config.User, config.Password, config.Host, config.Port, config.DBName)
		return client, err
	}
	return nil, errors.New("Invalid DatabaseType")
}
