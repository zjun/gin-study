package main

import (
	"gin-study/common/lib"
	"gin-study/public"
	"gin-study/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//lib.InitModule("./conf/dev/",[]string{"base","mysql","redis",})
	lib.InitModule("./conf/dev/", []string{"base"})
	defer lib.Destroy()
	//public.InitMysql()
	public.InitValidate()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
