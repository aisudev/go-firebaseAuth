package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi")
	})

<<<<<<< HEAD
	e.Start(":8080")
	// if err := e.StartAutoTLS(":9999"); err != nil {
	// 	log.Fatalln(err.Error())
	// }
=======
	if err := e.StartAutoTLS(":8080"); err != nil {
		log.Fatalln(err.Error())
	}
>>>>>>> 31f98fd151de4c5ee91d31de001a5c5e5c09f999
	// if err := e.StartTLS(":9999", "server.cert", "server.key"); err != nil {
	// 	log.Fatalln(err.Error())
	// }
}
