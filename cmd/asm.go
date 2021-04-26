package cmd

import (
	"github.com/digarok/appy/core"

	"github.com/spf13/cobra"
)

// 'appy asm' command
var asmCmd = &cobra.Command{
	Use:   "asm",
	Short: "Assemble the source files in your project",
	Long: `This will run your 6502/65816 assembler, such as Merlin32,
against all of the files specified for assembly in your project.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Assemble()
	},
}

func init() {
	rootCmd.AddCommand(asmCmd)
}
