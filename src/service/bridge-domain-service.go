package service

import (
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var bridgeDomainServiceInstance *BridgeDomainService

const BD_RESOURCE_NAME_PREFIX = "BD"
const BD_OBJECT_CLASS = "fvBD"

type BridgeDomainService struct {
	ResourceService
}

func GetBridgeDomainService(client *Client) *BridgeDomainService {
	if bridgeDomainServiceInstance == nil {
		bridgeDomainServiceInstance = &BridgeDomainService{ResourceService{
			ObjectClass:        BD_OBJECT_CLASS,
			ResourceNamePrefix: BD_RESOURCE_NAME_PREFIX,
			HasParent:          true,
		}}
	}
	return bridgeDomainServiceInstance
}

/* New creates a new BridgeDomain with the appropriate default values */
func (bds BridgeDomainService) New(name string, description string) *models.BridgeDomain {

	b := models.BridgeDomain{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  BD_OBJECT_CLASS,
		ResourceName: bds.getResourceName(name),
	},
		nil,
		nil,
	}
	//Do any additional construction logic here.
	return &b
}

func (bds BridgeDomainService) Save(t *models.BridgeDomain) error {

	err := bds.ResourceService.Save(t)
	if err != nil {
		return err
	}

	return nil

}

func (bds BridgeDomainService) Get(domainName string) (*models.BridgeDomain, error) {

	data, err := bds.ResourceService.Get(domainName)

	if err != nil {
		return nil, err
	}

	newBridgeDomain, err := bds.fromJSON(data)

	if err != nil {
		return nil, err
	}

	return newBridgeDomain, nil
}

func (bds BridgeDomainService) GetById(id string) (*models.BridgeDomain, error) {
	var bridgeDomains []*models.BridgeDomain

	data, err := bds.ResourceService.GetById(id)

	if err != nil {
		return nil, err
	}

	bridgeDomains, err = bds.fromDataArray(data)

	if err != nil {
		return nil, err
	}

	if len(bridgeDomains) > 0 {
		return bridgeDomains[0], nil
	} else {
		return nil, nil
	}
}

func (bds BridgeDomainService) GetByName(name string) ([]*models.BridgeDomain, error) {

	data, err := bds.ResourceService.GetByName(name)

	if err != nil {
		return nil, err
	}

	return bds.fromDataArray(data)
}

func (bds BridgeDomainService) GetAll() ([]*models.BridgeDomain, error) {

	data, err := bds.ResourceService.GetAll()

	if err != nil {
		return nil, err
	}

	return bds.fromDataArray(data)
}

func (bds BridgeDomainService) fromDataArray(data *gabs.Container) ([]*models.BridgeDomain, error) {
	var bridgeDomains []*models.BridgeDomain
	var errors error

	fvBridgeDomains, err := data.S("imdata").Children()
	if err != nil {
		return nil, err
	}

	// For each bridgeDomain in the payload
	for _, fvBridgeDomain := range fvBridgeDomains {

		newBridgeDomain, err := bds.fromJSON(fvBridgeDomain)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			bridgeDomains = append(bridgeDomains, newBridgeDomain)

		}
	}

	return bridgeDomains, err
}

func (bds BridgeDomainService) fromJSON(data *gabs.Container) (*models.BridgeDomain, error) {
	/*
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

		newBridgeDomain := bds.New(name, desc)
		return newBridgeDomain, nil
	*/

	resourceAttributes, err := bds.fromJSONToAttributes(bds.ObjectClass, data)

	if err != nil {
		return nil, err
	}

	// TODO: process child collections

	return &models.BridgeDomain{
		resourceAttributes,
		nil,
		nil,
	}, nil
}
