package userroute

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/config"
	"my-gin-app/internal/cart/v1"
	"my-gin-app/internal/order/v1"
	review "my-gin-app/internal/product/reviews/v1"
	"my-gin-app/internal/product/v1"
	"my-gin-app/internal/user/v1/auth"
	"my-gin-app/internal/wishlist/v1"
	"my-gin-app/pkg/middleware"
)

func UserRouter(r *gin.Engine, config *config.Config) {

	v1 := r.Group("/v1")
	{
		userRepo := auth.NewUserRepoV1(config)
		userServices := auth.NewUserServiceV1(userRepo)
		userHandler := auth.NewUserHandlerV1(userServices)

		userV1 := v1.Group("/user")
		{

			userV1.POST("/register", userHandler.Register)
			userV1.POST("/login", userHandler.Logine)
			userV1.POST("/logout", userHandler.LogOut)

		}

		productRepo := product.NewProducctRepoV1(config)
		productServices := product.NewProducctServicesV1(productRepo)
		productHandler := product.NewProductHandlerV1(productServices)

		product := v1.Group("/product")
		{
			product.GET("/", productHandler.GetAllProduct)
			product.GET("/:id", productHandler.GetProductById)
			product.GET("/search", productHandler.SerchProduct)
			product.GET("/filter", productHandler.FilterProducts)
		}

		auth := v1.Group("/auth", middleware.AuthMiddleware())
		{
			cartrepo := cart.NewrepoCartV1(config)
			cartSrvices := cart.NewServiceCartV1(cartrepo)
			carthandler := cart.NewHandleCartV1(cartSrvices)

			cart := auth.Group("/cart")
			{
				cart.POST("/", carthandler.AddToCarthancler)
				cart.GET("/", carthandler.GetCartItemsHandler)
			}

			wishList := wishlist.NewWishListrepo(config)
			wislistSerives := wishlist.NewWishlistServises(wishList)
			wishlistHanlder := wishlist.NewWishListHanlder(wislistSerives)

			wishlist := auth.Group("/wishlist")
			{
				wishlist.POST("/", wishlistHanlder.WishListController)
				wishlist.GET("/", wishlistHanlder.GetAllwishlistItem)

			}

			orderRepo := order.NewOrderRepoV1(*config)
			orderServices := order.NewOrderServiceV1(orderRepo)
			orderHnalder := order.NewOrderHnalderV1(orderServices)

			order := auth.Group("/order")
			{
				order.POST("/", orderHnalder.OrderItemsChckOut)
				order.GET("/", orderHnalder.GetAllOrder)
				order.PUT("/", orderHnalder.CancellOrder)
				order.POST("/checkout", orderHnalder.OrderChckOut)
			}

			reviewRepo := review.NewRviewRepo(config)
			reviewServices := review.NewRviewServies(reviewRepo)
			reviewHnalder := review.NewRviewHandler(reviewServices)

			review := auth.Group("/review")
			{
				review.POST("/add", reviewHnalder.AddReview)
				review.DELETE("/delete", reviewHnalder.DeleteReviews)
				review.PUT("/update", reviewHnalder.Updatedreview)
				// review.PUT("/", reviewHnalder.CancellOrder)
			}
		}
	}
}
