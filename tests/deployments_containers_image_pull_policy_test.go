package tests

import (
	"app/tests/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"
	"testing"
)

func Test_Deployments_Containers_ImagePullPolicy(t *testing.T) {
	schema := utils.SchemaLoader()

	tests := []struct {
		given   string
		isValid bool
	}{
		// valid
		{given: "Always", isValid: true},
		{given: "IfNotPresent", isValid: true},
		{given: "Never", isValid: true},
		// invalid
		{given: "SomeTimes", isValid: false},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.given, func(t *testing.T) {
			values := utils.DefaultValues()
			values.Deployments["nginx"].Containers[0].ImagePullPolicy = tt.given

			result, err := gojsonschema.Validate(schema, gojsonschema.NewGoLoader(values))

			require.NoError(t, err)
			assert.Equalf(t, tt.isValid, result.Valid(), "Schema validation failed: %s", result.Errors())
		})
	}
}
