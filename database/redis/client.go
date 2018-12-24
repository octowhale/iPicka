package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	// "github.com/gomodule/redigo/redis"
)

type Config struct {
	pool *redis.Client

	RedisHost     string
	RedisPort     string
	RedisDB       int
	RedisPassword string
}

// InitRedis return redis client
func (r *Config) InitRedis() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.RedisHost, r.RedisPort),
		Password: r.RedisPassword, // no password set
		DB:       r.RedisDB,       // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	r.pool = client
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
