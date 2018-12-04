package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/shotastage/instalink/models"
)

func main() {

	e := InitRouter()

	e.Logger.Fatal(e.Start(":1323"))

	setupDB()
}

func setupDB() {

	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(models.URLEntry{})
}
