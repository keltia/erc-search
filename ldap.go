// ldap.go
//

/*
    Package to implement a thin layer over the ldap server
 */
package main

import (

	"github.com/go-ldap/ldap"
	"github.com/keltia/erc-search/config"
	"fmt"
	"log"
	"errors"
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

	// Even in non-verbose, display something
	log.Printf("Connecting to %s\n", connstr)

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