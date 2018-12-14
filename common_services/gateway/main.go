package main

import "fmt"

func init() {
	fmt.Println("LOAD PLUGIN : GATEWAY")
}

func Run() {
	app := NewApp()
	app.Run()
}
