/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/adel-hadadi/gotisan/services"
	"github.com/spf13/cobra"
)

var modelCmd = &cobra.Command{
	Use:   "make:model [name]",
	Short: "make model",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		service := services.ModelCommand{
			Destination: args[0],
		}

		service.Make(cfgFile)
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}
