package main

import (
	"fmt"
	"github.com/gizzahub/gzh-manager-go/cmd"
	"os"
)

var version = ""

func main() {
	if err := cmd.Execute(version); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
