package utils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type Values struct {
	DefaultImage           string      `json:"defaultImage"`
	DefaultImageTag        string      `json:"defaultImageTag"`
	DefaultImagePullPolicy string      `json:"defaultImagePullPolicy"`
	Deployments            Deployments `json:"deployments"`
}

type Deployments map[string]*Deployment

type Deployment struct {
	Replicas           interface{}  `json:"replicas"`
	ServiceAccountName string       `json:"serviceAccountName"`
	Containers         []Containers `json:"containers"`
}

type Containers struct {
	Name            string  `json:"name"`
	Image           string  `json:"image"`
	ImageTag        float64 `json:"imageTag"`
	ImagePullPolicy string  `json:"imagePullPolicy"`
}

// DefaultValues returns new instance of Values with valid dummy setup.
func DefaultValues() Values {
	values := Values{
		DefaultImage:           "nginx",
		DefaultImageTag:        "latest",
		DefaultImagePullPolicy: "IfNotPresent",
		Deployments: Deployments{
			"nginx": &Deployment{
				Replicas:           2, //nolint:gomnd
				ServiceAccountName: "deployer",
				Containers: []Containers{
					{
						Name:            "nginx",
						Image:           "nginx",
						ImageTag:        1.19, //nolint:gomnd
						ImagePullPolicy: "Always",
					},
				},
			},
		},
	}

	return values
}

func (v *Values) AsJSON(t *testing.T) string {
	result, err := json.Marshal(v)
	require.NoErrorf(t, err, "Failed to marshal values to JSON: %s", err)
	return string(result)
}
