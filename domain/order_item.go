package domain

type OrderItemRequest struct {
	OrderId        int    `json:"order_id" gorm:"column:order_id;"`
	MerchantItemId int    `json:"merchant_item_id" gorm:"column:merchant_item_id;"`
	Notes          string `json:"notes" gorm:"column:notes;"`
	TotalCount     int    `json:"total_count" gorm:"column:total_count;"`
	TotalPrice     int    `json:"total_price" gorm:"column:total_price;"`
	CreatedAt      string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt      string `json:"updated_at" gorm:"column:updated_at;"`
}

type OrderItemData struct {
	Id int `json:"id" gorm:"column:id;"`
	OrderItemRequest
}

func (OrderItemData) TableName() string {
	return "order_items"
}
