package main


import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/DLLenjoyer/TrollNetGram/server/models"
)

func main() {
	r := gin.Default()

	db := initDB()
	migrateDB(db)
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run("1043")
	if err != nil {
		log.Fatal("Не получилось стартануть сервер", err)
	}
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test123.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Не получлось законектится с базой данных", err)
	}
	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
