package authSevices

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"my-gin-app/internal/models"
	authInterface "my-gin-app/internal/user/auth/interface"
	"my-gin-app/pkg/utils"
)

type UserServiceImpl struct {
	userRepo authInterface.UserRepo
}

func NewUserServiceV1(repo authInterface.UserRepo) authInterface.UserService {
	return &UserServiceImpl{userRepo: repo}
}

func (us UserServiceImpl) CreateUser(user models.User) (models.User, string, string, error) {
	existingUser := models.User{}
	if err := us.userRepo.FindByEmail(user.Email, &existingUser); err == nil {
		return user, "", "", errors.New("email already exists")
	}
	if err := us.userRepo.Save(&user); err != nil {
		return user, "", "", err
	}
	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return user, "", "", err
	}
	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return user, "", "", err
	}
	refreshTokenSet := models.UserToken{
		Token:     refreshToken,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	}
	if err := us.userRepo.TokenSave(&refreshTokenSet); err != nil {
		return user, "", "", err
	}
	return user, accessToken, refreshTokenSet.Token, nil
}

func (us UserServiceImpl) LogineUser(user models.User) (models.User, string, string, error) {
	var existingUser models.User
	if err := us.userRepo.FindByEmail(user.Email, &existingUser); err != nil {
		return user, "", "", errors.New("email can't find pleas register")
	}
	if !utils.CheckPasswordHash(user.Password, existingUser.Password) {
		return user, "", "", errors.New("password can't match")
	}
	if existingUser.Ban == true {
		return user, "", "", errors.New("your bloked user")
	}
	accessToken, err := utils.GenerateAccessToken(existingUser.ID)
	if err != nil {
		return user, "", "", err
	}
	var refreshToken models.UserToken
	err = us.userRepo.FindRefreshToken(&refreshToken, existingUser.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			refreshTokenSet, err := utils.GenerateRefreshToken(existingUser.ID)
			if err != nil {
				return user, "", "", err
			}
			refreshTokenCreat := models.UserToken{
				Token:     refreshTokenSet,
				UserID:    existingUser.ID,
				CreatedAt: time.Now(),
				ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
			}
			if err := us.userRepo.TokenSave(&refreshTokenCreat); err != nil {
				return user, "", "", err
			}
			return existingUser, accessToken, refreshTokenCreat.Token, nil
		} else {
			return user, "", "", err
		}
	}
	return existingUser, accessToken, refreshToken.Token, nil
}
