package model

import "github.com/labstack/echo/v4"

type ConfigRouter struct {
	App           *echo.Echo
	RandomUserUrl string
}
