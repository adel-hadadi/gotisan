package cmd

import (
	"github.com/adel-hadadi/gotisan/services"
	"github.com/spf13/cobra"
)

var (
	isRestful   bool = false
	createModel bool = false
)

var handlerCmd = &cobra.Command{
	Use:   "make:handler [name]",
	Short: "make handler",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		service := services.HandlerCommand{
			Destination: args[0],
			Options: services.HandlerOptions{
				IsRestful:   isRestful,
				CreateModel: createModel,
			},
		}

		service.Make()
	},
}

func init() {
	handlerCmd.Flags().BoolVarP(
		&isRestful,
		"restful",
		"r",
		false,
		"create restfull handler",
	)

	rootCmd.AddCommand(handlerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// handlerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// handlerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
