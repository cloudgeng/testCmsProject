package test

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	mvc.New(app).Handle(new(ExampleController))
	app.Run(iris.Addr(":11111"))
}
	type ExampleController struct {

	}

	func (c *ExampleController) Get() mvc.Result{
		return mvc.Response{
			ContentType: "text/html",
			Text: "<h1>Welcome</h1>",
		}
	}

	func (c *ExampleController) GetPing() string {
		return "pong"
	}

	func (c *ExampleController) GetHello() interface{} {
		return map[string]string{"message":"Hello Iris!"}
	}