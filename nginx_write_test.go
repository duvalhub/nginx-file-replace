package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestNginxWrite(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		request  NginxConfigWriter
	}{
		{
			"allo",
			"nginx_write_noproxy.txt",
			NginxConfigWriter{
				make([]Proxy, 0),
			},
		},
		{
			"allo",
			"nginx_write_proxies.txt",
			NginxConfigWriter{
				[]Proxy{
					{
						"toto.com",
						8080,
						"toto",
					},
					{
						"banaa.com",
						80,
						"bana",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := ioutil.TempFile("", "test-write")
			if err != nil {
				log.Fatal(err)
			}
			defer os.Remove(tmpFile.Name())

			tt.request.Write(tmpFile)

			expected, err := os.ReadFile(fmt.Sprintf("testdata/%s", tt.expected))
			if err != nil {
				log.Fatal(err)
			}

			output, err := os.ReadFile(tmpFile.Name())
			if err != nil {
				log.Fatal(err)
			}
			if !bytes.Equal(expected, output) {
				t.Errorf("NginxConfigWriter.Write() = \n%v\n Want \n%v\n", string(output), string(expected))
			}

		})
	}
}
