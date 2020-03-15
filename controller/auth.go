package controller

import (
	"errors"
	"gin-study/dto"
	"gin-study/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Auth struct {
}

func (auth *Auth) Login(c *gin.Context) {
	api := &dto.LoginInput{}
	if err := api.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	//todo 这块应该从db验证
	if api.Username == "admin" && api.Password == "123456" {
		session := sessions.Default(c)
		session.Set("user", api.Username)
		session.Save()
		middleware.ResponseSuccess(c, api)
	} else {
		middleware.ResponseError(c, 2002, errors.New("账号或密码错误"))
	}
	return
}

func (demo *Auth) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
	return
}
