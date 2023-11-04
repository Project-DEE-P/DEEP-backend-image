package conteoller

import (
	"DEEP-backend-image/cerrors"
	"DEEP-backend-image/database"
	"DEEP-backend-image/database/ent"
	"DEEP-backend-image/middleware"
	"DEEP-backend-image/model"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Route(app *fiber.App) {
	app.Post("/api/images/image", middleware.Authenticate, CreateImage)
	app.Get("/api/images/:ident", SelectImgae)
	app.Put("/api/images/:ident", UpdateImage)
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
	cerrors.Sniff(err)

	// response 해주기
	return c.Status(fiber.StatusCreated).JSON(model.OutGeneral{
		Status:  fiber.StatusCreated,
		Message: "Success Created",
		Data: model.OutImageCreate{
			Ident: database.Get().CreateImageX(c.Context(), fileBytes).ID.String(),
		},
	})
}

func SelectImgae(c *fiber.Ctx) error {
	clientRequest := new(model.InImageSelect).ParseX(c)

	file := database.Get().SelectImageX(c.Context(), uuid.Must(uuid.Parse(clientRequest.Ident))).Instance
	return c.Status(fiber.StatusOK).Send(file)
}

func UpdateImage(c *fiber.Ctx) error {
	// parsing
	clientRequest := new(model.InImageUpdate).ParseX(c)

	// multipart/form-data로 넘어온 image를 bytes로 변환
	fileContent, err := clientRequest.Image.Open()
	cerrors.Sniff(err)

	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)
	cerrors.Sniff(err)

	// call udpate repository
	database.Get().UpdateImageX(c.Context(), &ent.Image{
		ID:       uuid.Must(uuid.Parse(clientRequest.Ident)),
		Instance: fileBytes,
	})

	// response
	return c.Status(fiber.StatusOK).JSON(model.OutGeneral{
		Status:  fiber.StatusOK,
		Message: "Success Updated",
		Data:    nil,
	})
}
