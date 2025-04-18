package configuration

import (
	craveConfiguration "crave/shared/configuration"
)

type Secret struct {
	Controller string
}

type Dependency struct {
}

type Variable struct {
	Database *craveConfiguration.Database
	Secret   *Secret
}

func NewVariable() *Variable {
	return &Variable{
		Database: &craveConfiguration.Database{
			Uri:      "127.0.0.1:18000",
			Username: "root",
			Password: "root",
		},
		Secret: &Secret{
			Controller: "secretValue",
		},
	}
}
