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
	"errors"
)

const (
	RcFile = "erc-search"
)

// Wrap the ldap parameters
type ldapServer struct {
	c      *ldap.Conn
	site   string
	config config.Config
}

// Create a new client instance
func newServer(config *config.Config) (*ldapServer, error) {
	s := new(ldapServer)
	// Connect to the server
	c, err := doConnect(config.Site, config.Port); if err != nil {
		return s, err
	}
	s.c = c
	s.site = config.Site

	// Save config for later operations.
	s.config = *config

	return s, nil
}

// Do the connection
func doConnect(site string, port int) (*ldap.Conn, error) {
	// Build connection string
	connstr := fmt.Sprintf("%s:%d", site, port)

	if fVerbose {
		log.Printf("Connecting to %s\n", connstr)
	}
	// Connect
	c, err := ldap.Dial("tcp", connstr);
	if err != nil {
		return c, errors.New(fmt.Sprintf("Error: Can't connect to %s\n", site))
	}

	return c, nil
}

// Close the connection
func (myldap *ldapServer) Close() (error) {
	myldap.c.Close()
	return nil
}

// Search the specific attribute
func (myldap *ldapServer) searchAttr(query, attr string) (*ldap.SearchResult, error) {

	filter := fmt.Sprintf(myldap.config.Filter, attr, query)
	if fVerbose {
		log.Printf("  Using %s as filter\n", filter)
	}
	sr := ldap.NewSearchRequest(myldap.config.Base,
	 	ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0, 0,
		false,
		filter,
		myldap.config.Attrs,
		nil,
	)
	res, err := myldap.c.Search(sr)
	if err != nil {
		log.Printf("  Warning searching %s failed\n", filter)
		return nil, err
	}
	return res, nil
}

// Do the actual search
func (myldap *ldapServer) doSearch (attrs map[string]bool, query string) (map[string]ldap.Entry, error) {

	allResults := make(map[string]ldap.Entry)
	for attr, yes := range attrs {
		if yes {
			if fVerbose {
				fmt.Printf("  Looking for %s in %s…\n", query, attr)
			}
			res, err := myldap.searchAttr(query, attr)
			if err != nil {
				log.Printf("Warning: search for %s failed: %v", attr, err)
			}

			// Merge entries with the previous searches ones
			if fVerbose {
				log.Printf("  Merging %d entries…\n", len(res.Entries))
			}
			for _, entry := range res.Entries {
				if fVerbose {
					entry.PrettyPrint(2)
				}
				allResults[entry.GetAttributeValue("uid")] = *entry
			}
		}
	}
	return allResults, nil
}

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

	myldap, err := newServer(config)

	// Minimum search is uid
	attrs := map[string]bool{
		"uid": true,
	}

	// Setup searches
	if fInclFull {
		attrs["ksn"] = true
		attrs["kgivenname"] = true
	}
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
	myldap.Close()
}

