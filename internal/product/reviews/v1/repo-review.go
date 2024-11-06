package review

import (
	"my-gin-app/config"
	"my-gin-app/internal/models"
)

type rviewreporImpl struct {
	config *config.Config
}

func NewRviewRepo(config *config.Config) Reviewrepo {
	return &rviewreporImpl{config: config}
}

func (rr rviewreporImpl) GetUserInfo(uid string, tempUser *models.User) error {
	return rr.config.DB.Where("id=?", uid).First(&tempUser).Error
}
func (rr rviewreporImpl) SaveCommand(review *models.Review) error {
	return rr.config.DB.Save(&review).Error
}
func (rr rviewreporImpl) DeleteReviewCommand(id string, userID string) error {
	return rr.config.DB.Where("id=? AND user_id=?", id, userID).Delete(&models.Review{}).Error
}
func (rr rviewreporImpl) Findreview(id string, userID string, review *models.Review) error {
	return rr.config.DB.Where("id=? AND user_id=?", id, userID).First(&review).Error
}

func (rr rviewreporImpl) Updatereview(review *models.Review) error {
	return rr.config.DB.Model(&models.Review{}).Where("id = ?", review.ID).Updates(map[string]interface{}{
		"comment": review.Comment,
		"rating":  review.Rating,
	}).Error
}
