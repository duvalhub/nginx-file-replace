package main

import (
	"fmt"

	"github.com/duvalhub/cloudconfigclient"
	"github.com/mitchellh/mapstructure"
)

type Change struct {
	key   string
	value string
}

type File struct {
	name    string
	changes []Change
}

type App struct {
	hello  string
	bybye  string
	mapped string
}

func mapFromConfig(config cloudconfigclient.Source) File {
	// j :=config.PropertySources[0]
	// y, err := yaml.JSONToYAML(j)
	// if err != nil {
	// fmt.Printf("err: %v\n", err)
	// return File{}
	// }
	// fmt.Println(string(y))
	// return File{}
	// var files []File = []File
	app := App{}
	mapstructure.Decode(config, &app)

	fmt.Printf("%v\n", app)

	// aasd := config.Get("files", "notfound")
	// fmt.Printf("\n\n%v\n", aasd)
	// for _, file := range aasd {
	// 	fmt.Printf("key=%s, value=%s\n", file.key, file.value)
	// }

	return File{}
}

func main() {
	config, err := ReadFromEnv().Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", config)
	port := config.Get("server.port", "").(float64)
	fmt.Println(port)
	Work(config)
	// mapFromConfig(config)
}
