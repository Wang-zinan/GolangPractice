package main

import (
	"ginEssential/Controller"
	"ginEssential/middleware"
	"github.com/gin-gonic/gin"
)

func  CollectRoute(r *gin.Engine)  *gin.Engine{
	r.POST("/api/auth/register", Controller.Register)      //注册
	r.POST("/api/auth/login", Controller.Login)     //登录
	r.GET("api/auth/info",middleware.AuthMiddleware(),Controller.Info)     //获取用户信息
	r.POST("api/application/install",Controller.Install)    //安装应用
	r.POST("apo/application/download",Controller.Download)
	return r
}