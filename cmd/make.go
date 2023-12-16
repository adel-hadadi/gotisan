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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// handlerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// handlerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
