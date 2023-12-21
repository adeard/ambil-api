package domain

type MerchantItemRequest struct {
	MerchantId string `json:"merchant_id" gorm:"column:merchant_id;"`
	Picture    int    `json:"picture" gorm:"column:picture;"`
	Price      int32  `json:"price" gorm:"column:price;"`
	Name       string `json:"name" gorm:"column:name;"`
	TotalSold  int32  `json:"total_sold" gorm:"column:total_sold;"`
	IsActive   int    `json:"is_active" gorm:"column:is_active;"`
	CreatedAt  string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  string `json:"updated_at" gorm:"column:updated_at;"`
}

type MerchantItemData struct {
	Id int `json:"id" gorm:"column:id;"`
	MerchantItemRequest
}

func (MerchantItemData) TableName() string {
	return "merchant_items"
}
