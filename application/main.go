package main

import (
	"app/adapters/firestore"
	"app/controller"
	"app/domain/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	//rep := inmemory.NewBancotemp()
	rep := firestore.NewFireStoreConnection()

	createRole := usecase.NewCreateRoleUseCase(rep)
	getRoleById := usecase.NewGetRoleUseCase(rep)
	getAllRoles := usecase.NewGetAllRolesUseCase(rep)
	deleteRole := usecase.NewDeleteRoleUseCase(rep)
	updateRole := usecase.NewUpdateRoleUseCase(rep)

	createRoleController := controller.CreateRoleController(createRole)
	createGetRoleController := controller.GetRoleController(getRoleById)
	createGetAllRolesController := controller.GetAllRolesController(getAllRoles)
	createDeleteRoleController := controller.DeleteRoleController(deleteRole)
	createUpdateRoleController := controller.UpdateRoleUseCaseController(updateRole)

	r := gin.Default()

	r.POST("/role", createRoleController)
	r.GET("/role/:id", createGetRoleController)
	r.GET("/roles", createGetAllRolesController)
	r.DELETE("/role/:id", createDeleteRoleController)
	r.PUT("/role/:id", createUpdateRoleController)

	r.Run(":8070")
}
