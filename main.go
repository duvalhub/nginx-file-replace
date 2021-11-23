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

	// Change Static Files
	changeFiles(config)

	// Generate Nginx Conf
	generateNginxconf(config)
}

func changeFiles(config cloudconfigclient.Source) {
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

func generateNginxconf(config cloudconfigclient.Source) {
	nginxConfigWriter, err := newNginxConfigWriterFromConfig(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	nginxConfigWriter.Write()

}
