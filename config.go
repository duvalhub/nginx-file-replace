package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/duvalhub/cloudconfigclient"
)

type BasicAuthTransport struct {
	Username string
	Password string
}

func (bat BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s",
		base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
			bat.Username, bat.Password)))))
	return http.DefaultTransport.RoundTrip(req)
}

func (bat *BasicAuthTransport) client() *http.Client {
	return &http.Client{Transport: bat}
}

type Configs struct {
	username string
	password string
	url      string
	name     string
	profiles string
	label    string
}

func ReadFromEnv() Configs {
	return Configs{
		url:      os.Getenv("CONFIG_URL"),
		username: os.Getenv("CONFIG_USERNAME"),
		password: os.Getenv("CONFIG_PASSWORD"),
		name:     os.Getenv("APPLICATION_NAME"),
		profiles: os.Getenv("APPLICATION_PROFILES"),
		label:    os.Getenv("CONFIG_LABEL"),
	}
}

func (configs Configs) Load() (cloudconfigclient.Source, error) {
	transport := BasicAuthTransport{Username: configs.username, Password: configs.password}
	client := transport.client()
	configConf := cloudconfigclient.Local(client, configs.url)
	configClient, err := cloudconfigclient.New(configConf)

	if err != nil {
		fmt.Println(err)
		return cloudconfigclient.Source{}, err
	}

	// Retrieves the configurations from the Config Server based on the application name, active profiles and label
	source, err := configClient.GetConfiguration(configs.name, strings.Split(configs.profiles, ","), configs.label)
	if err != nil {
		fmt.Println(err)
		return cloudconfigclient.Source{}, err
	}

	return source, nil
}

func Work(s cloudconfigclient.Source) map[string]interface{} {
	result := map[string]interface{}{}
	for _, propertySource := range s.PropertySources {
		for key, value := range propertySource.Source {
			entries := strings.Split(key, ".")
			result = insertInMap(entries, value, result)
		}
	}
	return result
}
func insertInMap(s []string, value interface{}, dest map[string]interface{}) map[string]interface{} {
	keys := s[:len(s)-1]
	last := s[len(s)-1]

	curr := dest
	for _, key := range keys {
		switch curr[key].(type) {
		case nil:
			curr[key] = map[string]interface{}{}
			curr = curr[key].(map[string]interface{})
		case map[string]interface{}:
			curr = curr[key].(map[string]interface{})
		}
	}
	if curr[last] == nil {
		curr[last] = value
	}
	return dest
}
func insertInMapRecursion(s []string, value interface{}, dest map[string]interface{}) map[string]interface{} {
	key := s[0]

	switch dest[key].(type) {
	case nil:
		if len(s) == 1 {
			dest[key] = value
			return dest
		} else {
			dest[key] = map[string]interface{}{}
			return insertInMap(s[1:], value, dest[key].(map[string]interface{}))
		}
	case map[string]interface{}:
		dest[key] = insertInMap(s[1:], value, dest[key].(map[string]interface{}))
		return dest
	default:
		return dest
	}

	// if dest[key] != nil {
	// 	switch dest[key].(type) {
	// 	case map[string]interface{}:
	// 		dest[key] = insertInMap(s[1:], value, dest[key].(map[string]interface{}))
	// 		return dest
	// 	default:
	// 		return dest
	// 	}
	// }
	// if len(s) == 1 {
	// 	dest[key] = value
	// } else {
	// 	dest[key] = insertInMap(s[1:], value, map[string]interface{}{})
	// }

	// return dest
}
