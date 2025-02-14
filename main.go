package main

import (
	"encoding/json"
	"os"
	"prime-gateway/config"
)

func main() {

	// os.Args is a slice of strings containing all arguments
	args := os.Args
	lengthOfCliArgs := len(args)
	if lengthOfCliArgs < 2 {
		panic("Error!, File Path was not Defined")
	}

	// The first argument (args[0]) is the program name

	//Second argument is meant to be the path to the JSON file config
	filePath := args[1]

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		panic("Error Reading File, Ensure the file path is correct")
	}

	var userConfig map[string]interface{}

	err = json.Unmarshal(fileData, &userConfig)
	if err != nil {
		panic("Error parsing JSON, Check your JSON File")
	}

	config.LoadConfig(userConfig)

}
