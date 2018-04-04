package service

import (
	"fmt"

	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
)

var bridgeDomainServiceInstance *BridgeDomainService

type BridgeDomainService struct {
	ResourceService
}

func GetBridgeDomainService(client *Client) *BridgeDomainService {
	if bridgeDomainServiceInstance == nil {
		bridgeDomainServiceInstance = &BridgeDomainService{ResourceService{
			ObjectClass: "@TODO",
		}}
	}
	return bridgeDomainServiceInstance
}

/* New creates a new BridgeDomain with the appropriate default values */
func (bds BridgeDomainService) New(name string, description string) *BridgeDomain {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	b := BridgeDomain{ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  "@TODO",
		ResourceName: resourceName,
	},
		nil,
		nil,
	}
	//Do any additional construction logic here.
	return &b
}

func (bds BridgeDomainService) Save(t *BridgeDomain) error {

	err := bds.ResourceService.Save(t)
	if err != nil {
		return err
	}

	return nil

}

func (bds BridgeDomainService) Get(domainName string) (*BridgeDomain, error) {

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

func (bds BridgeDomainService) GetAll() ([]*BridgeDomain, error) {
	var bridgeDomains []*BridgeDomain
	var errors error
	data, err := bds.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

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

func (bds BridgeDomainService) fromJSON(data *gabs.Container) (*BridgeDomain, error) {
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
}
