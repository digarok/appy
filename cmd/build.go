package cmd

import (
	"github.com/digarok/appy/core"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "This is the equivalent of 'asm' and 'disk' in a single command.",
	Long: `This will first assemble your files and then build the disk
image(s) specified in your project file.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Assemble()
		core.BuildDisk()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
