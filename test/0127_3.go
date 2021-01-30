package test

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.Use(before)
	app.Done(after)

	app.Get("/",indexHandler)
	app.Get("/contact",contactHandler)

	app.Run(iris.Addr(":8080"))
}

func before(ctx iris.Context) {

}

func after(ctx iris.Context) {

}

func indexHandler(ctx iris.Context) {
	ctx.HTML("<h1>Index</h1>")
	ctx.Next()
}

func contactHandler(ctx iris.Context) {
	ctx.HTML("<h1>Contact</h1>")
	ctx.Next()
}
