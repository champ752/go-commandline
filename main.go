package main

import (
	"goCmd/handler"
	"goCmd/util"
)

func main() {
	utility := util.Utility{}
	commandLineHandler := handler.CreateCommandLineHandler(utility)
	commandLineHandler.Start()
}
