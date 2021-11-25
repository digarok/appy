package core

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/digarok/appy/core/project"
)

func Run() {
	fmt.Println("Running an emulator")

	cmd := exec.Command(project.LocalConf.Programs.Gsplus)

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
