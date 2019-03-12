package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "https",
	Long:  "HTTP Stream CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	cobra.OnInitialize()
}
