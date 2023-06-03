package database

import (
	"fmt"

	"github.com/diegolopezcode/api-crud-complete-chi/configs"
	"github.com/diegolopezcode/api-crud-complete-chi/models/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Table struct {
	TableName string
	// more fields if needed...&loc=Loca
}

func Connect() *gorm.DB {
	conn_url := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configs.Config("user"), configs.Config("password"), configs.Config("port"), configs.Config("name"))
	db, err := gorm.Open(mysql.Open(conn_url), &gorm.Config{})
	defer recover()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	// Migrate the schema

	db.AutoMigrate(
		models.Role{},
		models.Permission{},
		models.RolePermission{},
		models.User{},
		models.Task{},
	)

	// roles := []models.Role{
	// 	{Name: "admin"},
	// 	{Name: "user"},
	// }
	// db.Create(&roles)
	// fmt.Println(roles)

	// fmt.Println("db.Migrator().CurrentDatabase()", db.Migrator())
	// fmt.Println("Database Migrated")
	fmt.Println("Database Migrated")
	return db
}

// func Connect() *gorm.DB {
// 	conn_url := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configs.Config("user"), configs.Config("password"), configs.Config("port"), configs.Config("name"))
// 	db, err := gorm.Open(mysql.Open(conn_url), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(err)
// 		panic("failed to connect database")
// 	}
// 	fmt.Println("Connection Opened to Database")
// 	// Migrate the schema
// 	db.AutoMigrate(
// 		models.Role{},
// 		models.Permission{},
// 		models.RolePermission{},
// 		models.User{},
// 		models.Task{},
// 	)
// 	fmt.Println("Database Migrated")
// 	return db
// }
