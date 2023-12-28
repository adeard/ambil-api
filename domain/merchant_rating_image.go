package domain

type MerchantRatingImageRequest struct {
	MerchantRatingId int    `json:"merchant_rating_id" gorm:"column:merchant_rating_id;"`
	Filepath         string `json:"filepath" gorm:"column:filepath;"`
	CreatedAt        string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt        string `json:"updated_at" gorm:"column:updated_at;"`
}

type MerchantRatingImageData struct {
	Id int `json:"id" gorm:"column:id;"`
	MerchantRatingImageRequest
}

func (MerchantRatingImageData) TableName() string {
	return "merchant_rating_images"
}

type MerchantRatingImageFilterRequest struct {
	MerchantRatingImageRequest
	FilterRequest
}
