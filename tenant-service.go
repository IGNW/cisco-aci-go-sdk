package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
)

var tenantServiceInstance *TenantService

type TenantService struct {
	ResourceService
}

func GetTenantService(client *Client) *TenantService {
	if tenantServiceInstance == nil {
		tenantServiceInstance = &TenantService{ResourceService{
			ObjectClass: "fvTenant",
		}}
	}
	return tenantServiceInstance
}

/* New creates a new Tenant with the appropriate default values */
func (ts TenantService) New(name string, description string) *Tenant {
	resourceName := fmt.Sprintf("tn-%s", name)

	t := Tenant{ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  "fvTenant",
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

func (ts TenantService) Save(t *Tenant) error {

	err := ts.ResourceService.Save(t)
	if err != nil {
		return err
	}

	return nil

}

func (ts TenantService) Get(params *map[string]string) (*[]Tenant, error) {

	data, err := ts.ResourceService.Get(params)

	fmt.Printf("DATA: %v\n\n", data)
	fmt.Printf("ERR : %s\n\n", err)

	return nil, nil
}

func (ts TenantService) GetAll() ([]*Tenant, error) {
	var tenants []*Tenant
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

func (ts TenantService) fromJSON(data *gabs.Container) (*Tenant, error) {
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
}
