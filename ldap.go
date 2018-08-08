// ldap.go
//

/*
Package ldap implements a thin layer over the ldap server
*/
package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"gopkg.in/ldap.v2"
)

// Server wraps the ldap parameters
type Server struct {
	c *ldap.Conn
	s *Source
}

// NewServer creates a new client instance
func NewServer(src *Source) (srv *Server, err error) {
	verbose("Adding %v as source", src)

	// Get one of the SRV records if .Site is empty
	if src.Site == "" {
		rec, lerr := GetServerName(src.Domain)
		if lerr != nil {
			return &Server{}, errors.Wrapf(err, "%+v - srv %+v\n", lerr, rec)
		}
		src.Site = rec
	}

	// Connect to the server
	c, err := doConnect(src.Site, src.Port)
	if err != nil {
		return nil, errors.Wrapf(err, "connection failed to %s", src.Site)
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
	verbose("Connecting to %s\n", connstr)

	// Connect
	c, err := ldap.Dial("tcp", connstr)
	if err != nil {
		return c, errors.Wrapf(err, "can not connect to %s\n", site)
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
		return nil, errors.Wrapf(err, "warning searching %s failed\n", filter)
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
	debug("allResults=%v", allResults)
	return allResults, nil
}
