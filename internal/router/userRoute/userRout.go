package userrouteHandler

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/config"
	authHandler "my-gin-app/internal/user/auth/handler"
	authrepo "my-gin-app/internal/user/auth/repo"
	authSevices "my-gin-app/internal/user/auth/services"
	cartHandler "my-gin-app/internal/user/cart/handler"
	cartRepo "my-gin-app/internal/user/cart/repo"
	cartServices "my-gin-app/internal/user/cart/services"
	orderHandler "my-gin-app/internal/user/order/handler"
	orderRepo "my-gin-app/internal/user/order/repo"
	orderServices "my-gin-app/internal/user/order/service"
	productHandler "my-gin-app/internal/user/product/handler"
	productRepo "my-gin-app/internal/user/product/repo"
	productServices "my-gin-app/internal/user/product/services"
	reviewHandler "my-gin-app/internal/user/reviews/handler"
	reviewrepo "my-gin-app/internal/user/reviews/repo"
	reviewServices "my-gin-app/internal/user/reviews/services"
	wishlistHandler "my-gin-app/internal/user/wishlist/handler"
	wishlistRepo "my-gin-app/internal/user/wishlist/repo"
	wishlistServices "my-gin-app/internal/user/wishlist/services"
	"my-gin-app/pkg/middleware"
)

func UserRouter(r *gin.Engine, config *config.Config) {

	v1 := r.Group("/v1")
	{
		userRepo := authrepo.NewUserRepoV1(config)
		userServices := authSevices.NewUserServiceV1(userRepo)
		userHandler := authHandler.NewUserHandlerV1(userServices)

		userV1 := v1.Group("/user")
		{
			userV1.POST("/register", userHandler.Register)
			userV1.POST("/login", userHandler.Logine)
			userV1.POST("/logout", userHandler.LogOut)
		}

		productRepo := productRepo.NewProducctRepoV1(config)
		productServices := productServices.NewProducctServicesV1(productRepo)
		productHandler := productHandler.NewProductHandlerV1(productServices)

		product := v1.Group("/product")
		{
			product.GET("/", productHandler.GetAllProduct)
			product.GET("/:id", productHandler.GetProductById)
			product.GET("/search", productHandler.SerchProduct)
			product.GET("/filter", productHandler.FilterProducts)
		}

		auth := v1.Group("/auth", middleware.AuthMiddleware())
		{
			cartrepo := cartRepo.NewrepoCartV1(config)
			cartSrvices := cartServices.NewServiceCartV1(cartrepo)
			carthandler := cartHandler.NewHandleCartV1(cartSrvices)

			cart := auth.Group("/cart")
			{
				cart.POST("/", carthandler.AddToCarthancler)
				cart.GET("/", carthandler.GetCartItemsHandler)
			}

			wishList := wishlistRepo.NewWishListrepo(config)
			wislistSerives := wishlistServices.NewWishlistServises(wishList)
			wishlistHanlder := wishlistHandler.NewWishListHanlder(wislistSerives)

			wishlist := auth.Group("/wishlist")
			{
				wishlist.POST("/", wishlistHanlder.WishListController)
				wishlist.GET("/", wishlistHanlder.GetAllwishlistItem)
			}

			orderRepo := orderRepo.NewOrderRepoV1(*config)
			orderServices := orderServices.NewOrderServiceV1(orderRepo)
			orderHnalder := orderHandler.NewOrderHnalderV1(orderServices)

			order := auth.Group("/order")
			{
				order.POST("/", orderHnalder.OrderItemsChckOut)
				order.GET("/", orderHnalder.GetAllOrder)
				order.PUT("/", orderHnalder.CancellOrder)
				order.POST("/checkout", orderHnalder.OrderChckOut)
			}

			reviewRepo := reviewrepo.NewRviewRepo(config)
			reviewServices := reviewServices.NewRviewServies(reviewRepo)
			reviewHnalder := reviewHandler.NewRviewHandler(reviewServices)

			review := auth.Group("/review")
			{
				review.POST("/add", reviewHnalder.AddReview)
				review.DELETE("/delete", reviewHnalder.DeleteReviews)
				review.PUT("/update", reviewHnalder.Updatedreview)

			}
		}
	}
}
