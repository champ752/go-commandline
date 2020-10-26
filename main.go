package main

import "goCmd/handler"

func main() {
	commandLineHandler := handler.CommandLineHandler{}
	commandLineHandler.Start()
}
