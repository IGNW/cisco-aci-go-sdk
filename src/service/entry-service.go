package service

import (
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var entryServiceInstance *EntryService

type EntryService struct {
	ResourceService
}

func GetEntryService(client *Client) *EntryService {
	if entryServiceInstance == nil {
		entryServiceInstance = &EntryService{ResourceService{
			ObjectClass:        models.ENTRY_OBJECT_CLASS,
			ResourceNamePrefix: models.ENTRY_RESOURCE_PREFIX,
		}}
	}
	return entryServiceInstance
}

// New creates a new Entry with the appropriate default values.
func (es EntryService) New(name string, description string) *models.Entry {

	e := models.Entry{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  models.ENTRY_OBJECT_CLASS,
		ResourceName: es.getResourceName(name),
	},
		"unspecified",
		"unspecified",
		false,
		"unspecified",
		"unspecified",
		"unspecified",
		"unspecified",
		false,
		"unspecified",
		&models.ToFrom{To: "", From: ""},
		&models.ToFrom{To: "", From: ""},
	}
	//Do any additional construction logic here.
	return &e
}

func (es EntryService) Save(e *models.Entry) error {

	err := es.ResourceService.Save(e)
	if err != nil {
		return err
	}

	return nil

}

func (es EntryService) Get(domainName string) (*models.Entry, error) {

	data, err := es.ResourceService.Get(domainName)

	if err != nil {
		return nil, err
	}

	newFilter, err := es.fromJSON(data)

	if err != nil {
		return nil, err
	}

	return newFilter, nil
}

func (es EntryService) GetByName(name string) ([]*models.Entry, error) {

	data, err := es.ResourceService.GetByName(name)
	if err != nil {
		return nil, err
	}

	return es.fromDataArray(data)
}

func (es EntryService) GetAll() ([]*models.Entry, error) {

	data, err := es.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	return es.fromDataArray(data)
}

func (es EntryService) fromDataArray(data []*gabs.Container) ([]*models.Entry, error) {
	var entries []*models.Entry
	var err, errors error

	// For each epg in the payload
	for _, item := range data {

		newEntry, err := es.fromJSON(item)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			entries = append(entries, newEntry)

		}
	}

	return entries, err
}

func (es EntryService) fromJSON(data *gabs.Container) (*models.Entry, error) {
	mapped, err := es.fromJSONToMap(models.NewEntryMap(), data)

	if err != nil {
		return nil, err
	}

	// TODO: process child collections
	return models.NewEntry(mapped), nil
}
