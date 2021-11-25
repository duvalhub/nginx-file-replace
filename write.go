package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/duvalhub/cloudconfigclient"
	"github.com/mitchellh/mapstructure"
)

type Replacement struct {
	Key   string
	Value string
}
type FileToChange struct {
	Name         string
	Replacements map[string]Replacement
}

func (f *FileToChange) Apply() {
	output, err := ioutil.ReadFile(f.Name)
	if err != nil {
		log.Fatal(err)
	}
	for _, replacement := range f.Replacements {
		output = bytes.Replace(output, []byte(replacement.Key), []byte(replacement.Value), -1)
	}

	ioutil.WriteFile(f.Name, output, 0666)
}

func filesToChange(source cloudconfigclient.Source) ([]FileToChange, error) {
	var filesToChange []FileToChange
	for _, file := range source.Data["files"].(map[string]interface{}) {
		fileToChange := &FileToChange{}
		mapstructure.Decode(file, &fileToChange)
		filesToChange = append(filesToChange, *fileToChange)
	}
	return filesToChange, nil
}
