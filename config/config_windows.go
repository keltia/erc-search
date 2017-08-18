// config.go
//
// Copyright 2015 © by Ollivier Robert <roberto@keltia.net>
//

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
		return filepath.Join(os.Getenv("%LOCALAPPDATA%"),
			"erc-search",
			"config.toml")
	}
	return file
}
