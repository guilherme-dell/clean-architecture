package controller

import (
	entities "app/domain/entities/role"
	"app/domain/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoleController(u usecase.GetRoleUseCase) gin.HandlerFunc {
	return func(g *gin.Context) {
		id := g.Param("id")

		role, err := u.Execute(id)
		if err != nil {
			g.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if role == nil {
			g.JSON(http.StatusNotFound, err.Error())
			return
		}

		roleFound := entities.RoleOut{
			ID:      id,
			Nome:    role.Nome,
			Local:   role.Local,
			Cidade:  role.Cidade,
			Horario: role.Horario,
			Dia:     role.Dia,
			Tipo:    role.Tipo,
		}

		g.JSON(http.StatusOK, roleFound)
	}
}
