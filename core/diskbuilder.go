package core

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/digarok/appy/core/project"
)

func CreateDisk(name string, file string, size string) {
	fmt.Printf("Creating Disk: \"%s\" -> %s \tSize: %s\n", name, file, size)

	cmd := exec.Command(project.LocalConf.Programs.Cadius, "CREATEVOLUME", file, name, size)

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func AddFiles(disk project.Disk) {
	fmt.Printf("Add files to:  \"%s\"\n", disk.Name)
	for _, file := range disk.Files {
		// fmt.Printf("%s ADDFILE %s %s %s\n", CadiusPath, disk.File, file.Output, file.Input)
		fmt.Printf(" Adding file: ----->  %s\n", file.Input)
		cmd := exec.Command(project.LocalConf.Programs.Cadius, "ADDFILE", disk.File, file.Output, file.Input)
		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
	}

}

func BuildDisk() {
	var p = project.AppyProj
	for _, disk := range p.Disks {
		CreateDisk(disk.Name, disk.File, disk.Size)
		AddFiles(disk)
	}
}
