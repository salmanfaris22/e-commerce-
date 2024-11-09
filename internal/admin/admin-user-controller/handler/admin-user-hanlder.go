package adminUserHandler

import (
	"github.com/gin-gonic/gin"

	adminUserInterface "my-gin-app/internal/admin/admin-user-controller/interface"
)

type adminUserHandlerImpl struct {
	service adminUserInterface.AdminUserServices
}

func NewAdminUserHandler(service adminUserInterface.AdminUserServices) adminUserInterface.AdminUserHandler {
	return &adminUserHandlerImpl{service: service}
}

func (auh adminUserHandlerImpl) GetUSer(ctx *gin.Context) {
	id := ctx.Query("user_id")
	if id == "" {
		ctx.JSON(400, gin.H{
			"error": "Query isshu",
		})
		return
	}
	user, err := auh.service.GetUserById(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "cant get user",
			"err":     err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": user,
	})
}

func (auh adminUserHandlerImpl) GetAllUSer(ctx *gin.Context) {
	user, err := auh.service.GetAlluser()
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "cant get user",
			"err":     err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": user,
	})
}

func (auh adminUserHandlerImpl) EditUser(ctx *gin.Context) {
	id := ctx.Query("user_id")
	if id == "" {
		ctx.JSON(400, gin.H{
			"error": "Query isshu",
		})
		return
	}
	user := make(map[string]interface{})
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "validate json",
		})
		return
	}
	newUser, err := auh.service.EditUser(user, id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": newUser,
	})
}

func (auh adminUserHandlerImpl) BlockUser(ctx *gin.Context) {
	id := ctx.Query("user_id")
	status := ctx.Query("status")
	if id == "" || status == "" {
		ctx.JSON(400, gin.H{
			"error": "Query isshu",
		})
		return
	}
	str, err := auh.service.BlockuserServieces(status, id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error":   err.Error(),
			"message": str,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": str,
	})
}
