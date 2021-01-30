package test

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/boltdb"
	"time"
)

func main() {
	//db,_ := boltdb.New("./sessions/sessions.db",0666,"users")
	db, _ := boltdb.New("./sessions/sessions.db", 0666, "users")
	//使用不同的go协程来同步数据库
	db.Async(true)
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	sess := sessions.New(sessions.Config{
		Cookie:"sessionscookieid",
		Expires: 45 * time.Minute,
	})

	sess.UseDatabase(db)

	app := iris.New()
	app.Get("/",func(ctx iris.Context) {
		ctx.Writef("You should navigate to the /set,/get,/delete,/clear,/destroy instead")
	})

	app.Get("/set",func(ctx iris.Context) {
		s := sess.Start(ctx)
		s.Set("name","iris")
		ctx.Writef("All ok session setted to: %s",s.GetString("name"))
	})

	app.Get("/set/{key}/{value}",func(ctx iris.Context) {
		key, value := ctx.Params().Get("key"),ctx.Params().Get("value")
		s := sess.Start(ctx)
		s.Set(key,value)
		ctx.Writef("all ok session setted to: %s",s.GetString(key))
	})

	app.Get("/get",func(ctx iris.Context) {
		name := sess.Start(ctx).GetString("name")
		ctx.Writef("the name on the /set was: %s",name)
	})

	app.Get("/get/{key}",func(ctx iris.Context) {
		name := sess.Start(ctx).GetString(ctx.Params().Get("key"))
		ctx.Writef("The name on the /set was: %s", name)

	})

	app.Get("/delete", func(ctx iris.Context) {
		sess.Start(ctx).Delete("name")

	})

	app.Get("/clear",func(ctx iris.Context) {
		sess.Start(ctx).Clear()
	})
	app.Get("/destroy",func(ctx iris.Context){
		//destroy,删除整个会话数据和cookie
		sess.Destroy(ctx)
	})

	app.Get("/update",func(ctx iris.Context) {
		//更新过期日期与新日期
		sess.ShiftExpiration(ctx)
	})

	app.Run(iris.Addr(":8182"))


}
