package test

import (
	"github.com/kataras/iris/v12"
	)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views",".html"))

	app.Get("/",func(ctx iris.Context) {
		ctx.ViewData("message","Hello world!")
		ctx.View("hello.html")
	})

	app.Get("/user/{id:long}", func(ctx iris.Context) {
		userID,_ := ctx.Params().GetInt64("id")
		ctx.Writef("User ID: %d", userID)
	})

	app.Run(iris.Addr(":8183"))
}
