package main

import (
	"flag"
	"fmt"
	"github.com/xor-gate/debpkg"
)

var gOutputFile string
var gConfigFile string

func init() {
	flag.StringVar(&gConfigFile, "c", "", "Yaml configuration file")
	flag.StringVar(&gOutputFile, "o", "", "Debian output file")
	flag.Parse()
}

func main() {
	deb := debpkg.New()

	if gConfigFile != "" {
		err := deb.Config(gConfigFile)
		if err != nil {
			fmt.Println("Error while loading config file", gConfigFile)
			return
		}
	}

	err := deb.Write(gOutputFile)
	if err != nil {
		fmt.Println("debpkg: error:", err)
		return
	}

	fmt.Println("debpkg: written:", gOutputFile)
}
