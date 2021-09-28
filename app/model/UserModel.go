package model

type UserModel struct {
	*BaseModel
	UserId uint32 `gorm:primary_key json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (u UserModel) TableName() string {
	return "users"
}