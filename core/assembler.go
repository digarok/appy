package core

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/digarok/appy/core/project"
	"github.com/fatih/color"
)

func Assemble() {
	// assemble all files in list
	for _, filename := range project.AppyProj.Assemble {
		fmt.Printf("Assembling %v\n", filename)
		fmt.Printf(">> %v %v %v\n", project.LocalConf.Programs.Merlin32, project.AppyProj.AssembleFlags, filename)

		out, err := exec.Command(project.LocalConf.Programs.Merlin32, project.AppyProj.AssembleFlags, filename).Output()
		if err != nil {
			color.Cyan(string(out))
			log.Fatal(err)
		}
	}
}
