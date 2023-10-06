package validator

import (
	"errors"
	"log"
	"reflect"
	"strings"
	"unicode"

	"github.com/3lur/go-mall/internal/common/reason"
	"github.com/3lur/go-mall/pkg/console"
	"github.com/3lur/go-mall/pkg/e"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Validate *MyValidator

type MyValidator struct {
	Validate *validator.Validate
	Tran     ut.Translator
}

// FormErrorField indicates the current form error content. which field is error and error message.
type FormErrorField struct {
	ErrorField string `json:"error_field"`
	ErrorMsg   string `json:"error_msg"`
}

func init() {
	// 国际化
	// tran := getTran()
	validate := validator.New()
	Validate = &MyValidator{
		Validate: validate,
		// Tran:     tran,
	}
}

func (v *MyValidator) Check(obj any) (errFields []*FormErrorField, err error) {
	defer func() {
		if len(errFields) == 0 {
			return
		}
		for _, field := range errFields {
			if len(field.ErrorField) == 0 {
				continue
			}
			firstRune := []rune(field.ErrorMsg)[0]
			if !unicode.IsLetter(firstRune) || !unicode.Is(unicode.Latin, firstRune) {
				continue
			}
			upperFirstRune := unicode.ToUpper(firstRune)
			field.ErrorMsg = string(upperFirstRune) + field.ErrorMsg[1:]
			if !strings.HasSuffix(field.ErrorMsg, ".") {
				field.ErrorMsg += "."
			}
		}
	}()
	err = v.Validate.Struct(obj)
	if err != nil {
		var validErrs validator.ValidationErrors
		if !errors.As(err, &validErrs) {
			console.ErrorIf(err)
			return nil, errors.New("validate check exception")
		}

		for _, fieldErr := range validErrs {
			errField := &FormErrorField{
				ErrorField: fieldErr.Field(),
				ErrorMsg:   fieldErr.Translate(v.Tran),
			}

			sns := fieldErr.StructNamespace()

			_, fieldName, found := strings.Cut(sns, ".")

			if found {
				originalTag := getObjectTagByFieldName(obj, fieldName)
				if len(originalTag) > 0 {
					errField.ErrorField = originalTag
				}
			}
			errFields = append(errFields, errField)
		}
		if len(errFields) > 0 {
			errMsg := ""
			if len(errFields) == 1 {
				errMsg = errFields[0].ErrorMsg
			}
			return errFields, e.BadRequest(reason.RequestBodyError).WithMsg(errMsg)
		}
	}

	return nil, nil
}

func getObjectTagByFieldName(obj any, fieldName string) (tag string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	objT := reflect.TypeOf(obj)
	objT = objT.Elem()

	structField, exists := objT.FieldByName(fieldName)
	if !exists {
		return ""
	}
	tag = structField.Tag.Get("json")
	if len(tag) == 0 {
		return structField.Tag.Get("form")
	}
	return tag
}

func getTran() ut.Translator {
	tran, ok := ut.New(zh.New()).GetTranslator("zh-CN")
	if !ok {
		console.Exit("failed to get translator")
	}
	return tran
}
