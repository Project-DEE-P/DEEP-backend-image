package conteoller

import (
	"DEEP-backend-image/cerrors"
	"DEEP-backend-image/database"
	"DEEP-backend-image/middleware"
	"DEEP-backend-image/model"
	"io"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Post("/api/images/image", middleware.Authenticate, CreateImage)
	app.Get("/api/images/:ident")
	app.Put("/api/images/:ident")
	app.Delete("/api/images/:ident")
}

func CreateImage(c *fiber.Ctx) error {
	// client reqeust parsing
	clientRequest := new(model.InImageCreate).ParseX(c)

	// multipart/form-data로 넘어온 image를 bytes로 변환
	fileContent, err := clientRequest.Image.Open()
	cerrors.Sniff(err)
	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)

	// response 해주기
	return c.Status(fiber.StatusCreated).JSON(model.OutGeneral{
		Status:  fiber.StatusCreated,
		Message: "Success Created",
		Data: model.OutImageCreate{
			Ident: database.Get().CreateImageX(c.Context(), fileBytes).ID.String(),
		},
	})
}

// func SelectImgae(c *fiber.Ctx) error {
// 	clientRequest := new(model.InImageSelect).ParseX(c)

// 	return c.Status(fiber.StatusOK).Send(database.Get().)
// }
