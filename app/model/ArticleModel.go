package model

type ArticleModel struct {
	*BaseModel
	ArticleId uint32 `gorm:"primary_key" json:"article_id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Content   string `json:"content"`
	CoverImg  string `json:"cover_img"`
}

func (a ArticleModel) TableName() string  {
	return "article"
}
