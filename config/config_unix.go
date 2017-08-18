// config.go
//
// Copyright 2015 © by Ollivier Robert <roberto@keltia.net>
// +build !windows

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
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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
