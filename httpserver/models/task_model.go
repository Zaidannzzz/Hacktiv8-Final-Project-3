package models

type TaskModel struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	UserID      uint   `json:"user_id"`
	Category_id string `json:"category_id" gorm:"foreignKey:UserID;references:ID"`
}

func (TaskModel) TableName() string {
	return "public.Task"
}
