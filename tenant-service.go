package main

var tenantServiceInstance *ResourceService

func GetTenantService(client *Client) *ResourceService {
	if tenantServiceInstance == nil {
		tenantServiceInstance = &ResourceService{
			ObjectClass: "fvTenant",
			New:         NewTenant,
			FromJSON:    TenantFromJSON,
		}
	}
	return tenantServiceInstance
}
