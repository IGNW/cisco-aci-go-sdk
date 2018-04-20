package service

import (
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var contractServiceInstance *ContractService

type ContractService struct {
	ResourceService
}

func GetContractService(client *Client) *ContractService {
	if contractServiceInstance == nil {
		contractServiceInstance = &ContractService{ResourceService{
			ObjectClass:        models.CONTRACT_OBJECT_CLASS,
			ResourceNamePrefix: models.CONTRACT_RESOURCE_PREFIX,
			HasParent:          true,
		}}
	}
	return contractServiceInstance
}

/* New creates a new Contract with the appropriate default values */
func (cs ContractService) New(name string, description string) *models.Contract {

	b := models.Contract{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  models.CONTRACT_OBJECT_CLASS,
		ResourceName: cs.getResourceName(name),
	},
		"",
		"",
		nil,
		nil,
	}
	//Do any additional construction logic here.
	return &b
}

func (cs ContractService) Save(c *models.Contract) error {

	err := cs.ResourceService.Save(c)
	if err != nil {
		return err
	}

	return nil

}

func (cs ContractService) Get(domainName string) (*models.Contract, error) {

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

func (cs ContractService) GetByName(name string) ([]*models.Contract, error) {

	data, err := cs.ResourceService.GetByName(name)
	if err != nil {
		return nil, err
	}

	return cs.fromDataArray(data)
}

func (cs ContractService) GetAll() ([]*models.Contract, error) {

	data, err := cs.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	return cs.fromDataArray(data)
}

func (cs ContractService) fromDataArray(data []*gabs.Container) ([]*models.Contract, error) {
	var contracts []*models.Contract
	var err, errors error

	// For each contract in the payload
	for _, fvContract := range data {

		newContract, err := cs.fromJSON(fvContract)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			contracts = append(contracts, newContract)

		}
	}

	return contracts, err
}

func (cs ContractService) fromJSON(data *gabs.Container) (*models.Contract, error) {
	mapped, err := cs.fromJSONToMap(models.NewContractMap(), data)

	if err != nil {
		return nil, err
	}

	// TODO: process child collections
	return models.NewContract(mapped), nil
}
