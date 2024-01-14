package cmd

import (
	"fmt"
	"github.com/adel-hadadi/gotisan/config"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var framework string = ""

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "to initialize Gotisan config and templates",
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			return
		}

		cfg, err := config.NewDefaultConfig(framework)
		if err != nil {
			log.Fatal(err)
			return
		}

		if err := os.MkdirAll(wd+"/.gotisan/templates", os.ModePerm); err != nil {
			log.Fatal(err)
			return
		}

		err = copyTemplateFiles(cfg)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("Lets hack Nasa 🙂")
	},
}

func init() {
	initCmd.Flags().StringVar(
		&framework,
		"framework",
		"",
		"create handlers with framework context",
	)

	rootCmd.AddCommand(initCmd)
}
