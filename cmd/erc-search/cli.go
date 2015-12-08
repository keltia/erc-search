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
	fVerbose  bool
)

func init() {
	flag.BoolVar(&fInclFull, "F", false, "Include full name search")
	flag.BoolVar(&fInclMail, "M", false, "Include mail search")
	flag.BoolVar(&fVerbose, "v", false, "Be verbose")
}
