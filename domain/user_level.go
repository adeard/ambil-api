package domain

type UserLevelRequest struct {
	UserId    int    `json:"user_id" gorm:"column:user_id;"`
	LevelName string `json:"level_name" gorm:"column:level_name;"`
	IsActive  int    `json:"is_active" gorm:"column:is_active;"`
	CreatedAt string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at;"`
}

type UserLevelData struct {
	Id int `json:"id" gorm:"column:id;"`
	UserLevelRequest
}

func (UserLevelData) TableName() string {
	return "user_levels"
}
