package test

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Get("/username/{name}",func(ctx iris.Context){
		path := ctx.Path()
		app.Logger().Info(path)
		ctx.Writef("Hello %s",ctx.Params().Get("name"))
	})

	/*app.Macros().Int.RegisterFunc("min",func(minValue int) func(string) bool {
		return func(paramValue string) bool {
			n,err := strconv.Atoi(paramValue)
			if err != nil {
				return false
			}
			return n >= minValue
		}
	})

	app.Get("/profile/{id:int min(1)}",func(ctx iris.Context) {
		id,_ := ctx.Params().GetInt("id")
		ctx.Writef("hello id: %id",id)
	})
 */
	app.Run(iris.Addr(":8080"))
}
