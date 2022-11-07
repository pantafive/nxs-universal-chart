package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
	"os"
	"path/filepath"
	"testing"
)

func TestExperimental(t *testing.T) {
	root, _ := os.Getwd()
	schemaPath := filepath.Join(root, "../schema.json")
	validValues := mustReadYAML(filepath.Join(root, "../samples/schema-demo-valid.values.yaml"))
	invalidValues := mustReadYAML(filepath.Join(root, "../samples/schema-demo-invalid.values.yaml"))

	schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", schemaPath))
	validLoader := gojsonschema.NewRawLoader(validValues)
	invalidLoader := gojsonschema.NewRawLoader(invalidValues)

	t.Run("validate Valid Schema", func(t *testing.T) {
		result, err := gojsonschema.Validate(schemaLoader, validLoader)
		assert.NoError(t, err)
		assert.Truef(t, result.Valid(), "Schema validation failed: %s", result.Errors())
	})

	t.Run("validate Invalid Schema", func(t *testing.T) {
		result, err := gojsonschema.Validate(schemaLoader, invalidLoader)
		assert.NoError(t, err)
		assert.Falsef(t, result.Valid(), "Schema validation failed: %s", result.Errors())
	})

	t.Run("validate Default Values Go Struct", func(t *testing.T) {
		result, err := gojsonschema.Validate(schemaLoader, gojsonschema.NewGoLoader(DefaultValues()))
		assert.NoError(t, err)
		assert.Truef(t, result.Valid(), "Schema validation failed: %s", result.Errors())
	})

	t.Run("validate ImagePullPolicy", func(t *testing.T) {
		values := DefaultValues()
		result, err := gojsonschema.Validate(schemaLoader, gojsonschema.NewGoLoader(values))
		assert.NoError(t, err)
		assert.Truef(t, result.Valid(), "Schema validation failed: %s", result.Errors())
	})

	t.Run("Main", func(t *testing.T) {

	})

}
