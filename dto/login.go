package dto

import (
	"errors"
	"gin-study/public"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type LoginInput struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

func (o *LoginInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

type ListPageInput struct {
	Page string `form:"page" json:"page" validate:"required"`
	Name string `form:"name" json:"name" validate:""`
}

func (o *ListPageInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}
