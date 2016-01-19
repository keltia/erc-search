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
	"log"
	"github.com/keltia/erc-search/config"
)

const (
	RcFile = "erc-search"
)

// Start here
func main () {
	// Load config file if any
	config, err := config.LoadConfig(RcFile); if err != nil {
		log.Printf("Warning: can't load %s, using defaults\n", RcFile)
		config.SetDefaults()
	}

	// Parse CLI
	flag.Parse()

	if fVerbose {
		log.Printf("Default config:\n%s", config.String())
	}

	// We need at least one argument
	if flag.Arg(0) == "" {
		log.Fatalln("Error: You must specify a search string")
	}

	// Do the actual connect
	myldap, err := newServer(config)
	if err != nil {
		log.Fatalf("Error: can not connect to %s: %s", config.Site, err.Error())
	}
	defer myldap.Close()

	// Minimum search is uid
	attrs := map[string]bool{
		"kgivenname": true,
		"ksn": true,
		"uid": true,
	}

	// Setup searches
	if fInclMail {
		attrs["mail"] = true
	}

	// Meat of the game, the search
	res, err := myldap.doSearch(attrs, flag.Arg(0))
	if err != nil {
		log.Printf("Error: searching failed: %v", err)
	}

	for _, entry := range res {
		entry.PrettyPrint(2)
	}
	log.Printf("Found %d results\n", len(res))

	// We're done
	log.Printf("Shutting down…")
}

