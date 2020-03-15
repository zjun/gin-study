package controller

import (
	"gin-study/middleware"
	"github.com/gin-gonic/gin"
)

type Calc struct{}

func CalcRegister(router *gin.RouterGroup) {
	calc := Calc{}
	router.GET("/healthcalc", calc.HealthCalc)
}

func (calc *Calc) HealthCalc(c *gin.Context) {
	middleware.ResponseSuccess(c, "")
	return
}
