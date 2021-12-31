package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var WorkDir string

type config struct {
	Bind string
	DSN  string
}

var Config *config

func init() {
	WorkDir, _ = os.Getwd()
	if os.Getenv("DOCKER_HUB_RECORDER_ROOT") != "" {
		WorkDir = os.Getenv("DOCKER_HUB_RECORDER_ROOT")
	}

	bytes, err := ioutil.ReadFile(fmt.Sprintf("%v/config.yaml", WorkDir))
	if err != nil {
		panic(err)
	}

	Config = &config{}

	if err := yaml.Unmarshal(bytes, Config); err != nil {
		panic(err)
	}
}
