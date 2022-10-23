package user

type Repository interface {
	IncrementBalance(userId int, sum float64) error
	ReserveAmount(userId int, sum float64, serviceId int, orderId int) error
}
