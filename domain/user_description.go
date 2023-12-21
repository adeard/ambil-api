package domain

type UserDescriptionRequest struct {
	UserId      int    `json:"user_id" gorm:"column:user_id;"`
	Fullname    string `json:"fullname" gorm:"column:fullname;"`
	Picture     string `json:"picture" gorm:"column:picture;"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;"`
	CreatedAt   string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt   string `json:"updated_at" gorm:"column:updated_at;"`
}

type UserDescriptionData struct {
	Id int `json:"id" gorm:"column:id;"`
	UserDescriptionRequest
}

func (UserDescriptionData) TableName() string {
	return "user_descriptions"
}
