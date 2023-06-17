package transactions

import (
	database "github.com/diegolopezcode/api-crud-complete-chi/models/database/config"
	"github.com/diegolopezcode/api-crud-complete-chi/models/database/models"
)

func CreatePermission(name string) (*models.Permission, error) {
	con := database.Connect()
	permission := &models.Permission{Name: name}
	err := con.Create(permission)
	if err.Error != nil {
		return nil, err.Error
	}

	return permission, nil
}

func GetPermissions() ([]*models.Permission, error) {
	con := database.Connect()
	permission := []*models.Permission{}
	err := con.Find(&permission)
	if err.Error != nil {
		return nil, err.Error
	}

	return permission, nil
}

func GetPermissionById(id int) (*models.Permission, error) {
	con := database.Connect()
	permission := &models.Permission{
		ID: uint(id),
	}
	err := con.First(permission)
	if err.Error != nil {
		return nil, err.Error
	}
	return permission, nil
}

type Permissions struct {
	Id   uint
	Name string
}

func UpdatePermission(id int, name string) (*models.Permission, error) {
	con := database.Connect()
	findPermission := &models.Permission{ID: uint(id),
		Name: name}
	err := con.Save(findPermission)
	if err.Error != nil {
		return nil, err.Error
	}

	return findPermission, nil
}
