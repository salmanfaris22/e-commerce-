package userInfo

type UserHandler interface {
}

type userHandlerImpl struct {
	service UserServices
}

func NewUserHanlder(service UserServices) UserHandler {
	return &userHandlerImpl{service: service}
}
