// db/connection.go

package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBOptions represents the options for configuring the database connection.
type DBOptions struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
	TimeZone string
	// Add more options as needed
}

// ConnectDB connects to the PostgreSQL database with the provided options and returns a *gorm.DB instance.
func ConnectDB(options DBOptions) (*gorm.DB, error) {
	dsn := "user=" + options.Username +
		" password=" + options.Password +
		" dbname=" + options.DBName +
		" host=" + options.Host +
		" port=" + options.Port +
		" sslmode=" + options.SSLMode +
		" TimeZone=" + options.TimeZone

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
