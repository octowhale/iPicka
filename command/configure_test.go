package command

import (
	"fmt"
	"testing"
)

func Test_configLoader(t *testing.T) {
	config := configLoader("../config/ipicka.json")

	fmt.Println(config)
}

func Test_ProfieLoader(t *testing.T) {
	key := "default"
	profile := ProfileLoader(key)

	fmt.Println(profile)
}
