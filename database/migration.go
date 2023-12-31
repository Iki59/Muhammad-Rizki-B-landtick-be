package database

import (
	"fmt"
	"landtick-be/models"
	"landtick-be/pkg/mysql"
)

// func RunMigration() {
// 	err := mysql.DB.AutoMigrate(&models.User{})

// 	if err != nil {
// 		fmt.Println(err)
// 		panic("Migration Failed")
// 	}

// 	fmt.Println("Migration Success")
// }

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Station{},
		&models.Ticket{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
