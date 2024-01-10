package domain

type BannerTypeRequest struct {
	Name      string `json:"name" gorm:"column:name;"`
	IsActive  int    `json:"is_active" gorm:"column:is_active;"`
	CreatedAt string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at;"`
}

type BannerTypeData struct {
	Id int `json:"id" gorm:"column:id;"`
	BannerTypeRequest
}

func (BannerTypeData) TableName() string {
	return "categories"
}

type BannerTypeFilterRequest struct {
	BannerTypeRequest
	FilterRequest
}
