package reviewHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
	review "my-gin-app/internal/user/reviews/interface"
)

type reviewHandlerImpl struct {
	services review.ReviewrServices
}

func NewRviewHandler(services review.ReviewrServices) review.ReviewHandler {
	return &reviewHandlerImpl{services: services}
}

func (rh reviewHandlerImpl) DeleteReviews(ctx *gin.Context) {
	pid := ctx.Query("review_id")
	id, _ := ctx.Get("user_Id")
	userIDStr, _ := id.(string)
	err := rh.services.DeleteReview(pid, userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "you Deleted Review",
	})
}

func (rh reviewHandlerImpl) Updatedreview(ctx *gin.Context) {
	var review models.Review
	pid := ctx.Query("review_id")
	id, _ := ctx.Get("user_Id")
	userIDStr, _ := id.(string)
	err := ctx.BindJSON(&review)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = rh.services.UpdateReview(pid, review, userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "you Updated Review",
	})
}
func (rh reviewHandlerImpl) AddReview(ctx *gin.Context) {
	var review models.Review
	pid := ctx.Query("productId")
	err := ctx.BindJSON(&review)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, _ := ctx.Get("user_Id")
	userIDStr, _ := id.(string)
	err = rh.services.AddreviewByUser(userIDStr, review, pid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "tnx for Review",
	})
}
