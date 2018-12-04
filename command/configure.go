package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/mkideal/cli"
)

type argConf struct {
	cli.Helper
	AccKeyID     string `cli:"accKeyID" usage:"ACCESS KEY for bucket" prompt:"ACCESS KEY"`
	AccKeySecret string `cli:"accKeySecret" usage:"ACCESS SECRET for bucket" prompt:"ACCESS SECRET"`
	BucketName   string `cli:"bucketName" usage:"bucket name" prompt:"BUCKET NAME"`
	Endpoint     string `cli:"endpoint" usage:"endpoint for your bucket" prompt:"BUCKET ENDPOINT"`
	Domain       string `cli:"domain" usage:"custom domain. Cname for bucket domain" prompt:"CUSTOM DOMAIN"`
}

// Configure generate config
func Configure() {
	cli.Run(new(argConf), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argConf)
		ctx.String("username=%s, password=%s\n", argv.AccKeyID, argv.AccKeySecret)

		profile := map[string]string{"AccKeyID": argv.AccKeyID, "AccKeySecret": argv.AccKeySecret, "BucketName": argv.BucketName, "Endpoint": argv.Endpoint, "Domain": argv.Domain}

		data, _ := json.Marshal(profile)

		fmt.Printf("%s\n", data)

		ioutil.WriteFile("config/ipicka.json", data, 0644)
		return nil
	})
}

type Config struct {
	Config map[string]Profile
}

type Profile struct {
	// Provider     string `json:"Provider,omitempty"`
	AccKeyID     string `json:"AccKeyID"`
	AccKeySecret string `json:"AccKeySecret"`
	BucketName   string `json:"BucketName"`
	Domain       string `json:"Domain,omitemtpy"`
	Endpoint     string `json:"Endpoint"`
}

// ConfigLoader loads config file and unmarshal it to json
func ConfigLoader() (config Config) {
	configFile := "config/ipicka.json"

	// check config file exist
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		log.Fatalf("config file : %s is not exist", configFile)
		os.Exit(1)
	}

	// unmarshal
	configByte, _ := ioutil.ReadFile(configFile)

	// var config Config
	json.Unmarshal(configByte, &config)
	// log.Info(config.Config["aliyun"].AccKeyID)

	return
}

func ProfileLoader(profileKey string) (profile Profile) {
	config := ConfigLoader()

	// profile
	// var profile Profile
	if _, ok := config.Config[profileKey]; !ok {
		log.Fatalf("Profile(%s) is not exist")
		os.Exit(1)
	}

	// log.Info(config.Config[profileKey])
	return config.Config[profileKey]
}
