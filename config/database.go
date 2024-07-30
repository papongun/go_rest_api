package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/papongun/go_todo/entity"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func GetDatabase() *gorm.DB {
	dbOnce.Do(func() {
		dbInstance = initDatabase()
	})
	return dbInstance
}

func initDatabase() (db *gorm.DB) {
	erro := godotenv.Load("./config/.env")
	if erro != nil {
		panic("failed to retrive env")
	}

	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASS")
	dbname := os.Getenv("PG_DB")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&entity.User{})

	return db
}
