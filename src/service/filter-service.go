package service

import (
	"fmt"
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var filterServiceInstance *FilterService

type FilterService struct {
	ResourceService
}

func GetFilterService(client *Client) *FilterService {
	if filterServiceInstance == nil {
		filterServiceInstance = &FilterService{ResourceService{
			ObjectClass: "@TODO",
		}}
	}
	return filterServiceInstance
}

/* New creates a new Filter with the appropriate default values */
func (fs FilterService) New(name string, description string) *models.Filter {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	b := models.Filter{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  "@TODO",
		ResourceName: resourceName,
	},
		nil,
		nil,
		nil,
	}
	//Do any additional construction logic here.
	return &b
}

func (fs FilterService) Save(f *models.Filter) error {

	err := fs.ResourceService.Save(f)
	if err != nil {
		return err
	}

	return nil

}

func (fs FilterService) Get(domainName string) (*models.Filter, error) {

	data, err := fs.ResourceService.Get(domainName)

	if err != nil {
		return nil, err
	}

	newFilter, err := fs.fromJSON(data)

	if err != nil {
		return nil, err
	}

	return newFilter, nil
}

func (fs FilterService) GetAll() ([]*models.Filter, error) {
	var epgs []*models.Filter
	var errors error
	data, err := fs.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	fvFilters, err := data.S("imdata").Children()
	if err != nil {
		return nil, err
	}

	// For each epg in the payload
	for _, fvFilter := range fvFilters {

		newFilter, err := fs.fromJSON(fvFilter)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			epgs = append(epgs, newFilter)

		}
	}

	return epgs, err
}

func (fs FilterService) fromJSON(data *gabs.Container) (*models.Filter, error) {
	var errors error
	var valPath, errMsg, name, desc string
	var ok bool

	errMsg = "Could not find value '%s' within child of imdata"
	valPath = ""

	valPath = "@TODO.attributfs.name"
	if name, ok = data.Path(valPath).Data().(string); !ok {
		errors = multierror.Append(errors, fmt.Errorf(errMsg, valPath))
	}

	valPath = "@TODO.attributfs.descr"
	if desc, ok = data.Path(valPath).Data().(string); !ok {
		errors = multierror.Append(errors, fmt.Errorf(errMsg, valPath))
	}

	if errors != nil {
		return nil, errors
	}

	newFilter := fs.New(name, desc)
	return newFilter, nil
}
