package main

// Source describe a given LDAP/AD server
type Source struct {
	Domain string
	Site   string
	Port   int
	Base   string
	Filter string
	Attrs  []string
}

func NewSource(name string) *Source {
	// Do the actual connect
	if s, ok := ctx.cnf.Sources[name]; ok {
		return s
	}
	return nil
}
