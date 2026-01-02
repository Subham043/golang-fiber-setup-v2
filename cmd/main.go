package main

import (
	"github.com/subham043/golang-fiber-setup/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module()).Run()
}
