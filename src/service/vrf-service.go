package service

import (
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var vrfServiceInstance *VRFService

type VRFService struct {
	ResourceService
}

func GetVRFService(client *Client) *VRFService {
	if vrfServiceInstance == nil {
		vrfServiceInstance = &VRFService{ResourceService{
			ObjectClass:        models.VRF_OBJECT_CLASS,
			ResourceNamePrefix: models.VRF_RESOURCE_PREFIX,
			HasParent:          true,
		}}
	}
	return vrfServiceInstance
}

/* New creates a new VRF with the appropriate default values */
func (vs VRFService) New(name string, description string) *models.VRF {

	t := models.VRF{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  models.VRF_OBJECT_CLASS,
		ResourceName: vs.getResourceName(name),
	},
		"",
		"",
		nil,
	}
	//Do any additional construction logic here.
	return &t
}

func (vs VRFService) Save(v *models.VRF) error {

	err := vs.ResourceService.Save(v)
	if err != nil {
		return err
	}

	return nil

}

func (vs VRFService) Get(domainName string) (*models.VRF, error) {

	data, err := vs.ResourceService.Get(domainName)

	if err != nil {
		return nil, err
	}

	newVRF, err := vs.fromJSON(data)

	if err != nil {
		return nil, err
	}

	return newVRF, nil
}

func (vs VRFService) GetByName(name string) ([]*models.VRF, error) {

	data, err := vs.ResourceService.GetByName(name)
	if err != nil {
		return nil, err
	}

	return vs.fromDataArray(data)
}

func (vs VRFService) GetAll() ([]*models.VRF, error) {

	data, err := vs.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	return vs.fromDataArray(data)
}

func (vs VRFService) fromDataArray(data []*gabs.Container) ([]*models.VRF, error) {
	var vrfs []*models.VRF
	var err, errors error

	// For each vrf in the payload
	for _, fvVRF := range data {

		newVRF, err := vs.fromJSON(fvVRF)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			vrfs = append(vrfs, newVRF)

		}
	}

	return vrfs, err
}

func (vs VRFService) fromJSON(data *gabs.Container) (*models.VRF, error) {
	mapped, err := vs.fromJSONToMap(models.NewVRFMap(), data)

	if err != nil {
		return nil, err
	}

	// TODO: process child collections

	return models.NewVRF(mapped), nil
}
