package mysql

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	onceGormDB sync.Once
	db         *gorm.DB
)

func NewMySQLConnection(ctx context.Context) (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}
	onceGormDB.Do(func() {
		cfg := getConfig()

		fmt.Println("CFG", cfg)

		// Log the MySQL configuration for debugging
		log.Printf("Connecting to MySQL with Config: User=%s, Password=%s, Host=%s, Port=%s, Name=%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

		// Build the DSN (Data Source Name)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

		fmt.Println("Preparing", dsn)
		var err error
		db, err = gorm.Open(mysqlDriver.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn), // Change log mode if you want to debug in local
		})
		if err != nil {
			err = errors.Wrapf(err, "gorm open data source: %s", dsn)
			log.Fatal(err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			err = errors.Wrap(err, "get DB")
			log.Fatal(err)
		}

		// Set database connection pool configurations
		dmic, err := strconv.Atoi(cfg.ConnMaxIdle)
		if err != nil {
			err = errors.Wrapf(err, "convert ConnMaxIdle %s", cfg.ConnMaxIdle)
			log.Fatal(err)
		}
		sqlDB.SetMaxIdleConns(dmic)

		dmoc, err := strconv.Atoi(cfg.ConnMaxOpen)
		if err != nil {
			err = errors.Wrapf(err, "convert ConnMaxOpen %s", cfg.ConnMaxOpen)
			log.Fatal(err)
		}
		sqlDB.SetMaxOpenConns(dmoc)

		dcmlt, err := strconv.Atoi(cfg.ConnMaxLifeTime)
		if err != nil {
			err = errors.Wrapf(err, "convert ConnMaxLifeTime %s", cfg.ConnMaxLifeTime)
			log.Fatal(err)
		}
		sqlDB.SetConnMaxLifetime(time.Duration(dcmlt) * time.Second)

		dcmit, err := strconv.Atoi(cfg.ConnMaxIdleTime)
		if err != nil {
			err = errors.Wrapf(err, "convert ConnMaxIdleTime %s", cfg.ConnMaxIdleTime)
			log.Fatal(err)
		}
		sqlDB.SetConnMaxIdleTime(time.Duration(dcmit) * time.Second)
	})
	return db, nil
}

// mySQLConfig holds the configuration for MySQL connection
type mySQLConfig struct {
	User            string
	Password        string
	Name            string
	Port            string
	Host            string
	ConnMaxIdle     string
	ConnMaxOpen     string
	ConnMaxLifeTime string
	ConnMaxIdleTime string
}

func getConfig() mySQLConfig {
	return mySQLConfig{
		User:            getEnv("DATABASE_USER", "default_user"),
		Password:        getEnv("DATABASE_PASSWORD", ""),
		Name:            getEnv("DATABASE_NAME", "default_db"),
		Port:            getEnv("DATABASE_PORT", "3306"),
		Host:            getEnv("DATABASE_HOSTNAME", "127.0.0.1"),
		ConnMaxIdle:     getEnv("DATABASE_MAX_IDLE_CONNS", "10"),
		ConnMaxOpen:     getEnv("DATABASE_MAX_OPEN_CONNS", "100"),
		ConnMaxLifeTime: getEnv("DATABASE_CONN_MAX_LIFE_TIME", "6"),
		ConnMaxIdleTime: getEnv("DATABASE_CONN_MAX_IDLE_TIME", "2"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
