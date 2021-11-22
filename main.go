package main

import (
	"fmt"
	"log"

	"github.com/duvalhub/cloudconfigclient"
)

func main() {
	configsEnv := cloudconfigclient.ConfigsEnv{}
	config, err := configsEnv.ReadFromEnv().Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	filesToChange, err := filesToChange(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range filesToChange {
		file.Apply()
	}

	log.Print("Done. Files processed.")
}
