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

rc will be serialized from TOML.

File location: $HOME/.config/erc-search/
*/
package main

import (
	"os"
	"path/filepath"
)

var (
	basedir = filepath.Join(os.Getenv("HOME"), ".config", rcFile)
)
