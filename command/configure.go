package command

import "github.com/mkideal/cli"

// ConfigureT is the config os project
type ConfigureT struct {
	cli.Helper
	Profile      string `cli:"profile" usage:"profile" dft:"default" prompt:"profile "`
	Provider     string `cli:"provider" usage:"Provider" prompt:"Provider[aliyunoss/qcloudcos/qiniu] "`
	Key          string `cli:"key" usage:"Bucket Access Key ID" prompt:"Access Key ID "`
	Sec          string `cli:"sec" usage:"Bucket Access Key Secret" prompt:"Access Key Secret "`
	Bucket       string `cli:"bucket" usage:"Bucket Name" prompt:"Bucket Name "`
	Endpoint     string `cli:"endpoint" usage:"Bucket Endpoint" prompt:"Bucket Endpoint "`
	Region       string `cli:"region" usage:"Region" prompt:"Bucket Region "`
	Schema       string `cli:"schema" usage:"Schema" prompt:"Schema[http/https] "`
	CustomDomain string `cli:"customdomain" usage:"Custom Domain" prompt:"Custom Domain "`
	Prefix       string `cli:"prefix" usage:"objectKey prefix" prompt:"objectKey Prefix "`
}
