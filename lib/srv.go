package ldap

import (
	"errors"
	"net"
)

var (
	// ErrEmptyDomain is for empty/nil parameter
	ErrEmptyDomain = errors.New("empty domain name")
	// ErrBadSRVRecord is for bad results
	ErrBadSRVRecord = errors.New("bad/empty SRV record")
)

// GetServerName returns one server amongst all SRV records
func GetServerName(domain string) (srv string, err error) {

    // domain must not be empty
    if domain == "" {
        err = ErrEmptyDomain
        srv = domain
        return
    }

	// Get the actual SRV records
	_, addrs, err := net.LookupSRV("ldap", "tcp", domain)
	if err != nil {
		srv = ""
		err = ErrBadSRVRecord
		return
	}

	// We suppose the nameserver does the randomize itself
	srv = addrs[0].Target
	return
}
