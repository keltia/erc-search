package ldap

import (
	"errors"
	"net"
	"fmt"
)

var (
	ErrEmptyDomain = errors.New("empty domain name")
	ErrBadSRVRecord = errors.New("bad/empty SRV record")
)

// GetServerName returns one server amongst all SRV records
func GetServerName(domain string) (srv string, err error) {
	if domain == "" {
		srv = domain
		err = ErrEmptyDomain
	}

	// Get the actual SRV records
	_, addrs, err := net.LookupSRV("ldap", "tcp", domain)
	if err != nil {
		srv = ""
		err = ErrBadSRVRecord
		return
	}

	fmt.Printf("srv: %+v\n", addrs)
	// We suppose the nameserver does the randomize itself
	srv = addrs[0].Target
	return
}
