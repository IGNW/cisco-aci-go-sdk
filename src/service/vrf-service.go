package service

import (
	"fmt"
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

// TODO: validate these settings are correct
const V_RESOURCE_NAME_PREFIX = "V"
const V_OBJECT_CLASS = "fvVRF"

var vrfServiceInstance *VRFService

type VRFService struct {
	ResourceService
}

func GetVRFService(client *Client) *VRFService {
	if vrfServiceInstance == nil {
		vrfServiceInstance = &VRFService{ResourceService{
			ObjectClass:        V_OBJECT_CLASS,
			ResourceNamePrefix: V_RESOURCE_NAME_PREFIX,
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
		ObjectClass:  V_OBJECT_CLASS,
		ResourceName: vs.getResourceName(name),
	},
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

func (vs VRFService) GetAll() ([]*models.VRF, error) {
	var vrfs []*models.VRF
	var errors error
	data, err := vs.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

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
	var errors error
	var valPath, errMsg, name, desc string
	var ok bool

	errMsg = "Could not find value '%s' within child of imdata"
	valPath = ""

	valPath = "fvVRF.attributes.name"
	if name, ok = data.Path(valPath).Data().(string); !ok {
		errors = multierror.Append(errors, fmt.Errorf(errMsg, valPath))
	}

	valPath = "fvVRF.attributes.descr"
	if desc, ok = data.Path(valPath).Data().(string); !ok {
		errors = multierror.Append(errors, fmt.Errorf(errMsg, valPath))
	}

	if errors != nil {
		return nil, errors
	}

	newVRF := vs.New(name, desc)
	return newVRF, nil
}
