package models

type Category struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeleteFlg   bool   `json:"delete_flg"`
}

func (c *Category) TableName() string {
	return "pms_categories"
}
