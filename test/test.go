package test
import "github.com/kataras/iris/v12"

func main(){
	app:= iris.New()
	app.Get("/",func(ctx iris.Context){})
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	app.Run(iris.Addr(":8686"))
}
