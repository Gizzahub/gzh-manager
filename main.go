package main

import (
	"fmt"
	"github.com/gizzahub/gzh-manager-go/cmd/gzh-manager"
	"os"
)

var version = ""

func main() {
	if err := gzh_manager.Execute(version); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
