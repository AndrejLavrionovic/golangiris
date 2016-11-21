package main

import (
	"github.com/kataras/iris"
	"github.com/valyala/fasthttp"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func main() {
	api := iris.New()

	api.Static("/*", "./public/*", 1)

	api.Get("/", getpage)

	api.Get("/mypath", func(ctx *iris.Context){
		ctx.Write("Hello from the server on path /mypath")
	})

	api.HandleFunc("GET", "/get", myhandler)

	api.API("/users", UserApi{}, myUsersMiddleware1, myUsersMiddleware2)

	api.API("/redirect", HackerNews{}, myUsersMiddleware1, myUsersMiddleware2)

	// Handler API


	// to use a custom server you have to call .Build after
	// route, sessions, templates, websockets, ssh... before server's listen
	api.Build()

	/*
	ln, err := net.Listen("tcp4", "0.0.0.0:9999")
	if err != nil{
		panic(err)
	}

	iris.Serve(ln)
	*/

	// create our custom fasthttp server and assign the Handler/Router
	fsrv := &fasthttp.Server{Handler: api.Router}
	fsrv.ListenAndServe(":9999")
}

type page struct{
	Title string
	Host string
	JObj string
	Text string
}

type JsonObj struct{
	by string
	descendants string
	id int
	kids []int
	score int
	text string
	time int
	title string
	Type string
	url string
}

func getpage(ctx *iris.Context){

	// Retrieving json object from HackerNews
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/121003.json?print=pretty")
	if err != nil{panic(err.Error())}
	body, err := ioutil.ReadAll(resp.Body)
	//jsonstring := fmt.Sprintf("%s", body)
	var f interface{}
	json.Unmarshal(body, &f)

	m := f.(map[string]interface{})

	println("m is -> ", m)

	j := JsonObj{}
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
			if k == "title" {j.title = vv}
			if k == "type" {j.Type = vv}
			if k == "by" {j.by = vv}
			if k == "descendants" {j.descendants = vv}
			if k == "text" {j.text = vv}
			if k == "url" {j.url = vv}
		case int:
			fmt.Println(k, "is int", vv)
			if k == "score" {j.score = vv}
			if k == "time" {j.time = vv}
			if k == "id" {j.id = vv}
		case []interface{}:
			fmt.Println(k, "is an array:")
			for u := range vv {
				fmt.Println(u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle", vv)
		}
	}

	println("Print something")
	println(j.title)
	for u := range j.kids{
		fmt.Println("j.kids -> ", u)
	}
	ctx.Next()

	ctx.Render("index.html", page{j.title, ctx.HostString(), fmt.Sprintf("%s", body), j.text})
}

func myhandler(c *iris.Context){
	c.Write("From %s. Implementation of handlerFunction", c.PathString())
}

type UserApi struct{*iris.Context}
// GET /users
func (u UserApi) Get(){
	//u.Write("Get from /users")
	u.HTML(iris.StatusOK, "<h3>Get all from users</h3>")
	//u.Redirect("https://hacker-news.firebaseio.com/v0/item/121003.json?print=pretty", iris.StatusOK)
}

func myUsersMiddleware1(ctx *iris.Context){
	println("From User middleware 1")
	ctx.Next()
}

func myUsersMiddleware2(ctx *iris.Context){
	println("From User middleware 2")
	ctx.Next()
}

// Retrieving json object from HackerNews and printing it into page
// 1) initialize structure with
type HackerNews struct{*iris.Context}
func (u HackerNews) Get(){
	//u.Write("Get from /users")
	//u.HTML(iris.StatusOK, "<h3>Get all from users</h3>")
	//u.Redirect("https://hacker-news.firebaseio.com/v0/item/121003.json?print=pretty", iris.StatusOK)
	//u.Request.SetRequestURI("https://hacker-news.firebaseio.com/v0/item/121003.json?print=pretty")
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/121003.json?print=pretty")
	if err != nil{panic(err.Error())}
	body, err := ioutil.ReadAll(resp.Body)
	u.Write("%s", body)
}