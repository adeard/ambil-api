package order

import (
	"ambil-api/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.OrderRequest) (domain.OrderData, error)
	Update(input domain.OrderData, updateData map[string]interface{}) error
	GetAll(input domain.OrderFilterRequest) ([]domain.OrderData, error)
	GetDetail(orderId string) (domain.OrderData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll(input domain.OrderFilterRequest) ([]domain.OrderData, error) {
	var drivers []domain.OrderData
	q := r.db.Debug().Table("orders")

	if input.CancelReason != "" {
		q = q.Where("cancel_reason = ?", input.CancelReason)
	}

	if input.Notes != "" {
		q = q.Where("notes = ?", input.Notes)
	}

	if input.OrderDate != "" {
		q = q.Where("order_date = ?", input.OrderDate)
	}

	if input.OrderNo != "" {
		q = q.Where("order_no = ?", input.OrderNo)
	}

	if input.PaymentDate != "" {
		q = q.Where("payment_date = ?", input.PaymentDate)
	}

	if input.MerchantId != 0 {
		q = q.Where("merchant_id = ?", input.MerchantId)
	}

	if input.OrderStatusId != 0 {
		q = q.Where("order_status_id = ?", input.OrderStatusId)
	}

	if input.UserId != 0 {
		q = q.Where("user_id = ?", input.UserId)
	}

	if input.TotalPrice != 0 {
		q = q.Where("total_price = ?", input.TotalPrice)
	}

	if input.OrderBy != "" {
		sort := "asc"
		order := input.OrderBy

		if input.SortBy != "" {
			sort = input.SortBy
		}

		q = q.Order(order + " " + sort)
	}

	offset := (input.Limit * (input.Page - 1))

	err := q.
		Limit(input.Limit).
		Offset(offset).
		Find(&drivers).
		Error

	return drivers, err
}

func (r *repository) Store(input domain.OrderRequest) (domain.OrderData, error) {

	order := domain.OrderData{OrderRequest: input}

	err := r.db.Create(&order).Error

	return order, err
}

func (r *repository) GetDetail(orderId string) (domain.OrderData, error) {
	var order domain.OrderData
	err := r.db.Where("id = ? ", orderId).First(&order).Error

	return order, err
}

func (r *repository) Update(order domain.OrderData, updateData map[string]interface{}) error {
	err := r.db.Debug().
		Model(&order).
		Where("id = ?", order.Id).
		Updates(updateData).Error

	return err
}
