// config.go
//
// Copyright 2015 © by Ollivier Robert <roberto@keltia.net>
// +build windows

package main

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	basedir = filepath.Join(os.Getenv("%LOCALAPPDATA%"), rcFile)
)
/*
File location: %LOCALAPPDATA%\ERC-SEARCH\
*/
