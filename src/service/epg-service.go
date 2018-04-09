package service

import (
	"fmt"
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var epgServiceInstance *EPGService

type EPGService struct {
	ResourceService
}

func GetEPGService(client *Client) *EPGService {
	if epgServiceInstance == nil {
		epgServiceInstance = &EPGService{ResourceService{
			ObjectClass: "@TODO",
		}}
	}
	return epgServiceInstance
}

/* New creates a new EPG  with the appropriate default values */
func (es EPGService) New(name string, description string) *models.EPG {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	e := models.EPG{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  "@TODO",
		ResourceName: resourceName,
	},
		nil,
	}
	//Do any additional construction logic here.
	return &e
}

func (es EPGService) Save(c *models.EPG) error {

	err := es.ResourceService.Save(c)
	if err != nil {
		return err
	}

	return nil

}

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

func (es EPGService) GetAll() ([]*models.EPG, error) {
	var epgs []*models.EPG
	var errors error
	data, err := es.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

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

func (es EPGService) fromJSON(data *gabs.Container) (*models.EPG, error) {
	var errors error
	var valPath, errMsg, name, desc string
	var ok bool

	errMsg = "Could not find value '%s' within child of imdata"
	valPath = ""

	valPath = "@TODO.attributes.name"
	if name, ok = data.Path(valPath).Data().(string); !ok {
		errors = multierror.Append(errors, fmt.Errorf(errMsg, valPath))
	}

	valPath = "@TODO.attributes.descr"
	if desc, ok = data.Path(valPath).Data().(string); !ok {
		errors = multierror.Append(errors, fmt.Errorf(errMsg, valPath))
	}

	if errors != nil {
		return nil, errors
	}

	newEPG := es.New(name, desc)
	return newEPG, nil
}
