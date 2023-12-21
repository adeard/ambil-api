package domain

type UserRequest struct {
	Email     string `json:"email" gorm:"column:email;"`
	Password  string `json:"password" gorm:"column:password;"`
	IsActive  int    `json:"is_active" gorm:"column:is_active;"`
	CreatedAt string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at;"`
}

type UserData struct {
	Id int `json:"id" gorm:"column:id;"`
	UserRequest
}

func (UserData) TableName() string {
	return "users"
}
