package router

import (
	"gin-study/controller"
	"gin-study/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	//写入gin日志
	//gin.DisableConsoleColor()
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//gin.DefaultErrorWriter = io.MultiWriter(f)
	router := gin.Default()
	router.Use(middlewares...)

	//demo
	//v1 := router.Group("/demo")
	//v1.Use(middleware.RecoveryMiddleware(), middleware.RequestLog(), middleware.IPAuthMiddleware())
	//{
	//	controller.CalcRegister(v1)
	//}

	//api
	store := sessions.NewCookieStore([]byte("secret"))
	authNormalGroup := router.Group("/auth")
	authController := &controller.Auth{}
	authNormalGroup.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog())
	authNormalGroup.POST("/login", authController.Login)
	authNormalGroup.GET("/loginout", authController.LoginOut)

	calcAuthGroup := router.Group("/calc")
	calcController := &controller.Calc{}
	calcAuthGroup.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware())
	calcAuthGroup.GET("/healthcalc", calcController.HealthCalc)

	return router
}
