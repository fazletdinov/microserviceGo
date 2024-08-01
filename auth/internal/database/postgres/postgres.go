package postgres

import (
	"auth/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

func NewPostgresClient(env *config.Config) (*gorm.DB, error) {
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		env.PostgresDB.Host,
		env.PostgresDB.User,
		env.PostgresDB.Password,
		env.PostgresDB.Name,
		env.PostgresDB.Port,
		env.PostgresDB.SSLMode)
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
	return db, nil
}
