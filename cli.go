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
	fVersion  bool
	fWorkStation bool
)

func init() {
	flag.BoolVar(&fInclMail, "M", false, "Include mail search")
	flag.BoolVar(&fVerbose, "v", false, "Be verbose")
	flag.BoolVar(&fVersion, "V", false, "Display version and quit")
	flag.BoolVar(&fWorkStation, "w", false, "Search Workstation name")
}
