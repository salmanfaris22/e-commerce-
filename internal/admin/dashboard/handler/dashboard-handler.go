package AdmindashboardHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	AdmindashboardInterface "my-gin-app/internal/admin/dashboard/interface"
)

type admindashBoardHhanlderImpl struct {
	services AdmindashboardInterface.AdminDashBoardServies
}

func NewAdminOrdeHanlder(services AdmindashboardInterface.AdminDashBoardServies) AdmindashboardInterface.AdminDashBoardInterface {
	return &admindashBoardHhanlderImpl{services: services}
}

func (adh admindashBoardHhanlderImpl) AdmindashBoardGetAll(ctx *gin.Context) {
	dashboard, err := adh.services.AdminDashBoardServices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve product analysis", "erro": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": dashboard,
	})
}
