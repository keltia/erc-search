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
	"github.com/keltia/erc-search/config"
	"log"
)

const (
	rcFile = "erc-search"
	Version = "0.1"
)

var (
	ctx *context
)

type context struct {
	cnf     *config.Config
	verbose bool
}

func (ctx *context) NewSource(name string) (config.Source) {
	// Do the actual connect
	return ctx.cnf.Sources[name]
}

// Start here
func main() {
	// Load config file if any
	cnf, err := config.LoadConfig(rcFile)
	if err != nil {
		log.Printf("Warning: can't load %s\n", rcFile)
		cnf.SetDefaults()
	}

	ctx = &context{
		cnf:     cnf,
		verbose: false,
	}

	// Parse CLI
	flag.Parse()

	if fVerbose {
		ctx.verbose = true
		log.Printf("Default config:\n%s", cnf.String())
	}

	// We need at least one argument
	if flag.Arg(0) == "" {
		log.Fatalln("Error: You must specify a search string")
	}

	// Are we trying to find a given machine?
	if fWorkStation {
		searchForMachine(ctx, flag.Arg(0))
	} else {
		searchForPeople(ctx, flag.Arg(0))
	}
	// We're done
	log.Printf("Shutting down…")
}
