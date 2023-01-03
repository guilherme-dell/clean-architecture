package controller

import (
	"app/domain/usecase"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllRolesController(u usecase.GetAllRolesUseCase) gin.HandlerFunc {
	return func(g *gin.Context) {
		roles, err := u.Execute()
		if err != nil {
			g.JSON(http.StatusBadRequest, errors.New("Nenhum role Cadastrado"))
			return
		}
		if roles == nil {
			g.JSON(http.StatusNotFound, errors.New("Nenhum role Cadastrado"))
			return
		}
		g.JSON(http.StatusOK, roles)
	}
}
