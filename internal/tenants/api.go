package tenants

import (
	"multitenancy/internal/tenants/app"

	"github.com/gin-gonic/gin"
)

type TenantRouter struct {
	tenantsService *app.TenantService
}

func BuildRoutes(tenantsPath *gin.RouterGroup, deps *DependencyTree) {
	tenantRouter := &TenantRouter{tenantsService: deps.TenantService}

	tenantsPath.GET("", tenantRouter.listTenantsHandler)
}

func (t *TenantRouter) listTenantsHandler(ctx *gin.Context) {
	tenants, err := t.tenantsService.ListTenants(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"tenants": tenants})
}
