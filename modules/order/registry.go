package order

import "gorm.io/gorm"

func OrderRegistry(db *gorm.DB) Service {
	orderRepository := NewRepository(db)
	orderService := NewService(orderRepository)

	return orderService
}
