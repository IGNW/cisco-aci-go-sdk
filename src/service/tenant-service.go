package service

import (
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var tenantServiceInstance *TenantService

type TenantService struct {
	ResourceService
}

func GetTenantService(client *Client) *TenantService {
	if tenantServiceInstance == nil {
		tenantServiceInstance = &TenantService{ResourceService{
			ObjectClass:        models.TENANT_OBJECT_CLASS,
			ResourceNamePrefix: models.TENANT_RESOURCE_PREFIX,
		}}
	}
	return tenantServiceInstance
}

/* New creates a new Tenant with the appropriate default values */
func (ts TenantService) New(name string, description string) *models.Tenant {
	t := models.Tenant{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  models.TENANT_OBJECT_CLASS,
		ResourceName: ts.getResourceName(name),
	},
		nil,
		nil,
		nil,
		nil,
		nil,
	}
	//Do any additional construction logic here.
	return &t
}

func (ts TenantService) Save(t *models.Tenant) error {

	err := ts.ResourceService.Save(t)
	if err != nil {
		return err
	}

	return nil

}

func (ts TenantService) Get(domainName string) (*models.Tenant, error) {

	data, err := ts.ResourceService.Get(domainName)

	if err != nil {
		return nil, err
	}

	return ts.fromJSON(data)
}

func (ts TenantService) GetById(id string) (*models.Tenant, error) {
	data, err := ts.ResourceService.GetById(id)

	if err != nil {
		return nil, err
	}

	return ts.fromJSON(data)
}

func (ts TenantService) GetByName(name string) ([]*models.Tenant, error) {
	data, err := ts.ResourceService.GetByName(name)

	if err != nil {
		return nil, err
	}

	return ts.fromDataArray(data)
}

func (ts TenantService) GetAll() ([]*models.Tenant, error) {

	data, err := ts.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	return ts.fromDataArray(data)
}

func (ts TenantService) fromDataArray(data []*gabs.Container) ([]*models.Tenant, error) {
	var tenants []*models.Tenant
	var newTenant *models.Tenant
	var err, errors error

	// For each tenant in the payload
	for _, fvTenant := range data {

		newTenant, err = ts.fromJSON(fvTenant)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			tenants = append(tenants, newTenant)

		}
	}

	return tenants, err
}

func (ts TenantService) fromJSON(data *gabs.Container) (*models.Tenant, error) {

	mapped, err := ts.fromJSONToMap(models.NewTenantMap(), data)

	if err != nil {
		return nil, err
	}

	// TODO: process child collections
	return models.NewTenant(mapped), nil

}
