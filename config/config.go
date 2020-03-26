package config

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	OLat             float64 `yaml:"oLat"`
	OLong            float64 `yaml:"oLong"`
	LimitDistance    float64 `yaml:"limitDistance"`
	CustomerDataFile string  `yaml:"customerDataFile"`
}

func GetConf(filename string) Config {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var str bytes.Buffer
	str.WriteString(dir)
	str.WriteString("/")
	str.WriteString(filename)

	source, err := ioutil.ReadFile(str.String())
	if err != nil {
		log.Fatal(err)
	}

	c := Config{}

	err = yaml.Unmarshal(source, &c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
