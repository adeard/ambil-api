package domain

type BannerRequest struct {
	UserId       int    `json:"user_id" gorm:"column:user_id;"`
	BannerTypeId int    `json:"banner_type_id" gorm:"column:banner_type_id;"`
	Filepath     string `json:"filepath" gorm:"column:filepath;"`
	IndexNo      int    `json:"index_no" gorm:"column:index_no;"`
	StartDate    string `json:"start_date" gorm:"column:start_date;"`
	EndDate      string `json:"end_date" gorm:"column:end_date;"`
	Link         string `json:"link" gorm:"column:link;"`
	IsActive     int    `json:"is_active" gorm:"column:is_active;"`
	CreatedAt    string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt    string `json:"updated_at" gorm:"column:updated_at;"`
}

type BannerData struct {
	Id int `json:"id" gorm:"column:id;"`
	BannerRequest
}

func (BannerData) TableName() string {
	return "categories"
}

type BannerFilterRequest struct {
	BannerRequest
	FilterRequest
}
