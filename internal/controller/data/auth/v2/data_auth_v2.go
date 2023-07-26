package data

import (
	"github.com/sucold/starter/internal/controller"
)

func init() {
	controller.Register("dataAuthV2", &controllerV2{})
}

type controllerV2 struct{}
