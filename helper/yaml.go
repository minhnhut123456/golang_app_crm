package helper

import (
	"os"

	"gopkg.in/yaml.v3"
)


func ReadYaml(path string, v interface{}) error{
	file, err := os.ReadFile(path)
	if(err != nil) {
		return err
	}

	err = yaml.Unmarshal(file, v)
	if(err != nil) {
		return err
	}

	return nil
}