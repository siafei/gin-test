package model

type ArticleTagModel struct {
	*BaseModel
	ArticleTagId uint32 `gorm:primary_key json:"article_tag_id"`
	ArticleId    uint32 `json:"article_id"`
	TagId        uint32 `json:"tag_id"`
}

func (a ArticleTagModel) TableName() string  {
	return "article_tag"
}