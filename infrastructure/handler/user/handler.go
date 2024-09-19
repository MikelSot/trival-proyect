package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/MikelSot/trival-proyect/domain/user"
)

type handler struct {
	useCase user.UseCase
}

func newHandler(useCase user.UseCase) handler {
	return handler{useCase}
}

func (h handler) GetUniqueUsers(ctx echo.Context) error {
	results, err := h.useCase.GetUniqueUsers()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println("len(results)")
	fmt.Println(len(results))

	return ctx.JSON(http.StatusOK, results)
}
