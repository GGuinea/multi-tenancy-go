package tenants

import (
	"multitenancy/internal/tenants/app"
	"multitenancy/internal/tenants/domain"

	"github.com/gin-gonic/gin"
)

type TenantRouter struct {
	tenantsService *app.TenantService
}

func BuildRoutes(tenantsPath *gin.RouterGroup, deps *DependencyTree) {
	tenantRouter := &TenantRouter{tenantsService: deps.TenantService}

	tenantsPath.GET("", tenantRouter.listTenantsHandler)
	tenantsPath.POST("", tenantRouter.addTenantHandler)
	tenantsPath.GET("/:id", tenantRouter.getTenantHandler)
}

func (t *TenantRouter) listTenantsHandler(ctx *gin.Context) {
	tenants, err := t.tenantsService.ListTenants(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := make([]domain.TenantResponseDto, len(tenants))

	for i, tenant := range tenants {
		response[i] = domain.TenantResponseDto{
			ID:     tenant.ID,
			Name:   tenant.Name,
			Active: tenant.Active,
		}
	}

	ctx.JSON(200, gin.H{"tenants": response})
}

func (t *TenantRouter) addTenantHandler(ctx *gin.Context) {
	var tenant domain.TenantRequestDto

	err := ctx.ShouldBindJSON(&tenant)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newTenantUUID, err := t.tenantsService.AddTenant(ctx, tenant)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"uuid": newTenantUUID})
}

func (t *TenantRouter) getTenantHandler(ctx *gin.Context) {
	tenantId := ctx.Param("id")

	tenant, err := t.tenantsService.ReadById(ctx, tenantId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := domain.TenantResponseDto{
		ID:     tenant.ID,
		Name:   tenant.Name,
		Active: tenant.Active,
	}

	ctx.JSON(200, gin.H{"tenant": response})
}
