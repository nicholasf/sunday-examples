package main

import (
    "github.com/nicholasf/sunday"
)

func main() {
    routes := sunday.NewRoutes()
    routes.Get("/hello", HelloWorldCreatorFunc)
    sunday.Run(routes)
}

func HelloWorldCreatorFunc() (c sunday.Controller, e error) {
    c = &HelloWorldController{}
    return
}

type HelloWorldController struct {
    sunday.DefaultController
}

func (h HelloWorldController) Do(r sunday.Request) (v sunday.View, e error) {
    v = sunday.ViewFunc(HelloWorld)
    return
}

func HelloWorld(r sunday.Request) (resp sunday.Response, e error) {
    resp = sunday.NewResponse()
    resp.SetData([]byte("Hello World!"))
    return 
}
