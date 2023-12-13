/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/adel-hadadi/tisan/commands/handler"
	"strings"

	"github.com/spf13/cobra"
)

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "make:handler",
	Short: "make handler",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		hcmd := handler.HandlerCommand{}

		hcmd.Make(args[0], cfgFile)

		fmt.Println(strings.Join(args, ", "))
		fmt.Println("make something :/")
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
