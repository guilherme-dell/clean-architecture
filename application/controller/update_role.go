package controller

import (
	entities "app/domain/entities/role"
	"app/domain/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateRoleUseCaseController(u usecase.UpdateRoleUseCase) gin.HandlerFunc {
	return func(g *gin.Context) {
		id := g.Param("id")

		var roleInput entities.RoleIn

		if err := g.ShouldBindJSON(&roleInput); err != nil {
			g.String(http.StatusBadRequest, err.Error())
			return
		}
		role := &entities.Role{
			ID:      id,
			Nome:    roleInput.Nome,
			Local:   roleInput.Local,
			Cidade:  roleInput.Cidade,
			Horario: roleInput.Horario,
			Dia:     roleInput.Dia,
			Tipo:    roleInput.Tipo,
		}
		roleModif, err := u.Execute(role)
		if err != nil {
			if err.Error() == "Role não encontrado" {
				g.JSON(http.StatusNotFound, err.Error())
				return
			}
			g.JSON(http.StatusBadRequest, err.Error())
			return
		}

		g.JSON(http.StatusOK, roleModif)
	}
}
