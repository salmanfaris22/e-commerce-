package Admindashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type admindashBoardHhanlderImpl struct {
	services AdminDashBoardServies
}

func NewAdminOrdeHanlder(services AdminDashBoardServies) AdminDashBoardInterface {
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
