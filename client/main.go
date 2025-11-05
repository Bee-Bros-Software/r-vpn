package main

import (
	"os"

	"github.com/Bee-Bros-Software/r-vpn/client/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
