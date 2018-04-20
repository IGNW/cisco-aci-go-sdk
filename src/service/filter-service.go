package service

import (
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
			ObjectClass:        models.FILTER_OBJECT_CLASS,
			ResourceNamePrefix: models.FILTER_RESOURCE_PREFIX,
			HasParent:          true,
		}}
	}
	return filterServiceInstance
}

// New creates a new Filter with the appropriate default values.
func (fs FilterService) New(name string, description string) *models.Filter {

	b := models.Filter{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  models.FILTER_OBJECT_CLASS,
		ResourceName: fs.getResourceName(name),
	},
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

func (fs FilterService) GetByName(name string) ([]*models.Filter, error) {

	data, err := fs.ResourceService.GetByName(name)
	if err != nil {
		return nil, err
	}

	return fs.fromDataArray(data)
}

func (fs FilterService) GetAll() ([]*models.Filter, error) {

	data, err := fs.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	return fs.fromDataArray(data)
}

func (fs FilterService) fromDataArray(data []*gabs.Container) ([]*models.Filter, error) {
	var epgs []*models.Filter
	var err, errors error

	// For each epg in the payload
	for _, fvFilter := range data {

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
	resourceAttributes, err := fs.fromJSONToAttributes(fs.ObjectClass, data)

	if err != nil {
		return nil, err
	}

	// TODO: process child collections

	return &models.Filter{
		resourceAttributes,
		nil,
		nil,
	}, nil
}
