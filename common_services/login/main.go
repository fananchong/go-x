package main

import "fmt"

func init() {
	fmt.Println("LOAD PLUGIN : LOGIN")
}

func Run() {
	app := NewApp()
	app.Run()
}
