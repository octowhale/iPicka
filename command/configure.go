package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mkideal/cli"
)

type configureT struct {
	cli.Helper
	Profile      string `cli:"profile" usage:"specify profile" dft:"default" prompt:"Profile Name "`
	Provider     string `cli:"provider" usage:"Provider" dft:"oss" prompt:"Profile Name "`
	AccKeyID     string `cli:"accKeyID" usage:"Bucket Access Key ID" prompt:"Access Key ID "`
	AccKeySecret string `cli:"accKeySecret" usage:"Bucket Access Key Secret" prompt:"Access Key Secret "`
	BucketName   string `cli:"bucketName" usage:"Bucket Name" prompt:"Bucket Name "`
	Endpoint     string `cli:"endpoint" usage:"Bucket Endpoint" prompt:"Bucket Endpoint "`
	Domain       string `cli:"domain" usage:"Custom Domain" prompt:"Custom Domain "`
}

var configure = &cli.Command{
	Name: "configure",
	Desc: "configure your profile",
	Argv: func() interface{} { return new(configureT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*configureT)
		// ctx.String("Hello, child command, I am %s\n", argv.Profile)
		configureCommand(argv)
		return nil
	},
}

func configureCommand(argv *configureT) {
	// fmt.Println("Hello, ", argv)
	filePath := "config/ipicka.json"
	ConfigWriter(filePath, argv)
}

// Config is the config struct of ipicka , by json type
type Config struct {
	Config map[string]Profile `json:"config"`
}

// Profile is config struct for image docker provider
type Profile struct {
	Provider     string `json:"provider,omitempty"`
	AccKeyID     string `json:"AccKeyID"`
	AccKeySecret string `json:"AccKeySecret"`
	BucketName   string `json:"BucketName"`
	Endpoint     string `json:"Endpoint"`
	Domain       string `json:"Domain,omitempty"`
}

func configLoader(filePath string) (config Config) {

	// Assuming config file exists
	cByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// unmarshal
	if len(cByte) == 0 {
		cByte = []byte(`{"config":{"default":{}}}`)
	}
	err = json.Unmarshal(cByte, &config)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return
}

// ProfileLoader return profile for image dock
func ProfileLoader(key string) (profile Profile) {
	filePath := "config/ipicka.json"
	config := configLoader(filePath)

	profile = config.Config[key]

	if _, ok := config.Config[key]; !ok {
		log.Fatalf("%s profile not exist", key)
		os.Exit(1)
	}

	profile = config.Config[key]
	return
}

// ConfigWriter dump config file into file
func ConfigWriter(filePath string, argv *configureT) {

	// Assuming filePath exists
	s := fmt.Sprintf(`{
"Provider": "%s",
"AccKeyID": "%s",
"AccKeySecret": "%s",
"BucketName": "%s",
"Domain": "%s",
"Endpoint": "%s"
}`, argv.Provider,
		argv.AccKeyID,
		argv.AccKeySecret,
		argv.BucketName,
		argv.Domain,
		argv.Endpoint,
	)

	// fmt.Println(s)
	log.Println(s)

	config := configLoader(filePath)
	var profile Profile
	err := json.Unmarshal([]byte(s), &profile)
	if err != nil {
		log.Fatal(err)
	}
	config.Config[argv.Profile] = profile

	cData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	flag := os.O_WRONLY | os.O_TRUNC | os.O_CREATE
	fobj, err := os.OpenFile(filePath, flag, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fobj.Close()
	n, err := fobj.WriteString(string(cData))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)
}
