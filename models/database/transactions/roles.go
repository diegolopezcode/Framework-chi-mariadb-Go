package transactions

import (
	"fmt"

	database "github.com/diegolopezcode/api-crud-complete-chi/models/database/config"
	"github.com/diegolopezcode/api-crud-complete-chi/models/database/models"
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
