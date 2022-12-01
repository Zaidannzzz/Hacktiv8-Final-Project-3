package models

type TaskModel struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	UserID      uint   `json:"user_id"`
	CategoryID  string `json:"category_id" gorm:"foreignKey:CategoryID;references:ID"`
}

func (TaskModel) TableName() string {
	return "public.Task"
}
