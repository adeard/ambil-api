package order

import (
	"ambil-api/domain"
	"ambil-api/utils"
)

type Service interface {
	GetAll(input domain.OrderFilterRequest) ([]domain.OrderData, error)
	GetDetail(merchantId string) (domain.OrderData, error)
	Store(input domain.OrderRequest) (domain.OrderData, error)
	Update(driverId string, updatedData domain.OrderRequest) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.OrderFilterRequest) ([]domain.OrderData, error) {

	if input.Limit == 0 {
		input.Limit = 20
	}

	if input.Page == 0 {
		input.Page = 1
	}

	merchants, err := s.repository.GetAll(input)
	if err != nil {
		return []domain.OrderData{}, err
	}

	return merchants, err
}

func (s *service) GetDetail(merchantId string) (domain.OrderData, error) {
	merchant, err := s.repository.GetDetail(merchantId)

	return merchant, err
}

func (s *service) Store(input domain.OrderRequest) (domain.OrderData, error) {

	input.CreatedAt = utils.GetCurrentDateTime()
	input.UpdatedAt = utils.GetCurrentDateTime()

	merchant, err := s.repository.Store(input)
	if err != nil {
		return domain.OrderData{}, err
	}

	return merchant, err
}

func (s *service) Update(orderId string, input domain.OrderRequest) error {

	order, err := s.repository.GetDetail(orderId)
	if err != nil {
		return err
	}

	updateData := map[string]interface{}{
		"updated_at": utils.GetCurrentDateTime(),
	}

	if input.UserId != 0 {
		updateData["user_id"] = input.UserId
	}

	if input.CancelReason != "" {
		updateData["cancel_reason"] = input.CancelReason
	}

	if input.Notes != "" {
		updateData["notes"] = input.Notes
	}

	if input.OrderDate != "" {
		updateData["order_date"] = input.OrderDate
	}

	if input.OrderNo != "" {
		updateData["order_no"] = input.OrderNo
	}

	if input.PaymentDate != "" {
		updateData["payment_date"] = input.PaymentDate
	}

	if input.MerchantId != 0 {
		updateData["merchant_id"] = input.MerchantId
	}

	if input.OrderStatusId != 0 {
		updateData["order_status_id"] = input.OrderStatusId
	}

	if input.TotalPrice != 0 {
		updateData["total_price"] = input.TotalPrice
	}

	err = s.repository.Update(order, updateData)

	return err
}
