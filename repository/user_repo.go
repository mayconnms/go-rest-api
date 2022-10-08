package repository

import (
	"log"
	"rest-api/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getConnection() *gorm.DB {

	dsn := "host=localhost user=postgres password=admin dbname=bank port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db

}

func FindUser(id uint) *domain.User {
	db := getConnection()
	var user domain.User
	db.First(&user, id)
	return &user
}

func DeleteUser(id uint) {
	db := getConnection()
	user := FindUser(id)
	db.Delete(&user)
}

func SaveUser(user *domain.User) {
	db := getConnection()
	db.Save(&user)
}
