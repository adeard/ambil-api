package domain

type MerchantGalleryRequest struct {
	MerchantId int    `json:"merchant_id" gorm:"column:merchant_id;"`
	Filepath   string `json:"filepath" gorm:"column:filepath;"`
	CreatedAt  string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  string `json:"updated_at" gorm:"column:updated_at;"`
}

type MerchantGalleryData struct {
	Id int `json:"id" gorm:"column:id;"`
	MerchantGalleryRequest
}

func (MerchantGalleryData) TableName() string {
	return "merchant_galleries"
}

type MerchantGalleryFilterRequest struct {
	MerchantGalleryRequest
	FilterRequest
}
