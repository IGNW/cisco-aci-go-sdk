package service

import (
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var epgServiceInstance *EPGService

// EPGService is used to manage EPG resources.
type EPGService struct {
	ResourceService
}

// GetEPGService will construct or return the singleton for the EPGService.
func GetEPGService(client *Client) *EPGService {
	if epgServiceInstance == nil {
		epgServiceInstance = &EPGService{ResourceService{
			ObjectClass:        models.EPG_OBJECT_CLASS,
			ResourceNamePrefix: models.EPG_RESOURCE_NAME_PREFIX,
		}}
	}
	return epgServiceInstance
}

// New creates a new EPG  with the appropriate default values.
func (es EPGService) New(name string, description string) *models.EPG {

	e := models.EPG{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  models.EPG_OBJECT_CLASS,
		ResourceName: es.getResourceName(name),
	},
		false,
		"",
		"",
		"",
	}

	//Do any additional construction logic here.
	return &e
}

// Save a new EPG or update an existing one.
func (es EPGService) Save(c *models.EPG) (string, error) {

	dn, err := es.ResourceService.Save(c)
	if err != nil {
		return "", err
	}

	return dn, nil

}

// Get will retrieve an EPG by it's domain name.
func (es EPGService) Get(domainName string) (*models.EPG, error) {

	data, err := es.ResourceService.Get(domainName)

	if err != nil {
		return nil, err
	}

	newEPG, err := es.fromJSON(data)

	if err != nil {
		return nil, err
	}

	return newEPG, nil
}

// GetById will retrieve an EPG by it's unique identifier.
func (es EPGService) GetById(id string) (*models.EPG, error) {

	data, err := es.ResourceService.GetById(id)

	if err != nil {
		return nil, err
	}

	return es.fromJSON(data)
}

// GetByName will retrieve EPG(s) by common name.
func (es EPGService) GetByName(name string) ([]*models.EPG, error) {

	data, err := es.ResourceService.GetByName(name)
	if err != nil {
		return nil, err
	}

	return es.fromDataArray(data)
}

// GetByName will retrieve all EPG(s).
func (es EPGService) GetAll() ([]*models.EPG, error) {

	data, err := es.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	return es.fromDataArray(data)
}

// fromDataArray will convert an array of gabs.Container (JSON) to EPG(s)
func (es EPGService) fromDataArray(data []*gabs.Container) ([]*models.EPG, error) {
	var epgs []*models.EPG
	var err, errors error

	// For each epg in the payload
	for _, fvEPG := range data {

		newEPG, err := es.fromJSON(fvEPG)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			epgs = append(epgs, newEPG)

		}
	}

	return epgs, err
}

// fromJSON will convert a gabs.Container (JSON) to EPG
func (es EPGService) fromJSON(data *gabs.Container) (*models.EPG, error) {

	if data == nil {
		return nil, nil
	}

	mapped, err := es.fromJSONToMap(models.NewEPGMap(), data)

	if err != nil {
		return nil, err
	}

	// TODO: process child collections
	return models.NewEPG(mapped), nil
}
