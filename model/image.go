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
	InCreateImage struct {
		Image *multipart.FileHeader `form:"image" validate:"image,required"`
	}

	// OutImageCreate 구조체는 /api/images/image를 호출의 응답의 사용되는 모델입니다.
	OutCreateImage struct {
		Ident string `json:"ident"`
	}

	InSelectImage struct {
		// path/value
		Ident string `validate:"required"`
	}

	InUpdateImage struct {
		Ident string                `validate:"required"`
		Image *multipart.FileHeader `form:"image" validate:"image,required"`
	}

	InDeleteImage struct {
		Ident string `validate:"required"`
	}
)

func (i *InCreateImage) ParseX(c *fiber.Ctx) *InCreateImage {
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

func (i *InSelectImage) ParseX(c *fiber.Ctx) *InSelectImage {

	// parsing Image
	i.Ident = c.Params("ident")

	// validate
	validate(i)

	// retunn
	return i

}

func (i *InUpdateImage) ParseX(c *fiber.Ctx) *InUpdateImage {
	var err error

	// parsing Image
	if i.Image, err = c.FormFile("image"); err != nil {
		cerrors.ParsingErr(err.Error())
	}

	i.Ident = c.Params("ident")

	// 필드 유효성 검사
	validate(i)

	// 반환
	return i
}

func (i *InDeleteImage) ParseX(c *fiber.Ctx) *InDeleteImage {

	// parsing
	i.Ident = c.Params("ident")

	// validate
	validate(i)

	return i
}
