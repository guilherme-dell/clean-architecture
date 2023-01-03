package controller

import (
	"app/domain/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteRoleController(u usecase.DeleteRoleUseCase) gin.HandlerFunc {
	return func(g *gin.Context) {
		id := g.Param("id")

		err := u.Execute(id)
		if err != nil {
			g.JSON(http.StatusBadRequest, err.Error())
			return
		}
		g.JSON(http.StatusOK, "Role apagado :0")

	}
}
