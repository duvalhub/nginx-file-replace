package main

import (
	"html/template"
	"io"
	"log"
	"os"
)

type Proxy struct {
	Url  string
	Port int
	Path string
}
type NginxConfigWriter struct {
	Proxies []Proxy
}

func (n *NginxConfigWriter) Write(writer io.Writer) {
	file, err := os.Open("nginx.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	tmpl, err := template.New(file.Name()).ParseFiles(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(writer, n)
	if err != nil {
		log.Fatal(err)
	}
}
