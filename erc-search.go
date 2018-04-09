// erc-search.go
//
// Small interface to get information from our corporate LDAP
//
// Copyright 2015 © by Ollivier Robert <roberto@keltia.net>

/*
 Package main implement a basic wrapper around LDAP search function.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	rcFile  = "erc-search"
	Version = "0.3"
)

var (
	ctx *context
)

type context struct {
	cnf *Config
}

func NewSource(name string) *Source {
	// Do the actual connect
	if s, ok := ctx.cnf.Sources[name]; ok {
		return s
	}
	return nil
}

// Start here
func main() {
	// Load file if any
	cnf, err := LoadConfig()
	if err != nil {
		log.Printf("Warning: can't load %s\n", rcFile)
		cnf.SetDefaults()
	}

	ctx = &context{
		cnf: cnf,
	}

	// Parse CLI
	flag.Parse()

	if fVersion {
		fmt.Printf("%s v%v\n", rcFile, Version)
		os.Exit(0)
	}

	verbose("Default config:\n%s", cnf.String())

	// We need at least one argument
	if flag.Arg(0) == "" {
		log.Fatalln("Error: You must specify a search string")
	}

	// Are we trying to find a given machine?
	if fWorkStation {
		searchForMachine(flag.Arg(0))
	} else {
		searchForPeople(flag.Arg(0))
	}
	// We're done
	log.Printf("Shutting down…")
}
