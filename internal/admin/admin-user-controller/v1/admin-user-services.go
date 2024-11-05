package adminUser

import (
	"my-gin-app/internal/models"
)

type adminUserServicesImpl struct {
	repo AdminUserrepo
}

func NewAdminUserServices(repo AdminUserrepo) AdminUserServices {
	return &adminUserServicesImpl{repo: repo}
}

func (aus adminUserServicesImpl) GetAlluser() ([]models.User, error) {
	var users []models.User
	err := aus.repo.GetAlluserRepo(&users)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (aus adminUserServicesImpl) GetUserById(id string) (models.User, error) {
	var user models.User
	err := aus.repo.GetUserById(&user, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (aus adminUserServicesImpl) EditUser(updateUser map[string]interface{}, id string) (models.User, error) {
	var user models.User
	err := aus.repo.UpdateUser(updateUser, &user, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (aus adminUserServicesImpl) BlockuserServieces(UserStatus string, user_id string) (string, error) {
	var user models.User
	err := aus.repo.GetUserById(&user, user_id)
	if err != nil {
		return "user Not font", err
	}
	if UserStatus == "true" || UserStatus == "block" {
		user.Ban = true
		aus.repo.UserSave(&user)
		if err != nil {
			return "user Not font", err
		}
		return "user Bloked", nil
	} else {
		user.Ban = false
		aus.repo.UserSave(&user)
		if err != nil {
			return "user Not font", err
		}
		return "user Unbloked", nil
	}
}
