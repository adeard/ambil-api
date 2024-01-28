package domain

type CategoryRequest struct {
	Name      string `json:"name" gorm:"column:name;"`
	IsActive  int    `json:"is_active" gorm:"column:is_active;"`
	Image     string `json:"image" gorm:"column:image;"`
	CreatedAt string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at;"`
}

type CategoryData struct {
	Id int `json:"id" gorm:"column:id;"`
	CategoryRequest
}

func (CategoryData) TableName() string {
	return "categories"
}

type CategoryFilterRequest struct {
	CategoryRequest
	FilterRequest
}
