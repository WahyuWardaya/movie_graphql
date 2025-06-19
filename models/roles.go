package models

type Roles struct {
	ID   uint   `gorm:"primaryKey;column:id_roles;autoIncrement" json:"id"`
	Name string `gorm:"column:roles_name" json:"name"`
}

func (Roles) TableName() string {
	return "roles"
}
