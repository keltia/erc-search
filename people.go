package main

import (
	"github.com/pkg/errors"
)

// searchForPeople looks into the corporate LDAP
func searchForPeople(text string) error {
	// Do the actual connect
	src := NewSource("corporate")
	verbose("Source: %v CNF: %v", src, ctx.cnf)
	server, err := NewServer(&Source{
		Domain: src.Domain,
		Site:   src.Site,
		Port:   src.Port,
		Filter: src.Filter,
		Base:   src.Base,
		Attrs:  src.Attrs,
	})
	if err != nil {
		return errors.Wrapf(err, "can not connect to %s", src.Site)
	}
	defer server.Close()

	// Minimum search is uid
	attrs := map[string]bool{
		"kgivenname": true,
		"ksn":        true,
		"uid":        true,
		"eurocontrolworkstation": true,
	}

	// Setup searches
	if fInclMail {
		attrs["mail"] = true
	}

	// Meat of the game, the search
	res, err := server.Search(text, attrs)
	if err != nil {
		return errors.Wrap(err, "search failed")
	}

	for _, entry := range res {
		entry.PrettyPrint(2)
	}
	verbose("Found %d results\n", len(res))
	return nil
}
