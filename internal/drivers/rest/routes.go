package rest

import (
	"multitenancy/internal/tenants"

	"github.com/gin-gonic/gin"
)

func BuildRoutes(router *gin.Engine, tenantsDeps *tenants.DependencyTree) {
	router.GET("/ping", ping)

	tenantsPath := router.Group("/tenants")
	tenants.BuildRoutes(tenantsPath, tenantsDeps)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
