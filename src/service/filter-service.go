package service

import (
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var filterServiceInstance *FilterService

// FilterService is used to manage Filter resources.
type FilterService struct {
	ResourceService
}

// GetFilterService will construct or return the singleton for the FilterService.
func GetFilterService(client *Client) *FilterService {
	if filterServiceInstance == nil {
		filterServiceInstance = &FilterService{ResourceService{
			ObjectClass:        models.FILTER_OBJECT_CLASS,
			ResourceNamePrefix: models.FILTER_RESOURCE_PREFIX,
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

// Save a new Filter or update an existing one.
func (fs FilterService) Save(f *models.Filter) (string, error) {

	dn, err := fs.ResourceService.Save(f)
	if err != nil {
		return "", err
	}

	return dn, nil

}

// Get will retrieve an Filter by it's domain name.
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

// GetById will retrieve an Filter by it's unique identifier.
//TODO: Add

// GetByName will retrieve Filter(s) by common name.
func (fs FilterService) GetByName(name string) ([]*models.Filter, error) {

	data, err := fs.ResourceService.GetByName(name)
	if err != nil {
		return nil, err
	}

	return fs.fromDataArray(data)
}

// GetByName will retrieve all Filter(s).
func (fs FilterService) GetAll() ([]*models.Filter, error) {

	data, err := fs.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	return fs.fromDataArray(data)
}

// fromDataArray will convert an array of gabs.Container (JSON) to Filter(s)
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

// fromJSON will convert a gabs.Container (JSON) to Filter
func (fs FilterService) fromJSON(data *gabs.Container) (*models.Filter, error) {

	if data == nil {
		return nil, nil
	}

	mapped, err := fs.fromJSONToMap(models.NewFilterMap(), data)

	if err != nil {
		return nil, err
	}

	// TODO: process child collections
	return models.NewFilter(mapped), nil
}
