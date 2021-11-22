package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	// Arrange
	tmpFile, err := ioutil.TempFile("", "test-write")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	source, err := os.Open("testdata/write_input")
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	nBytes, err := io.Copy(tmpFile, source)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Copied '%d' bytes", nBytes)

	fileToChange := FileToChange{tmpFile.Name(), map[string]Replacement{
		"url": {
			"CONFIG_URL",
			"http://myurl.com",
		},
	}}

	// Act
	fileToChange.Apply()

	// Assert
	fmt.Println(fileToChange)
	expected, err := os.ReadFile("testdata/write_expected")
	if err != nil {
		log.Fatal(err)
	}
	output, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	if !bytes.Equal(expected, output) {
		t.Errorf("fileToChange.Apply() = \n%v\n Want \n%v\n", string(output), string(expected))
	}
}
