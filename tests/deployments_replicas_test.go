package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"

	"github.com/nixys/nxs-universal-chart/tests/utils"
)

func Test_Deployments_Replicas(t *testing.T) {
	schema := utils.SchemaLoader()

	tests := []struct {
		given   interface{}
		isValid bool
	}{
		// valid
		{given: 0, isValid: true},
		{given: 1, isValid: true},
		{given: 2, isValid: true},
		{given: 3, isValid: true},
		// invalid
		{given: nil, isValid: false},
		{given: "", isValid: false},
		{given: "1", isValid: false},
		{given: -1, isValid: false},
		{given: 0.1, isValid: false},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(fmt.Sprintf("%T:%#v", tt.given, tt.given), func(t *testing.T) {
			values := utils.DefaultValues()
			values.Deployments["nginx"].Replicas = tt.given

			result, err := gojsonschema.Validate(schema, gojsonschema.NewGoLoader(values))

			require.NoError(t, err)
			assert.Equalf(t, tt.isValid, result.Valid(), "Schema validation failed: %s", result.Errors())
		})
	}
}
