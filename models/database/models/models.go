package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name"`
	Password  string         `json:"password"`
	Id_role   uint           `json:"id_role"`            // reference to Role.ID (belongs to)
	Role      Role           `gorm:"foreignkey:Id_role"` // Add a foreign key to the User model
	Email     string         `gorm:"email;unique"`
	IsActive  bool           `gorm:"default:true"`
}

type Role struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"name unique"`
}

type Permission struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name unique"`
}

type RolePermission struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Id_role       uint           `json:"id_role"`
	Role          Role           `gorm:"foreignkey:Id_role"`
	Id_permission uint           `json:"id_permission"`
	Permission    Permission     `gorm:"foreignkey:Id_permission"`
}

type Task struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Id_user     uint           `json:"id_user"`
	User        User           `gorm:"foreignkey:Id_user"`
}
