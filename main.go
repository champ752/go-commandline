package main

import (
	"goCmd/handler"
	"goCmd/usecase"
	"goCmd/util"
)

func main() {
	utility := util.CreateUtility()
	uc := usecase.CreateUsecase()
	commandLineHandler := handler.CreateCommandLineHandler(utility,uc)

	commandLineHandler.Serve()
}
