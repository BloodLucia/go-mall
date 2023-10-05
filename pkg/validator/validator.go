package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thedevsaddam/govalidator"
)

type Validator struct {
	Validate *validator.Validate
}

// FormErrorField indicates the current form error content. which field is error and error message.
type FormErrorField struct {
	ErrorField string `json:"error_field"`
	ErrorMsg   string `json:"error_msg"`
}

// ValidatorFunc
type ValidatorFunc func(any, *gin.Context) map[string][]string

// Validate
func BindAndValidate(ctx *gin.Context, obj any, fn ValidatorFunc) bool {
	// 1. 解析请求，支持JSON、Form、Query
	if err := ctx.ShouldBind(obj); err != nil {
		// TODO response.BadRequest()
		return false
	}

	// 2. 验证
	errs := fn(obj, ctx)

	// 3. 判断验证是否通过
	if len(errs) > 0 {
		// TODO response.BadRequest
		return false
	}

	return true
}

func Validate(data any, rules govalidator.MapData, msgs govalidator.MapData) map[string][]string {
	o := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      msgs,
	}

	return govalidator.New(o).ValidateStruct()
}
