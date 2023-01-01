package core

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/digarok/appy/core/project"
	"github.com/fatih/color"
)

func CreateDisk(name string, file string, size string) {
	fmt.Printf("Creating Disk: \"%s\" -> %s \tSize: %s\n", name, file, size)

	out, err := exec.Command(project.LocalConf.Programs.Cadius, "CREATEVOLUME", file, name, size).Output()
	if err != nil {
		color.Cyan(string(out))
		log.Fatal(err)
	}
}

func AddFiles(disk project.Disk) {
	fmt.Printf("Add files to:  \"%s\"\n", disk.Name)
	for _, file := range disk.Files {
		fmt.Printf(" Adding file: ----->  %s\n", file.Input)

		out, err := exec.Command(project.LocalConf.Programs.Cadius, "ADDFILE", disk.File, file.Output, file.Input).Output()
		if err != nil {
			color.Cyan(string(out))
			log.Fatal(err)
		}
	}
}

func BuildDisk() {
	for _, disk := range project.AppyProj.Disks {
		CreateDisk(disk.Name, disk.File, disk.Size)
		AddFiles(disk)
	}
}
