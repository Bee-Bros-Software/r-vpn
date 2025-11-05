package main

import (
	"github.com/Bee-Bros-Software/r-vpn/management/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
