package model

import (
	"DEEP-backend-image/cerrors"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

// In 접미사는 client에게 요청된 모델임을 뜻합니다.
// Out 접미사는 client에게 반환하는 모델임을 뜻합니다.
type (
	// InImageCreate 구조체는 /api/images/image를 호출했을 때 사용되는 모델입니다.
	InImageCreate struct {
		Image *multipart.FileHeader `form:"image" validate:"image"`
	}

	// OutImageCreate 구조체는 /api/images/image를 호출의 응답의 사용되는 모델입니다.
	OutImageCreate struct {
		Ident string `json:"ident"`
	}

	InImageSelect struct {
		// path/value
		Ident int `validate:"requried"`
	}
)

func (i *InImageCreate) ParseX(c *fiber.Ctx) *InImageCreate {
	var err error

	// parsing Image
	if i.Image, err = c.FormFile("image"); err != nil {
		cerrors.ParsingErr(err.Error())
	}

	// 필드 유효성 검사
	validate(i)

	// 반환
	return i
}

// func (i *InImageSelect) ParseX(c *fiber.Ctx) *InImageSelect {
// 	var err error

// 	// parsing Image
// 	if i.Ident, err = strconv.Atoi(c.Params("ident")); err != nil {
// 		cerrors.ParsingErr(err.Error())
// 	}

// 	// validate
// 	validate(i)

// 	// retunn
// 	return i

// }
