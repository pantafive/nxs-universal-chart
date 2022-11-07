package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/sjson"
	"github.com/xeipuuv/gojsonschema"

	"github.com/nixys/nxs-universal-chart/tests/utils"
)

func TestForbidAdditionalProperties(t *testing.T) {
	schema := utils.SchemaLoader()
	values := utils.DefaultValues()
	valuesAsJSON := values.AsJSON(t)

	tests := []struct {
		jsonPath string
		key      string
		val      interface{}
	}{
		{jsonPath: "deployments.nginx.containers.0", key: "pullPolicy", val: "Always"},
	}
	for i := range tests {
		tt := tests[i]
		t.Run("deployments.nginx.containers", func(t *testing.T) {
			valuesWithAdditionalProperty, err := sjson.Set(valuesAsJSON, tt.jsonPath+"."+tt.key, tt.val)
			require.NoError(t, err)

			result, err := gojsonschema.Validate(schema, gojsonschema.NewStringLoader(valuesWithAdditionalProperty))
			require.NoError(t, err)
			assert.Equalf(t, false, result.Valid(), "Schema validation failed: %s", result.Errors())
			assert.Equal(t, fmt.Sprintf("%s: Additional property %s is not allowed", tt.jsonPath, tt.key), result.Errors()[0].String())
		})
	}
}
