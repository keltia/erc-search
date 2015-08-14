// config.go
//
// Copyright 2015 © by Ollivier Robert <roberto@keltia.net>
//

/*
 Package implement my homemade configuration class

 Looks into a YAML file for configuration options and returns a config.Config
 struct.

 	import "config"

 	rc := config.LoadConfig("foo.yml")

 rc will be serialized from YAML.
 */
package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"log"
)

type Config struct {
	Site       string
	Port       string
	LdapBase   string
	LdapFilter string
	Attrs      []string
}

// Load a file as a YAML document and return the structure
func LoadConfig(file string) (Config, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	c := new(Config)
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		log.Println("Error parsing yaml")
		return Config{}, err
	}

	return *c, err
}
