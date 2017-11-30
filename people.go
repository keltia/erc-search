package main

import (
    "log"
    "flag"

)

// searchForPeople looks into the corporate LDAP
func searchForPeople(ctx *context, text string) {
    // Do the actual connect
    src := NewSource("corporate")
    log.Printf("Source: %v CNF: %v", src, ctx.cnf)
    server, err := NewServer(&Source{
        Domain: src.Domain,
        Site:   src.Site,
        Port:   src.Port,
        Filter: src.Filter,
        Base:   src.Base,
    })
    if err != nil {
        log.Fatalf("Error: can not connect to %s: %v", src.Site, err)
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
    res, err := server.Search(attrs, flag.Arg(0))
    if err != nil {
        log.Printf("Error: searching failed: %v", err)
    }

    for _, entry := range res {
        entry.PrettyPrint(2)
    }
    log.Printf("Found %d results\n", len(res))
}


