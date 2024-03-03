package cmd

import (
	"fmt"
	"github.com/adel-hadadi/gotisan/config"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const (
	promptSelectFrameworkMsg = "Choose one of the frameworks"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "to initialize Gotisan config and templates",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()

		prompt := promptui.Select{
			Label: promptSelectFrameworkMsg,
			Items: config.Frameworks,
		}

		_, framework, err := prompt.Run()
		if err != nil {
			log.Println(err)
			return
		}

		cfg, err := config.NewConfig(framework)
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
	rootCmd.AddCommand(initCmd)
}
