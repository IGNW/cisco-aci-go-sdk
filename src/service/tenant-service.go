package service

import (
	"fmt"
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var tenantServiceInstance *TenantService

const RESOURCE_NAME_PREFIX = "tn"
const OBJECT_CLASS = "fvTenant"

type TenantService struct {
	ResourceService
}

func GetTenantService(client *Client) *TenantService {
	if tenantServiceInstance == nil {
		tenantServiceInstance = &TenantService{ResourceService{
			ObjectClass:        OBJECT_CLASS,
			ResourceNamePrefix: RESOURCE_NAME_PREFIX,
		}}
	}
	return tenantServiceInstance
}

/* New creates a new Tenant with the appropriate default values */
func (ts TenantService) New(name string, description string) *models.Tenant {
	resourceName := fmt.Sprintf("tn-%s", name)

	t := models.Tenant{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  OBJECT_CLASS,
		ResourceName: resourceName,
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
	/*
		var errors error
		var valPath, errMsg, name, desc string
		var ok bool

		errMsg = "Could not find value '%s' within child of imdata"
		valPath = ""

		valPath = "fvTenant.attributes.name"
		if name, ok = data.Path(valPath).Data().(string); !ok {
			errors = multierror.Append(errors, fmt.Errorf(errMsg, valPath))
		}

		valPath = "fvTenant.attributes.descr"
		if desc, ok = data.Path(valPath).Data().(string); !ok {
			errors = multierror.Append(errors, fmt.Errorf(errMsg, valPath))
		}

		if errors != nil {
			return nil, errors
		}



		newTenant := ts.New(name, desc)
		return newTenant, nil
	*/

	resourceAttributes, err := ts.fromJSONToAttributes(ts.ObjectClass, data)

	if err != nil {
		return nil, err
	}

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
