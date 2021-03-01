package server

import (
	"github.com/gin-gonic/gin"
	"github.com/laybatin/bodymirror-app-backend/controllers"
	v1 "github.com/laybatin/bodymirror-app-backend/internal/v1"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)
	router.GET("/version", v1.VersionHandler)

	return router

}
