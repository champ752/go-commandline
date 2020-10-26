package main

import "goCmd/handler"

func main() {
	commandLineHandler := handler.CreateCommandLineHandler()
	commandLineHandler.Start()
}
