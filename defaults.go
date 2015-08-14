// defaults.go
//
// Holds default values for when the config file is not available

/*
 This file hold various default values
 */

package main

const (
	DEF_SERVER = "ldap.eurocontrol.fr"
	DEF_PORT = 389

	DEF_BASE = "ou=eurousers,o=eurocontrol,o=ec"
	DEF_FILTER = "(&(objectclass=eurocontrolperson)(%s=*%s*))"
)

var (
	DEF_ATTRS = []string{
		"ksn",
		"kgivenname",
		"telephonenumber",
		"eurocontrolroomid",
		"mail",
		"uid",
		"eurocontrolunitcode",
		"eurocontrolgrade",
	}
)
