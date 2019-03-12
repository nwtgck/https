package main

import (
	"github.com/nwtgck/https/cmd"
	"os"
)

func main () {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
