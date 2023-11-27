package main

import (
	"DEEP-backend-image/cerrors"
	"DEEP-backend-image/controller"
	"DEEP-backend-image/middleware"
	"flag"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var port string

func init() {
	flogPort := flag.Int("p", 8080, "Enter the port")
	flag.Parse()
	port = ":" + strconv.Itoa(*flogPort)
}

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "sexy image server for deep",
		Prefork:      true,
		ErrorHandler: cerrors.NewErrorHandler(),
	})

	app.Use(recover.New())
	app.Use(middleware.Authenticate)
	controller.Route(app)

	// 데이터베이스 테이블 ㅅ생성 확인
	log.Fatal(app.Listen(port))
}
