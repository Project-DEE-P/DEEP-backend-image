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
		Image *multipart.FileHeader `form:"image" validate:"image,requried"`
	}

	// OutImageCreate 구조체는 /api/images/image를 호출의 응답의 사용되는 모델입니다.
	OutImageCreate struct {
		Ident string `json:"ident"`
	}

	InImageSelect struct {
		// path/value
		Ident string `validate:"requried"`
	}

	InImageUpdate struct {
		Ident string                `validate:"required"`
		Image *multipart.FileHeader `form:"image" validate:"image,required"`
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

func (i *InImageSelect) ParseX(c *fiber.Ctx) *InImageSelect {

	// parsing Image
	i.Ident = c.Params("ident")

	// validate
	validate(i)

	// retunn
	return i

}

func (i *InImageUpdate) ParseX(c *fiber.Ctx) *InImageUpdate {
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
