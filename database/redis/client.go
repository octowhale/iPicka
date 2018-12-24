package redis

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	// "github.com/gomodule/redigo/redis"
)

type Config struct {
	pool *redis.Client

	RedisHost     string
	RedisPort     string
	RedisDB       string
	RedisPassword string
}

func NewRedis(host, port, password, db string) *Config {
	// portInt, _ := strconv.Atoi(port)

	return &Config{RedisHost: host,
		RedisPort:     port,
		RedisPassword: password,
		RedisDB:       db}
}

// InitRedis return redis client
func (r *Config) InitRedis() *redis.Client {

	redisdb, err := strconv.Atoi(r.RedisDB)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.RedisHost, r.RedisPort),
		Password: r.RedisPassword, // no password set
		DB:       redisdb,         // use default DB
	})

	_, err = client.Ping().Result()
	if err != nil {
		panic(err)
	}
	// r.pool = client
	return client

}

func (r *Config) Set(k, v string) error {

	r.pool = r.InitRedis()
	defer r.pool.Close()
	err := r.pool.Set(k, v, 0).Err()
	if err != nil {
		// panic(err)
		return err
	}
	return nil
}

func (r *Config) Get(k string) (string, error) {

	r.pool = r.InitRedis()
	defer r.pool.Close()
	val, err := r.pool.Get(k).Result()
	if err != nil {
		// panic(err)
		return "", err
	}
	return val, nil
}
