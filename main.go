package main

import "github.com/kataras/iris"

func main() {
	iris.Get("/", hi)
	iris.Listen(":9999")
}

func hi(ctx *iris.Context){
	ctx.Write("Hi %s", "Andrej.")
}