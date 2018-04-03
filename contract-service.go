package cage

import (
	"fmt"

	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
)

var contractServiceInstance *ContractService

type ContractService struct {
	ResourceService
}

func GetContractService(client *Client) *ContractService {
	if contractServiceInstance == nil {
		contractServiceInstance = &ContractService{ResourceService{
			ObjectClass: "@TODO",
		}}
	}
	return contractServiceInstance
}

/* New creates a new Contract with the appropriate default values */
func (cs ContractService) New(name string, description string) *Contract {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	b := Contract{ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  "@TODO",
		ResourceName: resourceName,
	},
		"",
		nil,
		nil,
	}
	//Do any additional construction logic here.
	return &b
}

func (cs ContractService) Save(c *Contract) error {

	err := cs.ResourceService.Save(c)
	if err != nil {
		return err
	}

	return nil

}

func (cs ContractService) Get(domainName string) (*Contract, error) {

	data, err := cs.ResourceService.Get(domainName)

	if err != nil {
		return nil, err
	}

	newContract, err := cs.fromJSON(data)

	if err != nil {
		return nil, err
	}

	return newContract, nil
}

func (cs ContractService) GetAll() ([]*Contract, error) {
	var contracts []*Contract
	var errors error
	data, err := cs.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	fvContracts, err := data.S("imdata").Children()
	if err != nil {
		return nil, err
	}

	// For each contract in the payload
	for _, fvContract := range fvContracts {

		newContract, err := cs.fromJSON(fvContract)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			contracts = append(contracts, newContract)

		}
	}

	return contracts, err
}

func (cs ContractService) fromJSON(data *gabs.Container) (*Contract, error) {
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

	newContract := cs.New(name, desc)
	return newContract, nil
}
