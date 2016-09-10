// config.go
//
// Copyright 2015 © by Ollivier Robert <roberto@keltia.net>
//

/*
Package config implements my homemade configuration class

Looks into a YAML file for configuration options and returns a config.Config struct.

 	import "config"

 	rc := config.LoadConfig("tag")

 	or

    rc := config.LoadConfig("foo.toml")

In the first case, $HOME/.tag/config.toml will be loaded.  On Windows
rc will be serialized from TOML.
*/
package config

import (
	"io/ioutil"

	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/naoina/toml"
	"log"
)

// Source describe a given LDAP/AD server
type Source struct {
	Site   string
	Port   int
	Base   string
	Filter string
	Attrs  []string
}

// Config is the outer shell for config data
type Config struct {
	Verbose bool
	Sources map[string]Source
}

// Check the parameter for either tag or filename
func checkName(file string) string {
	// Check for tag
	if !strings.HasSuffix(file, ".toml") {
		// file must be a tag so add a "."
		return filepath.Join(os.Getenv("HOME"),
			fmt.Sprintf(".%s", file),
			"config.toml")
	}
	return file
}

// Basic Stringer for Config
func (c *Config) String() string {
	str := fmt.Sprintf("Verbose = %v\n", c.Verbose)
	for _, s := range c.Sources {
		str = str + fmt.Sprintf("ldap://%s:%d/%s\n  Filter: %s\n  Attrs: %v\n",
			s.Site, s.Port, s.Base, s.Filter, s.Attrs)
	}
	return str
}

// LoadConfig loads a file as a TOML document and return the structure
func LoadConfig(file string) (*Config, error) {
	// Check for tag
	sFile := checkName(file)

	c := new(Config)
	buf, err := ioutil.ReadFile(sFile)
	if err != nil {
		return c, fmt.Errorf("Can not read %s", sFile)
	}

	err = toml.Unmarshal(buf, &c)
	if err != nil {
		return c, fmt.Errorf("Error parsing toml %s: %v", sFile, err)
	}

	return c, err
}

// SetDefaults does what the name implies
func (c *Config) SetDefaults() {
	log.Fatalf("Please set the defaults in the config.toml file.")
}
