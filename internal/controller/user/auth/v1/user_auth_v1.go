package user

import (
	"github.com/sucold/starter/internal/controller"
)

func init() {
	controller.Register("userAuthV1", &controllerV1{})
}

type controllerV1 struct{}
