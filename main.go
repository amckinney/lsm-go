package main

import (
	"github.com/amckinney/lsm-go/handler/lsm"
	"github.com/amckinney/lsm-go/server"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		lsm.Module,
		server.Module,
	).Run()
}
