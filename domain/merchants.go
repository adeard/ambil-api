package domain

type MerchantsRequest struct {
	UserId      int     `json:"user_id" gorm:"column:user_id;"`
	Picture     string  `json:"picture" gorm:"column:picture;"`
	Name        string  `json:"name" gorm:"column:name;"`
	Address     string  `json:"address" gorm:"column:address;"`
	Coordinate  string  `json:"coordinate" gorm:"column:coordinate;"`
	PhoneNumber string  `json:"phone_number" gorm:"column:phone_number;"`
	IsActive    int     `json:"is_active" gorm:"column:is_active;"`
	Score       float32 `json:"score" gorm:"column:score;"`
	CreatedAt   string  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt   string  `json:"updated_at" gorm:"column:updated_at;"`
}

type MerchantsData struct {
	Id int `json:"id" gorm:"column:id;"`
	MerchantsRequest
}

func (MerchantsData) TableName() string {
	return "merchants"
}
