package model

import (
	"DEEP-backend-image/cerrors"
	"encoding/json"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
)

// v(validate)는 이미지 검증을 당담하는 변수 입니다.
var v *validator.Validate

var allowExtentions = []string{".jpeg", ".jpg", ".png", ".mov", ".svg"}

func init() {
	v = validator.New()

	v.RegisterValidation("image", func(fl validator.FieldLevel) bool {
		field := fl.Field()

		if fileHeader, ok := field.Interface().(multipart.FileHeader); ok {
			ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
			for _, allowedExt := range allowExtentions {
				if ext == allowedExt {
					return true
				}
			}
		}
		return false
	})
}

func validate(model any) {
	err := v.Struct(model)

	if err != nil {
		var messages []map[string]any
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]any{
				"filed":   err.Field(),
				"message": "this filed is " + err.Tag(),
			})
		}
		marshalMessages, err := json.Marshal(messages)
		cerrors.Sniff(err)

		cerrors.ValidationErr(string(marshalMessages))
	}
}
