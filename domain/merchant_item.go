package domain

type MerchantItemRequest struct {
	MerchantId int    `json:"merchant_id" gorm:"column:merchant_id;"`
	Picture    string `json:"picture" gorm:"column:picture;"`
	Price      int    `json:"price" gorm:"column:price;"`
	Name       string `json:"name" gorm:"column:name;"`
	TotalSold  int    `json:"total_sold" gorm:"column:total_sold;"`
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

type MerchantItemFilterRequest struct {
	MerchantItemRequest
	FilterRequest
}
