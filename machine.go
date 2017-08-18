package main

import (
    "log"
    myldap "github.com/keltia/erc-search/lib"
)

// searchForMachine looks into AD for computers
func searchForMachine(ctx *context, name string) {
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


