package command

import (
	"encoding/json"
	"io/ioutil"
)

func configReader() (j Config) {
	b, _ := ioutil.ReadFile(configPath)
	j = Config{}
	json.Unmarshal(b, &j)

	return j
}

type Config struct {
	Bucket       string `json:"bucket,omitempty" example:"tangxin-test-02"`
	CustomDomain string `json:"customDomain,omitempty" example:"cdn.tangx.in"`
	Endpoint     string `json:"endpoint,omitempty" example:"cn-hangzhou.aliyuncs.com"`
	Key          string `json:"key,omitempty" example:"key-123"`
	Provider     string `json:"provider" example:"aliyunoss"`
	Region       string `json:"region,omitempty" example:"hangzhou"`
	Schema       string `json:"schema,omitempty" example:"https"`
	Sec          string `json:"sec,omitempty" example:"sec-123"`
}
