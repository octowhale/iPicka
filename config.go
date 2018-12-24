package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"github.com/octowhale/iPicka/database/redis"
	log "github.com/octowhale/iPicka/log"
	"github.com/octowhale/iPicka/storage"
)

// Config for aliyun
type StorageConfig = storage.Config

type Config struct {
	ConfigFile string
	File       string
	StorageConfig
}

var config Config

func init() {
	flag.StringVar(&config.ConfigFile, "config", "~/.ipic.json", "Config File")
	flag.StringVar(&config.AccKey, "key", "", "Access Key for Bucket")
	flag.StringVar(&config.AccSec, "sec", "", "Access Secret for Bucket")
	flag.StringVar(&config.CustomDomain, "domain", "", "返回用户域名替代默认域名")
	flag.BoolVar(&config.HTTPSSchema, "https", false, "是否使用 https (default: false)")
	// flag.StringVar(&config.Endpoint, "endpoint", "cn-hangzhou.aliyun", "是否使用 https")
	flag.StringVar(&config.Region, "endpoint", "cn-hangzhou", "Bucket 的地区")
	flag.StringVar(&config.Bucket, "bucket", "", "存储 Bucket 名称")
	flag.BoolVar(&config.Internal, "internal", false, "是否使用 vpc 内网模式 (default: false)")
	// flag.StringVar(&config.File, "file", "", "file or directory to uploads")
}

func initConfig() error {
	// configPath := config.ConfigFile

	s := strings.Split(config.ConfigFile, string(os.PathSeparator))

	if s[0] == "~" {
		s[0] = os.Getenv("HOME")
		config.ConfigFile = strings.Join(s, string(os.PathSeparator))
	}

	_, err := os.Stat(config.ConfigFile)
	if os.IsNotExist(err) {
		log.Debug("Default Config is not exists")
		return err
	} else {

		log.Debug("Loading ... ", config.ConfigFile)
		configBytes, err := ioutil.ReadFile(config.ConfigFile)
		if err != nil {
			return err
		}

		err = json.Unmarshal(configBytes, &config)
		if err != nil {
			return nil
		}

		// fmt.Println(config)
	}
	return nil
}

func initRedis() *redis.Config {
	redis := &redis.Config{
		RedisHost:     "172.18.8.88",
		RedisPort:     "53697",
		RedisDB:       11,
		RedisPassword: "",
	}

	// client := redis.InitRedis()
	// return redis, nil
	return redis
}
