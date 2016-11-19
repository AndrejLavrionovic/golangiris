package main

import (
	"net"
	"github.com/kataras/iris"
)

func main() {
	iris.Get("/", func(ctx *iris.Context){
		ctx.Write(ctx.Request.URI().String())
	})

	iris.Get("/server1", func(ctx *iris.Context){
		ctx.Render("hi.html", struct{Name string}{Name: "Welcom"})
	})
	iris.Get("/server2", server2)

	ln, err := net.Listen("tcp4", "127.0.0.1:9898")
	if err != nil{
		panic(err)
	}

	iris.Serve(ln)
}

func server2(ctx *iris.Context){
	ctx.Render("hi.html", struct { Name string }{ Name: "Welcome to secure server #2"})
}