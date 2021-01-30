package test

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Get("/",func(ctx iris.Context){
		ctx.HTML("<b>Hello!</b>")
	})
	app.Run(iris.Addr(":8882"),iris.WithConfiguration(iris.YAML("./config/aa.yml")))
}