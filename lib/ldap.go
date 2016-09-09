// ldap.go
//

/*
    Package to implement a thin layer over the ldap server
 */
package ldap

import (

	"github.com/go-ldap/ldap"
	"github.com/keltia/erc-search/config"
	"fmt"
	"log"
	"errors"
	"github.com/keltia/go/test/fixedbugs"
)

// Wrap the ldap parameters
type Server struct {
	c      *ldap.Conn
	Site   string
	Port   int
	Base   string
	Filter string
	Attrs  []string

	Verbose bool
}

// Create a new client instance
func NewServer(src config.Source) (srv *Server, err error) {
	log.Printf("Adding %v as source", src)
	// Connect to the server
	c, err := doConnect(src.Site, src.Port); if err != nil {
		return nil, err
	}
	return &Server{
		c: c,
		Site: src.Site,
		Port: src.Port,
		Base: src.Base,
		Filter: src.Filter,
		Attrs: src.Attrs,
	}
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

// SetVerbose sets verbose mode
func (myldap *Server) SetVerbose(v bool) {
	myldap.Verbose = v || false
}

// Close the connection
func (myldap *Server) Close() (error) {
	myldap.c.Close()
	return nil
}

// Search the specific attribute
func (myldap *Server) SearchAttr(query, attr string) (*ldap.SearchResult, error) {

	filter := fmt.Sprintf(myldap.Filter, attr, query)
	if myldap.Verbose {
		log.Printf("  Using %s as filter\n", filter)
	}
	sr := ldap.NewSearchRequest(myldap.Base,
	 	ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0, 0,
		false,
		filter,
		myldap.Attrs,
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
func (myldap *Server) Search(attrs map[string]bool, query string) (map[string]ldap.Entry, error) {

	allResults := make(map[string]ldap.Entry)
	for attr, yes := range attrs {
		if yes {
			if myldap.Verbose {
				fmt.Printf("  Looking for %s in %s…\n", query, attr)
			}
			res, err := myldap.SearchAttr(query, attr)
			if err != nil {
				log.Printf("Warning: search for %s failed: %v", attr, err)
			}

			// Merge entries with the previous searches ones
			if myldap.Verbose {
				log.Printf("  Merging %d entries…\n", len(res.Entries))
			}
			for _, entry := range res.Entries {
				if myldap.Verbose {
					entry.PrettyPrint(2)
				}
				allResults[entry.GetAttributeValue("uid")] = *entry
			}
		}
	}
	return allResults, nil
}
