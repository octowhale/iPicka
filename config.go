package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/octowhale/iPicka/database"
	"github.com/octowhale/iPicka/storage"
	"github.com/sirupsen/logrus"
)

// Config for aliyun
type StorageConfig = storage.Config
type DatabaseConfig = database.Config

type Config struct {
	ConfigFile string
	File       string
	StorageConfig
	DatabaseConfig
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

func init() {
	initConfig()
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
		logrus.Debugln("Default Config is not exists")
		return err
	}

	logrus.Debugln("Loading ... ", config.ConfigFile)
	configBytes, err := ioutil.ReadFile(config.ConfigFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return err
	}

	logrus.Debugln(config)
	return nil
}

var backend database.DatabaseClient

func init() {
	backend, _ = database.New(DatabaseConfig{
		DatabaseType: "mysql",
		Host:         "172.18.8.88",
		Port:         "60333",
		User:         "root",
		Password:     "SMdemT2Pm",
		DBName:       "Demo2",
	})
}

var client storage.StorageClient

func init() {
	client, err := storage.New(config.StorageConfig)

	if err != nil {
		panic(err)
	}
	fmt.Println(client)
}
