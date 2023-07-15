package db

import (
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"

	models "greenbone-task/internal/persistence/models/computer"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

// Setup opens a database and saves the reference to `Database` struct.
func Setup(logger *zap.Logger) {
	var db = DB
	dbLogger := zapgorm2.New(logger)
	dbLogger.SetAsDefault()

	database := viper.GetString("database_name")
	username := viper.GetString("database_username")
	password := viper.GetString("database_password")
	host := viper.GetString("database_host")
	port := viper.GetString("database_port")

	logger.Info("Database: using postgres driver for connecting.")
	dsn := "host=" + host + " port=" + port + " user=" + username + " dbname=" + database + "  sslmode=disable password=" + password
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		logger.Error("db err: ", zap.Error(err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("db err: ", zap.Error(err))
	}

	sqlDB.SetMaxIdleConns(viper.GetInt("database_idle_conns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("database_open_conns"))
	sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt("database_max_lifetime")) * time.Second)

	DB = db
	migration(logger)
}

// Auto migrate project models
func migration(logger *zap.Logger) {
	logger.Info("Database: creating service model tables")
	DB.AutoMigrate(&models.Computer{})
}

func Get() *gorm.DB {
	return DB
}
