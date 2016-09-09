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
	myldap "github.com/keltia/erc-search/lib"
	"log"
)

const (
	rcFile = "erc-search"
)

var (
	ctx context
)

type context struct {
	cnf     *config.Config
	verbose bool
}

func (ctx *context) NewSource(name string) (config.Source) {
	// Do the actual connect
	return ctx.cnf.Sources[name]
}

// searchForPeople looks into the corporate LDAP
func searchForPeople(ctx context, text string) {
	// Do the actual connect
	src := ctx.NewSource("corporate")
	log.Printf("Source: %v CNF: %v", src, ctx.cnf)
	server, err := myldap.NewServer(src)
	if err != nil {
		log.Fatalf("Error: can not connect to %s: %v", src.Site, err)
	}
	defer server.Close()

	if ctx.verbose {
		server.SetVerbose(true)
	}

	// Minimum search is uid
	attrs := map[string]bool{
		"kgivenname": true,
		"ksn":        true,
		"uid":        true,
		"eurocontrolworkstation": false,
	}

	// Setup searches
	if fInclMail {
		attrs["mail"] = true
	}

	if fWorkStation {
		attrs["eurocontrolworkstation"] = true
	}

	// Meat of the game, the search
	res, err := server.Search(attrs, flag.Arg(0))
	if err != nil {
		log.Printf("Error: searching failed: %v", err)
	}

	for _, entry := range res {
		entry.PrettyPrint(2)
	}
	log.Printf("Found %d results\n", len(res))
}

// searchForMachine looks into AD for computers
func searchForMachine(ctx context, name string) {
	src := ctx.NewSource("ad")
	myad, err := myldap.NewServer(src)
	if err != nil {
		log.Fatalf("Error: can not connect to %s: %s", src.Site, err.Error())
	}
	defer myad.Close()

	if ctx.verbose {
		myad.SetVerbose(true)
	}

}

// Start here
func main() {
	// Load config file if any
	cnf, err := config.LoadConfig(rcFile)
	if err != nil {
		log.Printf("Warning: can't load %s\n", rcFile)
		cnf.SetDefaults()
	}

	ctx = context{
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
