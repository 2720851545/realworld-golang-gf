package main

import (
	_ "github.com/2720851545/realworld-golang-gf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/2720851545/realworld-golang-gf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
