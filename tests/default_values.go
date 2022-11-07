package main

type Values struct {
	DefaultImage           string      `json:"defaultImage"`
	DefaultImageTag        string      `json:"defaultImageTag"`
	DefaultImagePullPolicy string      `json:"defaultImagePullPolicy"`
	Deployments            Deployments `json:"deployments"`
}

type Deployments map[string]*Deployment

type Deployment struct {
	Replicas           int        `json:"replicas"`
	ServiceAccountName string     `json:"serviceAccountName"`
	Containers         Containers `json:"containers"`
}

type Containers []struct {
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
				Replicas:           2,
				ServiceAccountName: "deployer",
				Containers: Containers{
					{
						Name:            "nginx",
						Image:           "nginx",
						ImageTag:        1.19,
						ImagePullPolicy: "Always",
					},
				},
			},
		},
	}

	return values
}
