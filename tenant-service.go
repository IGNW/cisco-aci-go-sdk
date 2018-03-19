package main

var tenantServiceInstance *ResourceService

func GetTenantService() *ResourceService {
	if tenantServiceInstance == nil {
		tenantServiceInstance = &ResourceService{
			ObjectClass: "fvTenant",
			New:         NewTenant,
			FromJSON:    TenantFromJSON,
		}
	}
	return tenantServiceInstance
}
