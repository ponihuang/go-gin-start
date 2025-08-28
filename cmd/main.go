package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ponihuang/go-gin-start/internal/config"
	"github.com/ponihuang/go-gin-start/internal/handlers"
	"github.com/ponihuang/go-gin-start/internal/model"
)

// setupRouter configures and returns a Gin engine.
func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.POST("/login", handlers.LoginHandler(db))
	uh := handlers.NewUserHandler(db)
	users := r.Group("/users")
	{
		users.GET("", uh.List)
		users.POST("", uh.Create)
		users.GET(":id", uh.Get)
		users.PUT(":id", uh.Update)
		users.DELETE(":id", uh.Delete)
	}
	return r
}

func main() {
	env := os.Getenv("APP_ENV")
	cfg, err := config.Load(env)
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	db, err := gorm.Open(sqlite.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}
	r := setupRouter(db)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
