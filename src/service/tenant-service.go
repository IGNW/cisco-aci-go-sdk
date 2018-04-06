package service

import (
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var tenantServiceInstance *TenantService

const TN_RESOURCE_NAME_PREFIX = "tn"
const TN_OBJECT_CLASS = "fvTenant"

type TenantService struct {
	ResourceService
}

func GetTenantService(client *Client) *TenantService {
	if tenantServiceInstance == nil {
		tenantServiceInstance = &TenantService{ResourceService{
			ObjectClass:        TN_OBJECT_CLASS,
			ResourceNamePrefix: TN_RESOURCE_NAME_PREFIX,
			HasParent:          false,
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
		ObjectClass:  TN_OBJECT_CLASS,
		ResourceName: ts.getResourceName(name),
	},
		"",
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

func (ts TenantService) Get(name string) (*models.Tenant, error) {

	data, err := ts.ResourceService.Get(name)

	if err != nil {
		return nil, err
	}

	// TODO: reafacor single item or array behavior of imdata into resource service
	fvTenants, err := data.S("imdata").Children()
	if err != nil {
		return nil, err
	}

	newTenant, err := ts.fromJSON(fvTenants[0])

	if err != nil {
		return nil, err
	}

	return newTenant, nil
}

func (ts TenantService) GetAll() ([]*models.Tenant, error) {
	var tenants []*models.Tenant
	var errors error
	data, err := ts.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	fvTenants, err := data.S("imdata").Children()
	if err != nil {
		return nil, err
	}

	// For each tenant in the payload
	for _, fvTenant := range fvTenants {

		newTenant, err := ts.fromJSON(fvTenant)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			tenants = append(tenants, newTenant)

		}
	}

	return tenants, err
}

func (ts TenantService) fromJSON(data *gabs.Container) (*models.Tenant, error) {

	resourceAttributes, err := ts.fromJSONToAttributes(ts.ObjectClass, data)

	if err != nil {
		return nil, err
	}

	// TODO: process child collections

	return &models.Tenant{
		resourceAttributes,
		"",
		nil,
		nil,
		nil,
		nil,
		nil,
	}, nil

}
