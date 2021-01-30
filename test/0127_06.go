package test

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	tmpl := iris.HTML("./templates",".html")
	tmpl.AddFunc("greet",func(s string) string {
		return "Greetings " + s + "!"
	})

	app.RegisterView(tmpl)

	app.Get("/",hi)
	app.Run(iris.Addr(":8184"))


}


func hi(ctx iris.Context) {
	ctx.View("hi.html")
}
