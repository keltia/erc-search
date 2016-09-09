package main

import (
	"errors"
	"fmt"
	"net"
	"os"
)

var (
	ErrEmptyDomain  = errors.New("empty domain name")
	ErrBadSRVRecord = errors.New("bad/empty SRV record")
)

// GetServerName returns one server amongst all SRV records
func GetServerName(domain string) (srv string, err error) {
	if domain == "" {
		srv = domain
		err = ErrEmptyDomain
		return
	}

	nss, err := net.LookupNS(domain)
	fmt.Printf("nss: %v\n", nss)
	if err != nil {
		srv = ""
		return
	}

	// Get the actual SRV records
	cname, addrs, err := net.LookupSRV("ldap", "tcp", domain)
	fmt.Printf("cname: %+v srv: %+v\n", cname, addrs)
	if err != nil {
		srv = ""
		return
	}

	// We suppose the nameserver does the randomize itself
	srv = addrs[0].Target
	return
}

func main() {
	fmt.Printf("ping/foobar\n")
	srv, err := GetServerName("sky.eurocontrol.int.")
	if err != nil {
		fmt.Printf("%+v - srv %+v\n", err, srv)
		os.Exit(2)
	}
	fmt.Printf("%s\n", srv)
}
