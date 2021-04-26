package project

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Project struct {
	name  string
	Disks []Disk
}

type Disk struct {
	Name  string
	File  string
	Size  string
	Files []File
}

type File struct {
	Input  string
	Output string
}

var AppyProj Project

func SelfConfigure() {
	AppyProj.name = "Selfie"
	err := viper.Unmarshal(&AppyProj)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

func Id() {
	fmt.Println("HI FROM:", AppyProj.name)
}
