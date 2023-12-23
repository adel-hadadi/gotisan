package cmd

import (
	"github.com/adel-hadadi/gotisan/services"
	"github.com/spf13/cobra"
)

var dtoCmd = &cobra.Command{
	Use:   "make:dto [name]",
	Short: "make dto objects (request and response)",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		service := services.DTOCommand{
			Destination: args[0],
		}

		service.Make()
	},
}

func init() {
	rootCmd.AddCommand(dtoCmd)
}
