package initializers

import (
	database "airlift/internal/connections"
	"airlift/schemas"
	"fmt"
)

func MakeMigrations() {
	if database.DB == nil {
		panic("database connection is not initialized")
	}

	err := database.DB.AutoMigrate(
		&schemas.Project{},
	)
	if err != nil {
		panic(fmt.Sprintf("cannot auto-migrate database: %s", err))
	}
}
