package domain

type OrderHistoryRequest struct {
	OrderId       int    `json:"order_id" gorm:"column:order_id;"`
	OrderStatusId int    `json:"order_status_id" gorm:"column:order_status_id;"`
	Notes         string `json:"notes" gorm:"column:notes;"`
	CreatedAt     string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt     string `json:"updated_at" gorm:"column:updated_at;"`
}

type OrderHistoryData struct {
	Id int `json:"id" gorm:"column:id;"`
	OrderHistoryRequest
}

func (OrderHistoryData) TableName() string {
	return "order_histories"
}
