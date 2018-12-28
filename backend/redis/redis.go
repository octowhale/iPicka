package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type RedisAgent struct {
	Host       string
	Port       string
	Password   string
	DBName     string
	Expiration time.Duration
	client     *redis.Client
}

func NewRedisAgent(host, port, password, dbname string) *RedisAgent {

	return &RedisAgent{
		Host:     host,
		Port:     port,
		Password: password,
		DBName:   dbname,
	}
}

func tryConnect(host, port, password, dbname string) (client *redis.Client, err error) {

	dbInt, err := strconv.Atoi(dbname)
	if err != nil {
		log.Errorln(err)
		panic(err)
	}

	// if dbInt < 0 || dbInt > 15 {
	// 	log.Errorln(error.)
	// }

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		DB:       dbInt,
		Password: password,
	})

	log.Debugln("Ping-pong test: this should be only run once")
	_, err = client.Ping().Result()
	if err != nil {
		log.Errorf("Redis Ping() Error: %v", err)
		panic(err)
		// return nil, err
	}
	return client, nil

}

func (r *RedisAgent) Conn() (client *redis.Client, err error) {

	if r.client != nil {
		_, err = r.client.Ping().Result()
		if err != nil {
			return nil, err
		}
	} else {
		r.client, err = tryConnect(r.Host, r.Port, r.Password, r.DBName)
		if err != nil {
			return nil, err
		}
	}

	return r.client, nil
}

func (r *RedisAgent) Set(k, v string) (ok bool, err error) {

	client, err := r.Conn()
	if err != nil {
		log.Errorln(err)
	}

	// stat := client.Set(k, v, 0)

	// fmt.Println("args:", stat.Args())
	// fmt.Println("err:", stat.Err())
	// fmt.Println("name:", stat.Name())
	// fmt.Printf("stat.Result: ")
	// fmt.Println(stat.Result())
	// fmt.Println("string:", stat.String())
	// fmt.Println("val:", stat.Val())
	err = client.Set(k, v, 0).Err()
	if err != nil {
		log.Errorln(err)
		return false, nil
	}

	return true, nil
}

func (r *RedisAgent) Get(k string) (s string, err error) {
	client, err := r.Conn()
	if err != nil {
		log.Errorln(err)
	}

	stat := client.Get(k)
	err = stat.Err()
	if err != nil {
		log.Errorln(err)
		return "", err
	}

	return stat.Val(), nil
}

func (r *RedisAgent) Status() {
	client, err := r.Conn()
	if err != nil {
		// log.Errorln(err)
		panic(err)
	}

	ps := client.PoolStats()
	log.Infoln("redis pool TotalConns:", ps.TotalConns)
	log.Infoln("redis pool IdleConns:", ps.IdleConns)
	log.Infoln("redis pool StaleConns:", ps.StaleConns)
	log.Infoln("redis pool Timeouts:", ps.Timeouts)
	log.Infoln("redis pool Hits:", ps.Hits)
	log.Infoln("redis pool Misses:", ps.Misses)
}
