package domain

type OrderStatusRequest struct {
	Name      string `json:"name" gorm:"column:name;"`
	CreatedAt string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at;"`
}

type OrderStatusData struct {
	Id int `json:"id" gorm:"column:id;"`
	OrderStatusRequest
}

func (OrderStatusData) TableName() string {
	return "order_statuses"
}
