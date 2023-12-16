/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/adel-hadadi/gotisan/config"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile *config.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotisan",
	Short: "an cmd tool for make programming more enjoyable",
	Long:  "\n   ______      __  _                \n  / ____/___  / /_(_)________ _____ \n / / __/ __ \\/ __/ / ___/ __ `/ __ \\\n/ /_/ / /_/ / /_/ (__  ) /_/ / / / /\n\\____/\\____/\\__/_/____/\\__,_/_/ /_/ \n                                    \n",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child services to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gotisan.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	//cfg, err := config.InitConfig()
	//cobra.CheckErr(err)
	//
	//cfgFile = cfg
}
