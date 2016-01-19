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
	fInclMail bool
	fVerbose  bool
)

func init() {
	flag.BoolVar(&fInclMail, "M", false, "Include mail search")
	flag.BoolVar(&fVerbose, "v", false, "Be verbose")
}
