package domain

type MerchantCategoryRequest struct {
	MerchantId int    `json:"merchant_id" gorm:"column:merchant_id;"`
	CategoryId int    `json:"category_id" gorm:"column:category_id;"`
	CreatedAt  string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  string `json:"updated_at" gorm:"column:updated_at;"`
}

type MerchantCategoryData struct {
	Id int `json:"id" gorm:"column:id;"`
	MerchantCategoryRequest
}

func (MerchantCategoryData) TableName() string {
	return "merchant_categories"
}

type MerchantCategoryFilterRequest struct {
	MerchantCategoryRequest
	FilterRequest
}
