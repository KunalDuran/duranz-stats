package data

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB instance
var DB *gorm.DB

// InitDB initializes the database connection
func InitDB(host, user, password, dbname string, port int) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")
	return nil
}

// CreateTables creates all the required tables in the database
func CreateTables() error {
	// List of all table structures
	tables := []interface{}{
		&CricketMatch{},
		&CricketPlayer{},
		&Team{},
		&Venue{},
		&MatchStats{},
		&PlayerMatchStats{},
		&ErrorLog{},
		&FileMapping{},
	}

	// Create tables
	for _, table := range tables {
		err := DB.AutoMigrate(table)
		if err != nil {
			return fmt.Errorf("failed to create table for %T: %v", table, err)
		}
	}

	log.Println("All tables created successfully")
	return nil
}

// TruncateTables deletes all data from the tables while maintaining the structure
func TruncateTables() error {
	// Disable foreign key checks
	err := DB.Exec("SET FOREIGN_KEY_CHECKS = 0").Error
	if err != nil {
		return fmt.Errorf("failed to disable foreign key checks: %v", err)
	}

	// List of tables in order (considering dependencies)
	tables := []interface{}{
		&PlayerMatchStats{},
		&MatchStats{},
		&CricketMatch{},
		&CricketPlayer{},
		&Team{},
		&Venue{},
		&ErrorLog{},
		&FileMapping{},
	}

	// Truncate each table
	for _, table := range tables {
		err := DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(table).Error
		if err != nil {
			return fmt.Errorf("failed to truncate table %T: %v", table, err)
		}
	}

	// Re-enable foreign key checks
	err = DB.Exec("SET FOREIGN_KEY_CHECKS = 1").Error
	if err != nil {
		return fmt.Errorf("failed to enable foreign key checks: %v", err)
	}

	log.Println("All tables truncated successfully")
	return nil
}
