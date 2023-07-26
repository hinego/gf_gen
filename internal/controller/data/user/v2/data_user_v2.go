package data

import (
	"github.com/sucold/starter/internal/controller"
)

func init() {
	controller.Register("dataUserV2", &controllerV2{})
}

type controllerV2 struct{}
