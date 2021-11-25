package main

import (
	"html/template"
	"log"
	"os"
	"path"

	"github.com/duvalhub/cloudconfigclient"
	"github.com/mitchellh/mapstructure"
)

type Proxy struct {
	Url  string
	Port int
	Path string
}

type NginxConfigWriter struct {
	Proxies  map[string]Proxy
	Output   string
	Template string
}

func (n *NginxConfigWriter) Write() {
	writer, err := os.Create(n.Output)
	if err != nil {
		log.Fatalf("Couldn't open configuration file '%s'. Error: %s", n.Output, err)
	}

	file, err := os.Open(n.Template)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	tmpl, err := template.New(path.Base(n.Template)).ParseFiles(n.Template)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(writer, n)
	if err != nil {
		log.Fatal(err)
	}
}

func newNginxConfigWriterFromConfig(source cloudconfigclient.Source) (NginxConfigWriter, error) {
	nginxConfigWriter := &NginxConfigWriter{}
	json := source.Data["nginx"]
	mapstructure.Decode(json, &nginxConfigWriter)
	return *nginxConfigWriter, nil
}
