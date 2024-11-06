package review

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type ReviewrServices interface {
	AddreviewByUser(uid string, review models.Review, pId string) error
	DeleteReview(pid string, uId string) error
	UpdateReview(pid string, review models.Review, uId string) error
}
type Reviewrepo interface {
	SaveCommand(review *models.Review) error
	GetUserInfo(uid string, tempUser *models.User) error
	DeleteReviewCommand(id string, userID string) error
	Findreview(id string, userID string, review *models.Review) error
	Updatereview(review *models.Review) error
}
type ReviewHandler interface {
	AddReview(ctx *gin.Context)
	DeleteReviews(ctx *gin.Context)
	Updatedreview(ctx *gin.Context)
}
