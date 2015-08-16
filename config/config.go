// config.go
//
// Copyright 2015 © by Ollivier Robert <roberto@keltia.net>
//

/*
 Package implement my homemade configuration class

 Looks into a YAML file for configuration options and returns a config.Config struct.

 	import "config"

 	rc := config.LoadConfig("tag")

 	or

    rc := config.LoadConfig("foo.yml")

 In the first case, $HOME/.tag/config.yml will be loaded.  On Windows
 rc will be serialized from YAML.
 */
package config

import (
	"io/ioutil"

	"log"
	"fmt"
	"os"
	"path/filepath"

	"github.com/naoina/toml"
)

type Config struct {
	Site   string
	Port   int
	Base   string
	Filter string
	Attrs  []string
}

// Basic Stringer for Config
func (c *Config) String() string {
	return fmt.Sprintf("ldap://%s:%d/%s\n  Filter: %s\n  Attrs: %v",
	c.Site, c.Port, c.Base, c.Filter, c.Attrs)
}

// Load a file as a TOML document and return the structure
func LoadConfig(file string) (*Config, error) {
	// Check for tag
	if _, err := os.Stat(file); err != nil {
		// file must be a tag so add a "."
		file = fmt.Sprintf(".%s", file)
		file = filepath.Join(os.Getenv("HOME"),
						     file,
			                 "config.toml")
	}

	c := new(Config)
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return c, err
	}

	err = toml.Unmarshal(buf, &c)
	if err != nil {
		log.Println("Error parsing toml %s", file)
		return c, err
	}

	return c, err
}

// Set defaults
func (c *Config) SetDefaults() {
	c.Site   = DEF_SERVER
	c.Port   = DEF_PORT
	c.Base   = DEF_BASE
	c.Filter = DEF_FILTER
	c.Attrs  = DEF_ATTRS
}