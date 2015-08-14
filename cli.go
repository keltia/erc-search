// cli.go
//
// Command-line flags handling

/*
  Implements the basic command-line flag handling
 */

package main

import (
	"flag"
)

var (
	fInclFull bool
	fInclMail bool
)

func init() {
	flag.BoolVar(&fInclFull, "F", false, "Include full name search")
	flag.BoolVar(&fInclMail, "M", false, "Include mail search")
}
