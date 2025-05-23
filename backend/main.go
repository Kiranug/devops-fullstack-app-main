package main

import (
	"employees/controller"
	"employees/repository"
	"employees/routes"
	"employees/service" // ✅ Add this line
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)


func main() {
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Initialize database
	db := initializeDatabaseConnection()
	repository.RunMigrations(db)

	// Setup dependencies
	employeeRepository := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeController := controller.NewEmployeeController(employeeService)

	// Register routes
	routes.RegisterRoute(app, employeeController)

	// Start the server
	err := app.Listen(":8080")
	if err != nil {
		log.Fatalln(fmt.Sprintf("error starting the server %s", err.Error()))
	}
}

func initializeDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  createDsn(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(fmt.Sprintf("error connecting with database %s", err.Error()))
	}
	return db
}

func createDsn() string {
	// Hardcoded values for testing
	dbHost := "172.174.2.135"
	dbUser := "myuser"
	dbPassword := "mypassword"
	dbName := "myappdb"
	dbPort := "5432"

	dsnFormat := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	return fmt.Sprintf(dsnFormat, dbHost, dbUser, dbPassword, dbName, dbPort)
}
