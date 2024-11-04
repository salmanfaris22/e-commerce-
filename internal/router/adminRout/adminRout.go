package adminrout

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/config"
	adminauth "my-gin-app/internal/admin/auth/v1"
)

func AdminRoutes(r *gin.Engine, config *config.Config) {
	v1 := r.Group("/v1")
	{

		userRepo := adminauth.NewAdminUserRepoV1(config)
		userServices := adminauth.NewAdminUserServiceV1(userRepo)
		userHandler := adminauth.NewAdminUserHandlerV1(userServices)
		auth := v1.Group("/admin")
		{
			auth.POST("/logine", userHandler.Logine)
		}
	}
}
