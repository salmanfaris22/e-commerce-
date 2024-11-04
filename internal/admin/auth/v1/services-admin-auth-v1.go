package adminauth

import (
	"errors"

	"my-gin-app/internal/models"
	"my-gin-app/pkg/utils"
)

type UserServiceImpl struct {
	userRepo UserRepo
}

func NewAdminUserServiceV1(repo UserRepo) UserService {
	return &UserServiceImpl{userRepo: repo}
}

func (us UserServiceImpl) LogineUser(admin models.Admin) (string, error) {
	var newAdmin models.Admin
	err := us.userRepo.FindByEmail(admin.Email, &newAdmin)
	if err != nil {
		return "", err
	}
	if admin.Password == newAdmin.Password {
		refreshToken, err := utils.GenerateRefreshToken(admin.ID)
		if err != nil {
			return "", err
		}
		return refreshToken, nil
	}
	return "", errors.New("password dont match")
}
