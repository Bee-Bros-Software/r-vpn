package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Bee-Bros-Software/r-vpn/version"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the R-VPN's client application version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.SetOut(cmd.OutOrStdout())
			cmd.Println(version.NetbirdVersion())
		},
	}
)
