package core

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/digarok/appy/core/project"
	"github.com/fatih/color"
)

var filesToAssemble []string

func Assemble() {

	// and assemble
	var p = project.AppyProj
	filesToAssemble := p.Assemble
	//filesToAssemble = viper.GetViper().GetStringSlice("assemble")

	for _, filename := range filesToAssemble {
		fmt.Printf("Assembling %v\n", filename)

		out, err := exec.Command(project.LocalConf.Programs.Merlin32, "-V", filename).Output()

		if err != nil {
			color.Cyan(string(out))
			log.Fatal(err)
		}
	}
}
