package redis

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_Redis(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	redis := RedisAgent{
		Host:       "172.18.8.88",
		Port:       "53697",
		Password:   "",
		DBName:     "3",
		Expiration: 0,
	}

	// redis.Conn()
	redis.Set("a", "https://www.a.com")
	redis.Set("b", "https://www.b.com")
	redis.Set("d", "https://www.d.com")
	// redis.Set("b", redis)

	s, err := redis.Get("a")
	fmt.Println(s)
	s, err = redis.Get("b")
	fmt.Println(s)
	s, err = redis.Get("d")
	fmt.Println(s)
	s, err = redis.Get("b")
	fmt.Println(s)

	fmt.Println(err)

	redis.Status()

}
