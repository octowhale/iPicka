package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/mkideal/cli"
)

const (
	configFile = "config/ipicka.json"
)

type argConf struct {
	cli.Helper
	Profile      string `cli:"profile"`
	AccKeyID     string `cli:"accKeyID" usage:"ACCESS KEY for bucket" prompt:"ACCESS KEY"`
	AccKeySecret string `cli:"accKeySecret" usage:"ACCESS SECRET for bucket" prompt:"ACCESS SECRET"`
	BucketName   string `cli:"bucketName" usage:"bucket name" prompt:"BUCKET NAME"`
	Endpoint     string `cli:"endpoint" usage:"endpoint for your bucket" prompt:"BUCKET ENDPOINT"`
	Domain       string `cli:"domain" usage:"custom domain. Cname for bucket domain" prompt:"CUSTOM DOMAIN"`
}

// Configure generate config
func Configure() {
	cli.Run(new(argConf), func(ctx *cli.Context) error {

		config := ConfigLoader()
		fmt.Println(config)

		argv := ctx.Argv().(*argConf)
		// ctx.String("username=%s, password=%s\n", argv.AccKeyID, argv.AccKeySecret)

		s := `{"AccKeyID": "%s", "AccKeySecret": "%s", "BucketName": "%s", "Endpoint": "%s", "Domain": "%s"}`
		s = fmt.Sprintf(s, argv.AccKeyID, argv.AccKeySecret, argv.BucketName, argv.Endpoint, argv.Domain)

		fmt.Println(s)

		var profile Profile
		err := json.Unmarshal([]byte(s), &profile)
		if err != nil {
			log.Fatalf("%s", err)
		}

		config.Config[argv.Profile] = profile

		// fmt.Println(config)

		configByte, _ := json.MarshalIndent(config, "", "  ")
		fmt.Println(string(configByte))

		ConfigWriter(configByte, configFile)
		return nil
	})
}

type Config struct {
	Config map[string]Profile `json:"config"`
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

func ConfigWriter(byte []byte, file string) {

	fobj, _ := os.OpenFile(file, os.O_RDONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer fobj.Close()
	fobj.WriteString(string(byte))

}
