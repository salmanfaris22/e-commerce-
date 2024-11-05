package adminrout

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/config"
	adminproduct "my-gin-app/internal/admin/admin-prodcut-controller/v1"
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

		admin := v1.Group("/auth")
		{
			produtRepo := adminproduct.NewAdminProductReposetries(config)
			produtServices := adminproduct.NewAdminProductServeces(produtRepo)
			productHnalder := adminproduct.NewAdminProductHandler(produtServices)
			product := admin.Group("/product")
			{
				product.POST("/add", productHnalder.AddProduct)
				product.PUT("/update", productHnalder.EditProduct)
				product.DELETE("/delete", productHnalder.DeleteProduct)
			}
		}

	}
}
