package middleware

import (
	"context"
	"net/http"
	"path/filepath"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func AuthMiddleware(auth *auth.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearer := c.Request().Header.Get("Authorization")
			authorizationToken := strings.Split(bearer, " ")

			if len(authorizationToken) != 2 {
				return c.String(http.StatusBadRequest, "invalid token")
			}

			token, err := auth.VerifyIDToken(context.Background(), authorizationToken[1])
			if err != nil {
				return c.String(http.StatusUnauthorized, "Unauthorization")
			}
			c.Set("uid", token.UID)

			return next(c)
		}
	}
}

func FirebaseAuthSetup() *auth.Client {
	serviceKey, err := filepath.Abs("serviceKey.json")
	if err != nil {
		panic("Can't load serviceKey.json")
	}

	options := option.WithCredentialsFile(serviceKey)

	app, err := firebase.NewApp(context.Background(), nil, options)
	if err != nil {
		panic(err.Error())
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(err.Error())
	}

	return auth
}
