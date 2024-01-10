package domain

type UserMerchantFavouriteRequest struct {
	UserId     int    `json:"user_id" gorm:"column:user_id;"`
	MerchantId int    `json:"merchant_id" gorm:"column:merchant_id;"`
	CreatedAt  string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  string `json:"updated_at" gorm:"column:updated_at;"`
}

type UserMerchantFavouriteData struct {
	Id int `json:"id" gorm:"column:id;"`
	UserMerchantFavouriteRequest
}

func (UserMerchantFavouriteData) TableName() string {
	return "user_merchant_favourites"
}
