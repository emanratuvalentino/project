package db

import (
	"code-be-docudigital/config"
	"fmt"

	"github.com/jinzhu/gorm"
	// For postgres connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// Init DB Connection
func Init() {
	configuration := config.GetConfig()

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", configuration.DB_HOST, configuration.DB_PORT, configuration.DB_USERNAME, configuration.DB_NAME, configuration.DB_PASSWORD)
	db, err = gorm.Open("postgres", connectString)

	// defer db.Close()
	if err != nil {
		panic("DB Connection Error cuy")
	}

	// db.AutoMigrate(&model.User{})

}

// Manager for conn
func Manager() *gorm.DB {
	return db
}

// ManagerRetrieve return db connection
func ManagerRetrieve() *gorm.DB {
	return db
}
