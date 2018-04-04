package service

import (
	"fmt"

	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
)

var subnetServiceInstance *SubnetService

type SubnetService struct {
	ResourceService
}

func GetSubnetService(client *Client) *SubnetService {
	if subnetServiceInstance == nil {
		subnetServiceInstance = &SubnetService{ResourceService{
			ObjectClass: "@TODO",
		}}
	}
	return subnetServiceInstance
}

/* New creates a new Subnet with the appropriate default values */
func (ss SubnetService) New(name string, description string) *Subnet {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	s := Subnet{ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  "@TODO",
		ResourceName: resourceName,
	},
		nil,
	}
	//Do any additional construction logic here.
	return &s
}

func (ss SubnetService) Save(s *Subnet) error {

	err := ss.ResourceService.Save(s)
	if err != nil {
		return err
	}

	return nil

}

func (ss SubnetService) Get(domainName string) (*Subnet, error) {

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

func (ss SubnetService) GetAll() ([]*Subnet, error) {
	var epgs []*Subnet
	var errors error
	data, err := ss.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	fvSubnets, err := data.S("imdata").Children()
	if err != nil {
		return nil, err
	}

	// For each epg in the payload
	for _, fvSubnet := range fvSubnets {

		newSubnet, err := ss.fromJSON(fvSubnet)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			epgs = append(epgs, newSubnet)

		}
	}

	return epgs, err
}

func (ss SubnetService) fromJSON(data *gabs.Container) (*Subnet, error) {
	var errors error
	var valPath, errMsg, name, desc string
	var ok bool

	errMsg = "Could not find value '%s' within child of imdata"
	valPath = ""

	valPath = "@TODO.attributss.name"
	if name, ok = data.Path(valPath).Data().(string); !ok {
		errors = multierror.Append(errors, fmt.Errorf(errMsg, valPath))
	}

	valPath = "@TODO.attributss.descr"
	if desc, ok = data.Path(valPath).Data().(string); !ok {
		errors = multierror.Append(errors, fmt.Errorf(errMsg, valPath))
	}

	if errors != nil {
		return nil, errors
	}

	newSubnet := ss.New(name, desc)
	return newSubnet, nil
}
