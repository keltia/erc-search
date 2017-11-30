package test

import (
	"fmt"
	"os"
)

func main() {
	srv, err := GetServerName("sky.corp.eurocontrol.int.")
	if err != nil {
		fmt.Printf("%+v - srv %+v\n", err, srv)
		os.Exit(2)
	}
	fmt.Printf("%s\n", srv)
}
