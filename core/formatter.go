package core

import (
	"fmt"

	"github.com/digarok/appy/core/project"
	"github.com/digarok/merlingo"
)

func Format(args []string) {
	if len(args) == 0 {
		// format all assembly files in appy.yaml
		for _, filename := range project.AppyProj.Assemble {
			fmt.Printf("Formatting %v\n", filename)
			merlingo.FmtFile(filename)
		}
	} else {
		// format all assembly files in args
		for _, filename := range args {
			fmt.Printf("Formatting %v\n", filename)
			merlingo.FmtFile(filename)
		}
	}
}
