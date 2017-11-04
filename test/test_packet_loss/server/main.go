package main

// usage: server --log_dir=./log -stderrthreshold 0

var (
	MSG = []byte("hello")
)

func main() {
	app := NewApp()
	app.Run()
}
