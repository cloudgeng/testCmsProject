package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

/**
 * 管理员控制器
 */
 type AdminController struct {
 	//iris框架自动为每个请求都绑定上下文对象
 	ctx iris.Context

 	//admin功能实体
 	Service service.AdminService

 	//session对象
 	Session *sessions.Session
 }

const (
	ADMINTABLENAME = "admin"
	ADMIN          = "admin"
)

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

/**
 *	管理员退出功能
 *	请求类型： Get
 *	请求url:   admin/singout
 */

