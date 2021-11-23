package main

import (
	"log"

	"github.com/duvalhub/cloudconfigclient"
)

func main() {
	configsEnv := cloudconfigclient.ConfigsEnv{}
	config, err := configsEnv.ReadFromEnv().Load()
	if err != nil {
		log.Fatal(err)
	}

	// Change Static Files
	changeFiles(config)

	// Generate Nginx Conf
	generateNginxconf(config)
}

func changeFiles(config cloudconfigclient.Source) {
	filesToChange, err := filesToChange(config)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range filesToChange {
		file.Apply()
	}

	log.Print("Done. Files processed.")
}

func generateNginxconf(config cloudconfigclient.Source) {
	nginxConfigWriter, err := newNginxConfigWriterFromConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	nginxConfigWriter.Write()

}
