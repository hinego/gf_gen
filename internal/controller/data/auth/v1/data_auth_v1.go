package data

import (
	"github.com/sucold/starter/internal/controller"
)

func init() {
	controller.Register("dataAuthV1", &controllerV1{})
}

type controllerV1 struct{}
