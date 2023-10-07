package controller

import (
	"github.com/3lur/go-mall/internal/service"
	"github.com/3lur/go-mall/pkg/response"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"

	en "github.com/go-playground/locales/en"
	zh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type UserReq struct {
	Username string `validate:"required"`
}

type userController struct {
	userSrv service.UserService
}

// Register implements UserController.
func (uc *userController) Register(ctx *gin.Context) {

	var req UserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(422, err.Error())
		return
	}
	var validate *validator.Validate
	var uni *ut.UniversalTranslator

	zh := zh.New()
	en := en.New()

	uni = ut.New(en, en, zh)

	trans, _ := uni.GetTranslator("zh-CN")

	validate = validator.New()

	zh_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(req)
	if err != nil {
		// errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(422, err)
		return
	}
	response.Build(ctx, nil, "ok!")
}

type UserController interface {
	Register(ctx *gin.Context)
}

func NewUserController(userSrv service.UserService) UserController {
	return &userController{userSrv}
}
