package models

import "time"

type Users struct {
	ID        uint      `gorm:"column:id_users;primary_key" json:"id"`
	Name      string    `gorm:"column:users_name" json:"name"`
	Email     string    `gorm:"column:users_email" json:"email"`
	Password  string    `gorm:"column:users_pass" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	IDRoles   uint      `gorm:"column:id_roles" json:"id_roles"` // Foreign key field
	Roles     Roles     `gorm:"foreignKey:IDRoles;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"roles"` // Relasi ke Roles
}

func (Users) TableName() string {
	return "users"
}
