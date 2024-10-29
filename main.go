package main

import (
	"fmt"

	database "github.com/ShabnamHaque/go-crm/database"
	lead "github.com/ShabnamHaque/go-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead) //takes an ID
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect DB")
	}
	fmt.Println("connection opened to DB")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}
func main() {
	//fmt.Println("yeahhh")
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3030)
	defer database.DBConn.Close()

}
