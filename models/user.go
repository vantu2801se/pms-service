package models

type User struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FullName  string `json:"full_name"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeleteFlg bool   `json:"delete_flg"`
}

func (u *User) TableName() string {
	return "pms_users"
}
