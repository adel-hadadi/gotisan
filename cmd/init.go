package cmd

import (
	"github.com/adel-hadadi/gotisan/config"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "to initialize gotisan config and templates",
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			return
		}

		if err := os.MkdirAll(wd+"/.gotisan/templates", os.ModePerm); err != nil {
			log.Fatal(err)
			return
		}

		err = copyTemplateFiles()
		if err != nil {
			log.Fatal(err)
			return
		}

		config.NewDefaultConfig()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
