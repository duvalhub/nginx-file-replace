package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestNginxWrite(t *testing.T) {
	tmpOutputFile, err := os.CreateTemp("", "asd")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpOutputFile.Name())

	mockTemplate := "testdata/nginx_write_template.tmpl"

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
				tmpOutputFile.Name(),
				mockTemplate,
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
				tmpOutputFile.Name(),
				mockTemplate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.request.Write()

			expected, err := os.ReadFile(fmt.Sprintf("testdata/%s", tt.expected))
			if err != nil {
				log.Fatal(err)
			}

			output, err := os.ReadFile(tt.request.output)
			if err != nil {
				log.Fatal(err)
			}
			if !bytes.Equal(expected, output) {
				t.Errorf("NginxConfigWriter.Write() = \n%v\n Want \n%v\n", string(output), string(expected))
			}

		})
	}
}
