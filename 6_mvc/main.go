package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	//mvc.Configure(app.Party("/user"),)
	//设置自定义控制器
	mvc.New(app).Handle(new(UserController))

	//路由组的mvc处理
	mvc.Configure(app.Party("/user"),func(context *mvc.Application) {
		context.Handle(new(UserController))
	})
	app.Run(iris.Addr(":8008"))
}

func (uc *UserController) Get() string {
	iris.New().Logger().Info("Get 请求 ")
	return "hello world"
}

func (uc *UserController) Post() {
	iris.New().Logger().Info(" post 请求 ")
}

func (uc *UserController) Put() {
	iris.New().Logger().Info(" put 请求")
}

func (uc *UserController) Getinfo() mvc.Result {
	iris.New().Logger().Info(" get 请求，请求路径为info")
	return mvc.Response{
		Object: map[string]interface{} {
			"code":	1,
			"message": "请求成功",
		},
	}
}
type UserController struct {

}

func (uc *UserController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET","/query","UserInfo")
}

func (uc *UserController) UserInfo() mvc.Result {
	//todo

	iris.New().Logger().Info(" user info query")
	return mvc.Response{}
}
