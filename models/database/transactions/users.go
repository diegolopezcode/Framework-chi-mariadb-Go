package transactions

import (
	database "github.com/diegolopezcode/api-crud-complete-chi/models/database/config"
	"github.com/diegolopezcode/api-crud-complete-chi/models/database/models"
)

type Users struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Id_role  uint   `json:"id_role"`
}

type UsersUpdate struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Id_role  uint   `json:"id_role"`
}

func CreateUser(users Users) (*models.User, error) {
	con := database.Connect()
	user := &models.User{
		Name:     users.Name,
		Password: users.Password,
		Email:    users.Email,
		Id_role:  users.Id_role,
	}
	err := con.Create(user)
	if err.Error != nil {
		return nil, err.Error
	}
	return user, nil

}

func GetUserById(id int) (*models.User, error) {
	con := database.Connect()
	user := &models.User{
		ID: uint(id),
	}
	err := con.First(user)
	if err.Error != nil {
		return nil, err.Error
	}
	return user, nil

}

func UpdateUser(user UsersUpdate) (*models.User, error) {
	con := database.Connect()
	findUser := &models.User{
		ID:       user.Id,
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
		Id_role:  user.Id_role,
	}
	err := con.Save(findUser)
	if err.Error != nil {
		return nil, err.Error
	}
	return findUser, nil
}

func GetAllUsers() ([]*models.User, error) {
	con := database.Connect()
	users := []*models.User{}
	err := con.Find(&users)
	if err.Error != nil {
		return nil, err.Error
	}
	return users, nil
}

func GetUsersByRoleId(
	id_role int,
) ([]*models.User, error) {
	con := database.Connect()
	users := []*models.User{}
	err := con.Where("id_role = ?", id_role).Find(&users)
	if err.Error != nil {
		return nil, err.Error
	}
	return users, nil
}
