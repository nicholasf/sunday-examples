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

func (h HelloWorldController) Do(r sunday.Request) (m sunday.Model, v sunday.View, e error) {
    v = sunday.ViewFunc(HelloWorld)
    m = HelloWorldModel{ Msg: "Hello World!"}
    return
}

func HelloWorld(m sunday.Model) (resp sunday.Response, e error) {
    resp = sunday.NewResponse()
    hwm :=m.(HelloWorldModel)
    resp.SetData([]byte(hwm.Msg))
    return 
}

type HelloWorldModel struct {
    Msg string
}