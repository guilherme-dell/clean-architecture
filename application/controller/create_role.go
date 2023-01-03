package controller

import (
	"app/domain/entities/role"
	"app/domain/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRoleController(u usecase.CreateRoleUseCase) gin.HandlerFunc {
	return func(g *gin.Context) {

		var roleInput entities.RoleIn
		if err := g.ShouldBindJSON(&roleInput); err != nil {
			g.String(http.StatusBadRequest, err.Error())
			return
		}

		UseCaseOut, err := u.Execute(&roleInput)
		if err != nil {
			g.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		g.JSON(http.StatusCreated, UseCaseOut)
	}
}
