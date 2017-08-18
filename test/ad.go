package main

import (
	"errors"
	"fmt"
	myldap "github.com/keltia/erc-search/lib"
	"os"
)

var (
	ErrEmptyDomain  = errors.New("empty domain name")
	ErrBadSRVRecord = errors.New("bad/empty SRV record")
)

func main() {
	srv, err := myldap.GetServerName("sky.corp.eurocontrol.int.")
	if err != nil {
		fmt.Printf("%+v - srv %+v\n", err, srv)
		os.Exit(2)
	}
	fmt.Printf("%s\n", srv)
}
