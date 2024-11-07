package userInfo

type UserServices interface {
}

type userServiesImpl struct {
	repo UserRepo
}

func NewUserServices(repo UserRepo) UserServices {
	return &userServiesImpl{repo: repo}
}
