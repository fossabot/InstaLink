package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/message/", func(c echo.Context) error {
		return c.String(http.StatusOK, "InstaLink is now under the construction.")
	})
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
	db.AutoMigrate(&MODEL_NAME{})
}
