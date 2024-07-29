package router

import (
	"github.com/gin-gonic/gin"
	"shortener/auth"
	"shortener/shortener"
)

func SetupRouter(cfg *Config) *gin.Engine {
	r := gin.Default()

	// auth
	authService := auth.NewService(cfg.SecretKey, cfg.Tarantool)
	r.Use(authService.ContextUserMiddleware())
	r.POST("/register", authService.Register)
	r.POST("/login", authService.Login)

	// links
	shortenerService := shortener.NewService(cfg.Tarantool)
	r.POST("/shortener", shortenerService.CreateShortLink)
	r.GET("/:shortURL", shortenerService.RedirectLink)
	return r
}
