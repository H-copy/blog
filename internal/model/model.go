package model

type Model struct {
	ID uint32 `json:"id" gorm:"primary_key"`
	CreatedBy string `json:"created_by"`
	CreatedOn uint32 `json:"created_on"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn uint32 `json:"modified_on"`
	DeleteOn uint32 `json:"delete_on"`
	IsDel uint8 `json:"is_del"`
}