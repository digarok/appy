package project

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Project struct {
	name          string
	Disks         []Disk
	Assemble      []string
	Indent        []string
	AssembleFlags string
	FormatFlags   string
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

type LocalConfig struct {
	Programs Programs `mapstructure:"programs"`
}

type Programs struct {
	Merlin32 string
	Cadius   string
	Gsplus   string
}

const Merlin32Path = "merlin32"
const CadiusPath = "cadius"
const GsplusPath = "gsplus"

var AppyProj Project
var LocalConf LocalConfig

func SelfConfigure() {
	AppyProj.name = "Default"
	AppyProj.AssembleFlags = ""
	LocalConf.Programs.Merlin32 = Merlin32Path
	LocalConf.Programs.Cadius = CadiusPath
	LocalConf.Programs.Gsplus = GsplusPath

	// first do main configuration
	err := viper.Unmarshal(&AppyProj)
	if err != nil {
		log.Fatalf("unable to decode MAIN config into struct, %v", err)
	}
	handleLocalConfigs()
}

func handleLocalConfigs() {
	// @todo: also allow flag for filename
	localConfigFile := "appy.user.yaml"
	if exists(localConfigFile) {
		v := viper.New()
		v.SetConfigName(localConfigFile)
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		err := v.ReadInConfig() // Find and read the config file
		if err != nil {         // Handle errors reading the config file
			log.Fatalf("unable to read LOCAL config, %v", err)
		}

		err = v.Unmarshal(&LocalConf)
		if err != nil {
			panic("Unable to unmarshal local")
		}

		fmt.Println("Loaded local conf")
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func Id() {
	fmt.Println("HI FROM:", AppyProj.name)
}
