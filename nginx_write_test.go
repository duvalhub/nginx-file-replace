package main

import (
	"bytes"
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

	mockTemplate := "nginx.tmpl"

	tests := []struct {
		name     string
		expected string
		request  NginxConfigWriter
	}{
		{
			"allo",
			"nginx_write_noproxy.txt",
			NginxConfigWriter{
				map[string]Proxy{},
				tmpOutputFile.Name(),
				mockTemplate,
			},
		},
		{
			"allo",
			"nginx_write_proxies.txt",
			NginxConfigWriter{
				map[string]Proxy{
					"aa": {
						"toto.com",
						8080,
						"toto",
					},
					"bb": {
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

			expected, err := os.ReadFile(tt.expected)
			if err != nil {
				log.Fatal(err)
			}

			output, err := os.ReadFile(tt.request.Output)
			if err != nil {
				log.Fatal(err)
			}
			if !bytes.Equal(expected, output) {
				t.Errorf("NginxConfigWriter.Write() = \n%v\n Want \n%v\n", string(output), string(expected))
			}

		})
	}
}
