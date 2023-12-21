package domain

type OrdersRequest struct {
	UserId        int    `json:"user_id" gorm:"column:user_id;"`
	MerchantId    int    `json:"merchant_id" gorm:"column:merchant_id;"`
	OrderStatusId int    `json:"order_status_id" gorm:"column:order_status_id;"`
	OrderNo       string `json:"order_no" gorm:"column:order_no;"`
	PaymentDate   string `json:"payment_date" gorm:"column:payment_date;"`
	OrderDate     string `json:"order_date" gorm:"column:order_date;"`
	TotalPrice    int32  `json:"total_price" gorm:"column:total_price;"`
	Notes         string `json:"notes" gorm:"column:notes;"`
	CancelReason  string `json:"cancel_reason" gorm:"column:cancel_reason;"`
	CreatedAt     string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt     string `json:"updated_at" gorm:"column:updated_at;"`
}

type OrdersData struct {
	Id int `json:"id" gorm:"column:id;"`
	OrdersRequest
}

func (OrdersData) TableName() string {
	return "orders"
}
