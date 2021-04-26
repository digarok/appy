package core

import (
	"fmt"
	"log"
	"os/exec"
)

const emulatorPath = "gsplus"

func Run() {
	fmt.Println("Running an emulator")

	cmd := exec.Command(emulatorPath)

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

}
