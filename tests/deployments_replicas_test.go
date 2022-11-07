package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"
	"testing"
)

func Test_Deployments_Replicas(t *testing.T) {
	schema := SchemaLoader()

	tests := []struct {
		given   int
		isValid bool
	}{
		// valid
		{given: 0, isValid: true},
		{given: 1, isValid: true},
		{given: 2, isValid: true},
		{given: 3, isValid: true},
		// invalid
		{given: -1, isValid: false},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			values := DefaultValues()
			values.Deployments["nginx"].Replicas = tt.given

			result, err := gojsonschema.Validate(schema, gojsonschema.NewGoLoader(values))

			require.NoError(t, err)
			assert.Equalf(t, tt.isValid, result.Valid(), "Schema validation failed: %s", result.Errors())
		})
	}
}
