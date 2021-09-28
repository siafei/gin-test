package model


type BaseModel struct {
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	CreatedOn uint32 `json:"created_on"`
	UpdatedOn uint32 `json:"updated_on"`
	DeletedOn uint32  `json:"deleted_on"`
	IsDel     uint8   `json:"is_del"`
}
