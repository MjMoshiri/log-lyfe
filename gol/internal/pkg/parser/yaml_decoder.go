package parser

import (
	"gopkg.in/yaml.v3"
	"os"
)

// FromYAML parses a yaml file and unmarshals it into the out interface
func FromYAML(filename string, out interface{}) error {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(fileContent, out)
	if err != nil {
		return err
	}
	return nil
}
