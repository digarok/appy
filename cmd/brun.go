package cmd

import (
	"github.com/digarok/appy/core"
	"github.com/spf13/cobra"
)

// brunCmd represents the brun command
var brunCmd = &cobra.Command{
	Use:   "brun",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Assemble()
		core.BuildDisk()
		core.Run()
	},
}

func init() {
	rootCmd.AddCommand(brunCmd)
}
