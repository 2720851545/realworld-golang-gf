package main

import (
	_ "realworld-golang-gf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"realworld-golang-gf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
