package cmd

import (
	"github.com/digarok/appy/core"
	"github.com/spf13/cobra"
)

// 'appy disk' command
var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Create disk images from your project",
	Long: `This will launch a disk creation utility (CADIUS), and 
create the disk images with the files as specified by
your project configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.BuildDisk()
	},
}

func init() {
	rootCmd.AddCommand(diskCmd)
}
