package main

import "os"

type Config struct {
	FileLocation string
}

func CollectConfig() (config Config) {

	// DB_ENGINE
	fileLocation := os.Getenv("FILE_LOCATION")
	if fileLocation == "" {
		config.FileLocation = "relay.jsonld"
	} else {
		config.FileLocation = fileLocation
	}

	return
}
