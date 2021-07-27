package main

import (
	"firebaseAuth/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	auth := middleware.FirebaseAuthSetup()

	e.Use(middleware.AuthMiddleware(auth))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": "Authorization",
		})
	})

	e.Start(":8080")
}
