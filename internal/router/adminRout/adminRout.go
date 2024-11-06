package adminrout

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/config"
	adminOrder "my-gin-app/internal/admin/admin-order-controller/v1"
	adminproduct "my-gin-app/internal/admin/admin-prodcut-controller/v1"
	adminUser "my-gin-app/internal/admin/admin-user-controller/v1"
	adminauth "my-gin-app/internal/admin/auth/v1"
	Admindashboard "my-gin-app/internal/admin/dashboard/v1"
	"my-gin-app/pkg/middleware"
)

func AdminRoutes(r *gin.Engine, config *config.Config) {
	v1 := r.Group("/v1")
	{

		userRepo := adminauth.NewAdminUserRepoV1(config)
		userServices := adminauth.NewAdminUserServiceV1(userRepo)
		userHandler := adminauth.NewAdminUserHandlerV1(userServices)
		auth := v1.Group("/admin")
		{
			auth.POST("/login", userHandler.Logine)
			auth.POST("/logout", userHandler.LogOut)
		}

		admin := v1.Group("/auth", middleware.AdminMiddlware())
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
			userRepo := adminUser.NewAdminUserrepo(config)
			userServices := adminUser.NewAdminUserServices(userRepo)
			userHndler := adminUser.NewAdminUserHandler(userServices)
			user := admin.Group("/user")
			{
				user.GET("/", userHndler.GetUSer)
				user.GET("/all", userHndler.GetAllUSer)
				user.PUT("/update", userHndler.EditUser)
				user.PUT("/block", userHndler.BlockUser)

			}

			orderRepo := adminOrder.NewAdminOrderrepo(config)
			orderServices := adminOrder.NewAdminOrdeServices(orderRepo)
			orderrHndler := adminOrder.NewAdminOrdeHandler(orderServices)
			order := admin.Group("/order")
			{
				order.GET("/all", orderrHndler.GetAllOrderAdmin)
				order.PUT("/controll", orderrHndler.AdminOrderControll)
			}

			dashboardRepo := Admindashboard.NewAdminDhasBoardpo(config)
			dashboardServices := Admindashboard.NewAdminOrdeServices(dashboardRepo)
			dashboardHndler := Admindashboard.NewAdminOrdeHanlder(dashboardServices)
			dashboar := admin.Group("/dashboar")
			{
				dashboar.GET("/", dashboardHndler.AdmindashBoardGetAll)

			}

		}

	}
}
