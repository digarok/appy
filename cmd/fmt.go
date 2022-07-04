package cmd

import (
	"github.com/digarok/appy/core"
	"github.com/spf13/cobra"
)

// 'appy fmt' command
var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Format the source files in your project",
	Long:  `This will run an internal formatter over the assembly language files in the appy.yaml 'assemble:' list, OR you can pass filename arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Format(args)
	},
}

func init() {
	rootCmd.AddCommand(fmtCmd)
}
