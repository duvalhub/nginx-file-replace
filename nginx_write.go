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
	Proxies  []Proxy
	output   string
	template string
}

func (n *NginxConfigWriter) Write() {
	writer, err := os.Create(n.output)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(n.template)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	tmpl, err := template.New(path.Base(n.template)).ParseFiles(n.template)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(writer, n)
	if err != nil {
		log.Fatal(err)
	}
}

func newNginxConfigWriterFromConfig(source cloudconfigclient.Source) (NginxConfigWriter, error) {
	nginxConfigWriter := NginxConfigWriter{}
	json := source.Data["proxies"]
	mapstructure.Decode(json, &nginxConfigWriter)
	return nginxConfigWriter, nil
}
