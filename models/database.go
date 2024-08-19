package models

import (
	"chesscaster-gin/helper"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the he database connection.
var DB *gorm.DB

// SetupDatabase migrates and sets up the database.
func SetupDatabase() {
	user := helper.GetEnv("DATABASE_USER", "chesscaster")
	pw := helper.GetEnv("DATABASE_PASSWORD", "")
	host := helper.GetEnv("DATABASE_HOST", "localhost:5432")
	dbName := helper.GetEnv("DATABASE_NAME", "chesscaster")

	// Assemble the connection string.
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s/%s", user, pw, host, dbName)

	// Connect to the database.
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	// Migrate the schema
	db.AutoMigrate(&Game{})

	if err != nil {
		panic("Could not open database connection")
	}

	DB = db
}
