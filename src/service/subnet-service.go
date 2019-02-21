package service

import (
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var subnetServiceInstance *SubnetService

type SubnetService struct {
	ResourceService
}

func GetSubnetService(client *Client) *SubnetService {
	if subnetServiceInstance == nil {
		subnetServiceInstance = &SubnetService{ResourceService{
			ObjectClass:        models.SUBNET_OBJECT_CLASS,
			ResourceNamePrefix: models.SUBNET_RESOURCE_NAME_PREFIX,
		}}
	}
	return subnetServiceInstance
}

/* New creates a new Subnet with the appropriate default values */
func (ss SubnetService) New(name string, description string) *models.Subnet {

	s := models.Subnet{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  models.SUBNET_OBJECT_CLASS,
		ResourceName: ss.getResourceName(name),
	},
		"nd",
		"",
		false,
		[]string{"private"},
		false,
	}

	//Do any additional construction logic here.
	return &s
}

func (ss SubnetService) Save(s *models.Subnet) (string, error) {

	// HACK: fix resource name this one has a difference scheme than all the others
	s.ResourceName = s.GetResourceName()

	dn, err := ss.ResourceService.Save(s)

	if err != nil {
		return "", err
	}

	return dn, nil

}

func (ss SubnetService) Get(domainName string) (*models.Subnet, error) {

	data, err := ss.ResourceService.Get(domainName)

	if err != nil {
		return nil, err
	}

	newSubnet, err := ss.fromJSON(data)

	if err != nil {
		return nil, err
	}

	return newSubnet, nil
}

func (ss SubnetService) GetByName(name string) ([]*models.Subnet, error) {

	data, err := ss.ResourceService.GetByName(name)
	if err != nil {
		return nil, err
	}

	return ss.fromDataArray(data)
}

func (ss SubnetService) GetAll() ([]*models.Subnet, error) {

	data, err := ss.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	return ss.fromDataArray(data)
}

func (ss SubnetService) fromDataArray(data []*gabs.Container) ([]*models.Subnet, error) {
	var epgs []*models.Subnet
	var err, errors error

	// For each epg in the payload
	for _, fvSubnet := range data {

		newSubnet, err := ss.fromJSON(fvSubnet)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			epgs = append(epgs, newSubnet)

		}
	}

	return epgs, err
}

func (ss SubnetService) fromJSON(data *gabs.Container) (*models.Subnet, error) {

	if data == nil {
		return nil, nil
	}

	mapped, err := ss.fromJSONToMap(models.NewSubnetMap(), data)

	if err != nil {
		return nil, err
	}

	return models.NewSubnet(mapped), nil
}
