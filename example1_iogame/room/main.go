package main

import "fmt"

func init() {
	fmt.Println("LOAD PLUGIN : ROOM")
}

func Run() {
	app := NewApp()
	app.Run()
}
