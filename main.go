package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/sucold/starter/internal/controller"
	_ "github.com/sucold/starter/internal/packed"

	"github.com/sucold/starter/internal/cmd"
)

func main() {
	controller.Data()
	cmd.Main.Run(gctx.New())
}
