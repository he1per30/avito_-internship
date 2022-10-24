package user

type Repository interface {
	IncrementBalance(userId int, sum float64) error
	ReserveAmount(userId int, sum float64, serviceId int, orderId int) error
	GetBalance(userId int) (float64, error)
	RevenueRecognition(userId int, sum float64, serviceId int, orderId int) error
	GetReport(year, month string, serviceId int) error
}
