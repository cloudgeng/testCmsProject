package test

import (
	"github.com/kataras/iris/v12"
	jwtmiddleware "github.com/kataras/iris/v12/middleware"
)
func myHandler(ctx iris.Context) {
	//如果解密成功，将会进入这里，获取了解密了的token
	token := ctx.Values().Get("jwt").(*jwt.Token)
	//或者这样
	//userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	//userMsg["id"].(float64)==1
	//userMsg["nick_name"].(string) == iris

	ctx.Writef("This is an anthenticated request\n")
	ctx.Writef("Claim conent:\n")
	//token数据结构
	ctx.Writef("%s",token.Signature)

}

func main() {
	app := iris.New()
	//jwt中间件
	jwtHandler := jwtMiddleware.New(jwtmiddleware.Config{

	})
}