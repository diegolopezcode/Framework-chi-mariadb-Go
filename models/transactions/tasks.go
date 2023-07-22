package transactions

import (
	"strconv"

	database "github.com/diegolopezcode/api-crud-complete-chi/models/database/config"
	"github.com/diegolopezcode/api-crud-complete-chi/models/models"
)

type Tasks struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Id_user     uint   `json:"id_user"`
	Is_complete bool   `json:"is_complete"`
}

type TasksUpdate struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Id_user     uint   `json:"id_user"`
	Is_complete string `json:"is_complete"`
}

func CreateTask(task Tasks) (*models.Task, error) {
	con := database.Connect()
	taskModel := &models.Task{
		Name:        task.Name,
		Description: task.Description,
		Id_user:     task.Id_user,
		Is_complete: task.Is_complete,
	}
	err := con.Create(taskModel)
	if err.Error != nil {
		return nil, err.Error
	}
	return taskModel, nil

}

func GetAllTasks() ([]*models.Task, error) {
	con := database.Connect()
	task := []*models.Task{}
	err := con.Find(&task)
	if err.Error != nil {
		return nil, err.Error
	}
	return task, nil
}

func GetTaskById(id int) (*models.Task, error) {
	con := database.Connect()
	task := &models.Task{
		ID: uint(id),
	}
	err := con.First(task)
	if err.Error != nil {
		return nil, err.Error
	}
	return task, nil

}

func GetTaskByUserId(id int) ([]*models.Task, error) {
	con := database.Connect()
	task := []*models.Task{}
	err := con.Where("id_user = ?", id).Find(&task)
	if err.Error != nil {
		return nil, err.Error
	}
	return task, nil

}

func GetTaskByState(state bool) ([]*models.Task, error) {
	con := database.Connect()
	task := []*models.Task{}
	err := con.Where("is_complete = ?", state).Find(&task)
	if err.Error != nil {
		return nil, err.Error
	}
	return task, nil

}

func GetTaskByStateAndUserId(state bool, id int) ([]*models.Task, error) {
	con := database.Connect()
	task := []*models.Task{}
	err := con.Where("is_complete = ? and id_user = ?", state, id).Find(&task)
	if err.Error != nil {
		return nil, err.Error
	}
	return task, nil

}

func UpdateTask(task TasksUpdate) (*models.Task, error) {
	con := database.Connect()
	data, err := strconv.ParseBool(task.Is_complete)
	if err != nil {
		return nil, err
	}

	findTask := &models.Task{
		ID:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		Id_user:     task.Id_user,
		Is_complete: data,
	}
	errs := con.Save(findTask)
	if errs.Error != nil {
		return nil, errs.Error
	}
	return findTask, nil
}

func DeleteTask(id int) error {
	con := database.Connect()
	task := &models.Task{
		ID: uint(id),
	}
	err := con.Delete(task)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
