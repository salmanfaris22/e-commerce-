package reviewServices

import (
	"errors"
	"strconv"

	"my-gin-app/internal/models"
	review "my-gin-app/internal/user/reviews/interface"
)

type rviewServiesImpl struct {
	repo review.Reviewrepo
}

func NewRviewServies(repo review.Reviewrepo) review.ReviewrServices {
	return &rviewServiesImpl{repo: repo}
}

func (rs rviewServiesImpl) AddreviewByUser(uid string, review models.Review, pId string) error {
	if pId == "" {
		return errors.New("cant find Product Id")
	}
	var tempUser models.User
	err := rs.repo.GetUserInfo(uid, &tempUser)
	if err != nil {
		return err
	}
	if review.Rating > 5 || review.Rating < 0 {
		return errors.New("rating isshu")
	}
	newId, err := strconv.ParseUint(uid, 10, 64)
	if err != nil {
		return err
	}
	newProduct, err := strconv.ParseUint(pId, 10, 64)
	if err != nil {
		return err
	}
	review.UserName = tempUser.FirstName
	review.UserID = uint(newId)
	review.ProductID = uint(newProduct)
	err = rs.repo.SaveCommand(&review)
	if err != nil {
		return err
	}
	return nil
}
func (rs rviewServiesImpl) DeleteReview(pid string, uId string) error {
	if pid == "" || uId == "" {
		return errors.New("cant Quaryu")
	}
	var review models.Review
	err := rs.repo.Findreview(pid, uId, &review)
	if err != nil {
		return err
	}
	err = rs.repo.DeleteReviewCommand(pid, uId)
	if err != nil {
		return err
	}
	return nil
}
func (rs rviewServiesImpl) UpdateReview(pid string, review models.Review, uId string) error {
	if pid == "" || uId == "" {
		return errors.New("cant Quaryu")
	}
	if review.Rating > 5 || review.Rating < 0 {
		return errors.New("rating isshu")
	}
	var newReview models.Review
	err := rs.repo.Findreview(pid, uId, &newReview)
	if err != nil {
		return err
	}
	review.ID = newReview.ID

	err = rs.repo.Updatereview(&review)
	if err != nil {
		return err
	}
	return nil
}
