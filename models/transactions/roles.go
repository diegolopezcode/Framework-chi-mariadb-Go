package transactions

import (
	"fmt"

	database "github.com/diegolopezcode/api-crud-complete-chi/models/database/config"
	"github.com/diegolopezcode/api-crud-complete-chi/models/models"
)

func CreateRole(name string) (*models.Role, error) {
	fmt.Println("Llega Name", name)
	con := database.Connect()
	role := &models.Role{Name: name}
	err := con.Create(role)
	if err.Error != nil {
		return nil, err.Error
	}

	return role, nil
}

func GetRoles() ([]*models.Role, error) {
	fmt.Println("Llega Solo")
	con := database.Connect()
	roles := []*models.Role{}
	err := con.Find(&roles)
	if err.Error != nil {
		return nil, err.Error
	}

	return roles, nil
}

func GetRoleById(id int) (*models.Role, error) {
	fmt.Println("Llega id", id)
	con := database.Connect()
	role := &models.Role{
		ID: uint(id),
	}
	err := con.First(role)
	if err.Error != nil {
		return nil, err.Error
	}
	fmt.Println("Llega role", role)
	return role, nil
}

type Roles struct {
	Id   uint
	Name string
}

func UpdateRole(id int, name string) (*models.Role, error) {
	fmt.Println("Llega con id and anme", id, name)
	con := database.Connect()
	findRole := &models.Role{ID: uint(id),
		Name: name}
	err := con.Save(findRole)
	if err.Error != nil {
		return nil, err.Error
	}

	fmt.Println("Llega role", findRole)
	return findRole, nil
}
