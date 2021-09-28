package model

type Tag struct {
	*BaseModel
	TagId uint32 `gorm:primary_key json:"tag_id"`
	TagName string `json:"tag_name"`
	State uint8 `json:"state"`
}

func (t Tag) TableName() string {
	return "tag"
}