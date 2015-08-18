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
	"github.com/go-ldap/ldap"
	"fmt"
)

const (
	RcFile = "erc-search"
)

// Do the actual search
func doSearch(query string) (error) {
	return nil
}

// Start here
func main () {
	config, err := config.LoadConfig(RcFile)
	if err != nil {
		log.Printf("Warning: can't load %s, using defaults\n", RcFile)
		config.SetDefaults()
	}
	flag.Parse()

	if fVerbose {
		log.Printf("Default config:\n%s", config.String())
	}

	if flag.Arg(0) == "" {
		log.Fatalln("Error: You must specify a search string")
	}

	site := fmt.Sprintf("%s:%d", config.Site, config.Port)
	c, err := ldap.Dial("tcp", site);
	if err != nil {
		log.Fatalf("Error: Can't connect to %s\n", config.Site)
	}

	err = doSearch(flag.Arg(0))
	if err != nil {
		log.Printf("Error: searching failed: %v", err)
	}
	log.Printf("Shutting down…")
	c.Close()
}

