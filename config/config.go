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
	"fmt"
)

type Config struct {
	Site       string
	Port       int
	LdapBase   string
	LdapFilter string
	Attrs      []string
}

// Basic Stringer for Config
func (c *Config) String() string {
	return fmt.Sprintf("ldap://%s:%d/%s\n  Filter: %s\n  Attrs: %v",
	c.Site, c.Port, c.LdapBase, c.LdapFilter, c.Attrs)
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

// Set defaults
func (c *Config) SetDefaults() {
	c.Site =     DEF_SERVER
	c.Port =     DEF_PORT
	c.LdapBase = DEF_BASE
	c.LdapFilter = DEF_FILTER
	c.Attrs =      DEF_ATTRS
}