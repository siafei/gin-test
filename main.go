package main

import "github/siafei/gin-test/bootstrap"

func main() {
	app := bootstrap.NewApp()
	app.Run()
}
