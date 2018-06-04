package service

import (
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var bridgeDomainServiceInstance *BridgeDomainService

// BridgeDomainService is used to manage BridgeDomain resources.
type BridgeDomainService struct {
	ResourceService
}

// GetBridgeDomainService will construct or return the singleton for the BridgeDomainService.
func GetBridgeDomainService(client *Client) *BridgeDomainService {
	if bridgeDomainServiceInstance == nil {
		bridgeDomainServiceInstance = &BridgeDomainService{ResourceService{
			ObjectClass:        models.BD_OBJECT_CLASS,
			ResourceNamePrefix: models.BD_RESOURCE_PREFIX,
		}}
	}
	return bridgeDomainServiceInstance
}

// New creates a new BridgeDomain with the appropriate default values.
func (bds BridgeDomainService) New(name string, description string) *models.BridgeDomain {

	b := models.BridgeDomain{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  models.BD_OBJECT_CLASS,
		ResourceName: bds.getResourceName(name),
	},
		"regular",
		false,
		false,
		"",
		false,
		false,
		true,
		true,
		"::",
		"00:22:BD:F8:19:FF",
		"bd-flood",
		false,
		true,
		"proxy",
		"flood",
		"not-applicable",
		nil,
		nil,
	}

	//Do any additional construction logic here.
	return &b
}

// Save a new BridgeDomain or update an existing one.
func (bds BridgeDomainService) Save(t *models.BridgeDomain) (string, error) {

	dn, err := bds.ResourceService.Save(t)
	if err != nil {
		return "", err
	}

	return dn, nil

}

// Get will retrieve an BridgeDomain by it's domain name.
func (bds BridgeDomainService) Get(domainName string) (*models.BridgeDomain, error) {

	data, err := bds.ResourceService.Get(domainName)

	if err != nil {
		return nil, err
	}

	return bds.fromJSON(data)

}

// GetById will retrieve an BridgeDomain by it's unique identifier.
func (bds BridgeDomainService) GetById(id string) (*models.BridgeDomain, error) {

	data, err := bds.ResourceService.GetById(id)

	if err != nil {
		return nil, err
	}

	return bds.fromJSON(data)
}

// GetByName will retrieve BridgeDomain(s) by common name.
func (bds BridgeDomainService) GetByName(name string) ([]*models.BridgeDomain, error) {

	data, err := bds.ResourceService.GetByName(name)

	if err != nil {
		return nil, err
	}

	return bds.fromDataArray(data)
}

// GetByName will retrieve all BridgeDomain(s).
func (bds BridgeDomainService) GetAll() ([]*models.BridgeDomain, error) {

	data, err := bds.ResourceService.GetAll()

	if err != nil {
		return nil, err
	}

	return bds.fromDataArray(data)
}

// fromDataArray will convert an array of gabs.Container (JSON) to BridgeDomain(s)
func (bds BridgeDomainService) fromDataArray(data []*gabs.Container) ([]*models.BridgeDomain, error) {
	var bridgeDomains []*models.BridgeDomain
	var err, errors error

	// For each bridgeDomain in the payload
	for _, fvBridgeDomain := range data {

		newBridgeDomain, err := bds.fromJSON(fvBridgeDomain)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			bridgeDomains = append(bridgeDomains, newBridgeDomain)

		}
	}

	return bridgeDomains, err
}

// fromJSON will convert a gabs.Container (JSON) to BridgeDomain
func (bds BridgeDomainService) fromJSON(data *gabs.Container) (*models.BridgeDomain, error) {

	modelMap, err := bds.fromJSONToMap(models.NewBridgeDomainMap(), data)

	if err != nil {
		return nil, err
	}

	return models.NewBridgeDomain(modelMap), nil

}
