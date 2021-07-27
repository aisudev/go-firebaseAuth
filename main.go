package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi")
	})

	if err := e.StartAutoTLS(":8080"); err != nil {
		log.Fatalln(err.Error())
	}
	// if err := e.StartTLS(":9999", "server.cert", "server.key"); err != nil {
	// 	log.Fatalln(err.Error())
	// }
}
