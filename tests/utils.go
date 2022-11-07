package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func SchemaLoader() gojsonschema.JSONLoader {
	root, _ := os.Getwd()
	schemaPath := filepath.Join(root, "../schema.json")
	schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", schemaPath))
	return schemaLoader
}

func mustReadYAML(path string) map[string]interface{} {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data map[string]interface{}
	err = yaml.NewDecoder(file).Decode(&data)
	if err != nil {
		panic(err)
	}

	return data
}
