package postgres

import (
	"fmt"
	"likes/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabse() error {
	env := config.ConfigEnvs
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		env.PostgresDB.Host,
		env.PostgresDB.User,
		env.PostgresDB.Password,
		env.PostgresDB.Name,
		env.PostgresDB.Port,
		env.PostgresDB.SSLMode)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
