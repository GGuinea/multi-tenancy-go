package rest

import (
	"multitenancy/internal/notes"
	"multitenancy/internal/tenants"

	"github.com/gin-gonic/gin"
)

func BuildRoutes(router *gin.Engine, tenantsDeps *tenants.DependencyTree, notesDeps *notes.DependencyTree) {
	router.GET("/ping", ping)

	tenantsPath := router.Group("/tenants")
	tenants.BuildRoutes(tenantsPath, tenantsDeps)

	notesPath := router.Group("/notes")
	notes.BuildRoutes(notesPath, notesDeps)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
