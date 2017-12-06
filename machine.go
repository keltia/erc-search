package main

import (
    "log"
)

// searchForMachine looks into AD for computers
func searchForMachine(name string) {
    src := NewSource("ad")
    myad, err := NewServer(src)
    if err != nil {
        log.Fatalf("Error: can not connect to %s: %s", src.Site, err.Error())
    }
    defer myad.Close()
}


