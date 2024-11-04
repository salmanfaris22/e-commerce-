package order

type orderSerivesimpl struct {
	repo OrderRepo
}

func NewOrderSerivesV1(repo OrderRepo) OrderService {
	return &orderSerivesimpl{repo: repo}
}
