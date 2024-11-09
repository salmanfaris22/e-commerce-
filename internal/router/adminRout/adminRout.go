package adminrout

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/config"
	adminOrderHandler "my-gin-app/internal/admin/admin-order-controller/handler"
	adminOrderRepo "my-gin-app/internal/admin/admin-order-controller/repo"
	adminOrderServices "my-gin-app/internal/admin/admin-order-controller/services"
	adminproductHandler "my-gin-app/internal/admin/admin-prodcut-controller/handler"
	adminproductRepo "my-gin-app/internal/admin/admin-prodcut-controller/repo"
	adminproductServices "my-gin-app/internal/admin/admin-prodcut-controller/services"
	adminUserHandler "my-gin-app/internal/admin/admin-user-controller/handler"
	adminUserRepo "my-gin-app/internal/admin/admin-user-controller/repo"
	adminUserServices "my-gin-app/internal/admin/admin-user-controller/services"
	adminauthHandler "my-gin-app/internal/admin/auth/handler"
	adminauthRepo "my-gin-app/internal/admin/auth/repo"
	adminauthServices "my-gin-app/internal/admin/auth/services"
	AdmindashboardHandler "my-gin-app/internal/admin/dashboard/handler"
	AdmindashboardRepo "my-gin-app/internal/admin/dashboard/repo"
	AdmindashboardServices "my-gin-app/internal/admin/dashboard/services"
	"my-gin-app/pkg/middleware"
)

func AdminRoutes(r *gin.Engine, config *config.Config) {
	v1 := r.Group("/v1")
	{

		userRepo := adminauthRepo.NewAdminUserRepoV1(config)
		userServices := adminauthServices.NewAdminUserServiceV1(userRepo)
		userHandler := adminauthHandler.NewAdminUserHandlerV1(userServices)
		auth := v1.Group("/admin")
		{
			auth.POST("/login", userHandler.Logine)
			auth.POST("/logout", userHandler.LogOut)
		}

		admin := v1.Group("/auth", middleware.AdminMiddlware())
		{
			produtRepo := adminproductRepo.NewAdminProductReposetries(config)
			produtServices := adminproductServices.NewAdminProductServeces(produtRepo)
			productHnalder := adminproductHandler.NewAdminProductHandler(produtServices)
			product := admin.Group("/product")
			{
				product.POST("/add", productHnalder.AddProduct)
				product.PUT("/update", productHnalder.EditProduct)
				product.DELETE("/delete", productHnalder.DeleteProduct)
			}
			userRepo := adminUserRepo.NewAdminUserrepo(config)
			userServices := adminUserServices.NewAdminUserServices(userRepo)
			userHndler := adminUserHandler.NewAdminUserHandler(userServices)
			user := admin.Group("/user")
			{
				user.GET("/", userHndler.GetUSer)
				user.GET("/all", userHndler.GetAllUSer)
				user.PUT("/update", userHndler.EditUser)
				user.PUT("/block", userHndler.BlockUser)

			}

			orderRepo := adminOrderRepo.NewAdminOrderrepo(config)
			orderServices := adminOrderServices.NewAdminOrdeServices(orderRepo)
			orderrHndler := adminOrderHandler.NewAdminOrdeHandler(orderServices)
			order := admin.Group("/order")
			{
				order.GET("/all", orderrHndler.GetAllOrderAdmin)
				order.GET("/byid", orderrHndler.AdminOrderByID)
				order.PUT("/controll", orderrHndler.AdminOrderControll)
			}

			dashboardRepo := AdmindashboardRepo.NewAdminDhasBoardpo(config)
			dashboardServices := AdmindashboardServices.NewAdminOrdeServices(dashboardRepo)
			dashboardHndler := AdmindashboardHandler.NewAdminOrdeHanlder(dashboardServices)
			dashboar := admin.Group("/dashboar")
			{
				dashboar.GET("/", dashboardHndler.AdmindashBoardGetAll)

			}

		}

	}
}
