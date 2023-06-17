package transactions

import (
	database "github.com/diegolopezcode/api-crud-complete-chi/models/database/config"
	"github.com/diegolopezcode/api-crud-complete-chi/models/database/models"
)

func CreateRolePermission(roleId int, permissionId int) (*models.RolePermission, error) {
	con := database.Connect()
	rolePermission := &models.RolePermission{Id_role: uint(roleId), Id_permission: uint(permissionId)}
	err := con.Create(rolePermission)
	if err.Error != nil {
		return nil, err.Error
	}

	return rolePermission, nil
}

func GetRolePermissions() ([]*models.RolePermission, error) {
	con := database.Connect()
	rolePermission := []*models.RolePermission{}
	err := con.Find(&rolePermission)
	if err.Error != nil {
		return nil, err.Error
	}

	return rolePermission, nil
}

func GetRolePermissionById(id int) (*models.RolePermission, error) {
	con := database.Connect()
	rolePermission := &models.RolePermission{
		ID: uint(id),
	}
	err := con.First(rolePermission)
	if err.Error != nil {
		return nil, err.Error
	}
	return rolePermission, nil
}

type RolePermissions struct {
	Id            uint
	Id_role       uint
	Id_permission uint
}

func UpdateRolePermission(id int, roleId int, permissionId int) (*models.RolePermission, error) {
	con := database.Connect()
	findRolePermission := &models.RolePermission{ID: uint(id),
		Id_role: uint(roleId), Id_permission: uint(permissionId)}
	err := con.Save(findRolePermission)
	if err.Error != nil {
		return nil, err.Error
	}

	return findRolePermission, nil
}

func DeleteRolePermission(id int) error {
	con := database.Connect()
	rolePermission := &models.RolePermission{ID: uint(id)}
	err := con.Delete(rolePermission)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func GetRolePermissionByRoleId(id int) ([]*models.RolePermission, error) {
	con := database.Connect()
	rolePermission := []*models.RolePermission{}
	err := con.Where("id_role = ?", id).Find(&rolePermission)
	if err.Error != nil {
		return nil, err.Error
	}

	return rolePermission, nil
}

func GetRolePermissionByPermissionId(id int) ([]*models.RolePermission, error) {
	con := database.Connect()
	rolePermission := []*models.RolePermission{}
	err := con.Where("id_permission = ?", id).Find(&rolePermission)
	if err.Error != nil {
		return nil, err.Error
	}

	return rolePermission, nil
}

func GetRolePermissionByRoleIdAndPermissionId(roleId int, permissionId int) ([]*models.RolePermission, error) {
	con := database.Connect()
	rolePermission := []*models.RolePermission{}
	err := con.Where("id_role = ? AND id_permission = ?", roleId, permissionId).Find(&rolePermission)
	if err.Error != nil {
		return nil, err.Error
	}

	return rolePermission, nil
}
