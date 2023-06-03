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
	Id_role   uint           `json:"id_role foreignkey:roles.id"`
	Email     string         `json:"email unique"`
}

type Role struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name unique"`
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
	Id_role       uint           `json:"id_role foreignkey:roles.id"`
	Id_permission uint           `json:"id_permission foreignkey:permissions.id"`
}

type Task struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Id_user     uint           `json:"id_user foreignkey:users.id"`
}
