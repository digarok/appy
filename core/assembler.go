package core

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/viper"
)

var filesToAssemble []string

const Merlin32Path = "/usr/local/bin/merlin32"

func Assemble() {
	//fmt.Println("asm called")
	filesToAssemble = viper.GetViper().GetStringSlice("assemble")

	//fmt.Fprintln(os.Stderr, "HEY", filesToAssemble)

	for _, filename := range filesToAssemble {
		fmt.Printf("Assembling %v\n", filename)

		cmd := exec.Command(Merlin32Path, "-V", filename)

		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
	}
}
