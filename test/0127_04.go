package test

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError,internalServerError)

	app.Get("/",index1)
	app.Run(iris.Addr(":8181"))
}

func notFound(ctx iris.Context) {
	ctx.View("errors/404.html")
}

func internalServerError(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong,try again")
}

func index1(ctx iris.Context) {
	ctx.View("index.html")
}
