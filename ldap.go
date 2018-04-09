// ldap.go
//

/*
Package ldap implements a thin layer over the ldap server
*/
package main

import (
	"fmt"
	"log"

	"gopkg.in/ldap.v2"
)

// Server wraps the ldap parameters
type Server struct {
	c *ldap.Conn
	s *Source
}

// Source describe a given LDAP/AD server
type Source struct {
	Domain string
	Site   string
	Port   int
	Base   string
	Filter string
	Attrs  []string
}

// NewServer creates a new client instance
func NewServer(src *Source) (srv *Server, err error) {
	log.Printf("Adding %v as source", src)

	// Get one of the SRV records if .Site is empty
	if src.Site == "" {
		rec, lerr := GetServerName(src.Domain)
		if lerr != nil {
			log.Printf("%+v - srv %+v\n", lerr, rec)
			err = lerr
			return
		}
		src.Site = rec
	}

	// Connect to the server
	c, err := doConnect(src.Site, src.Port)
	if err != nil {
		return nil, err
	}
	return &Server{
		c: c,
		s: src,
	}, err
}

// Do the connection
func doConnect(site string, port int) (*ldap.Conn, error) {
	// Build connection string
	connstr := fmt.Sprintf("%s:%d", site, port)

	// Even in non-verbose, display something
	log.Printf("Connecting to %s\n", connstr)

	// Connect
	c, err := ldap.Dial("tcp", connstr)
	if err != nil {
		return c, fmt.Errorf("Error: Can't connect to %s\n", site)
	}

	return c, nil
}

// Close the connection
func (s *Server) Close() error {
	s.c.Close()
	return nil
}

// SearchAttr explores the specific attribute
func (s *Server) SearchAttr(query, attr string) (*ldap.SearchResult, error) {

	filter := fmt.Sprintf(s.s.Filter, attr, query)
	verbose("  Using %s as filter\n", filter)
	sr := ldap.NewSearchRequest(s.s.Base,
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0, 0,
		false,
		filter,
		s.s.Attrs,
		nil,
	)
	res, err := s.c.Search(sr)
	if err != nil {
		log.Printf("  Warning searching %s failed\n", filter)
		return nil, err
	}
	return res, nil
}

// Search does the actual search
func (s *Server) Search(query string, attrs map[string]bool) (map[string]ldap.Entry, error) {

	allResults := make(map[string]ldap.Entry)
	for attr, yes := range attrs {
		if yes {
			verbose("  Looking for %s in %s…\n", query, attr)
			res, err := s.SearchAttr(query, attr)
			if err != nil {
				log.Printf("Warning: search for %s failed: %v", attr, err)
			}

			// Merge entries with the previous searches ones
			verbose("  Merging %d entries…\n", len(res.Entries))
			for _, entry := range res.Entries {
				allResults[entry.GetAttributeValue("uid")] = *entry
			}
		}
	}
	return allResults, nil
}
