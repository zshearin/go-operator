package controller

import (
	"github.com/zshearin/go-operator/pkg/controller/zachresource"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, zachresource.Add)
}
