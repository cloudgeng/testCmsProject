package test

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Get("/",before,mainHandler,after)
	app.Run(iris.Addr(":8081"))

}

func before(ctx iris.Context) {
	shareInformation := "this is a sharable information between handlers"
	requestPath := ctx.Path()
	println("before the mainHandler: " + requestPath)

	ctx.Values().Set("info",shareInformation)
	ctx.Next() //执行下一个处理器
}

func after(ctx iris.Context){
	println("After the mainHandler")
}

func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")

	info := ctx.Values().GetString("info")

	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() //execute the "afer"
}