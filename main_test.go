package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	fmt.Printf("AHJHHHH %s\n", os.Getenv("CONFIG_URL"))
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "allo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
