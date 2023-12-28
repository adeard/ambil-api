package domain

type MerchantRatingRequest struct {
	MerchantId int     `json:"merchant_id" gorm:"column:merchant_id;"`
	OrderId    int     `json:"order_id" gorm:"column:order_id;"`
	UserId     int     `json:"user_id" gorm:"column:user_id;"`
	Notes      string  `json:"notes" gorm:"column:notes;"`
	Score      float32 `json:"score" gorm:"column:score;"`
	Filepath   string  `json:"filepath" gorm:"column:filepath;"`
	CreatedAt  string  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  string  `json:"updated_at" gorm:"column:updated_at;"`
}

type MerchantRatingData struct {
	Id int `json:"id" gorm:"column:id;"`
	MerchantRatingRequest
}

func (MerchantRatingData) TableName() string {
	return "merchant_ratings"
}

type MerchantRatingFilterRequest struct {
	MerchantRatingRequest
	FilterRequest
}
