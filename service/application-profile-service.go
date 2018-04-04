package service

import (
	"fmt"

	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
)

var appProfileServiceInstance *AppProfileService

type AppProfileService struct {
	ResourceService
}

func GetAppProfileService(client *Client) *AppProfileService {
	if appProfileServiceInstance == nil {
		appProfileServiceInstance = &AppProfileService{ResourceService{
			ObjectClass: "@TODO",
		}}
	}
	return appProfileServiceInstance
}

/* New creates a new AppProfile with the appropriate default values */
func (aps AppProfileService) New(name string, description string) *AppProfile {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	a := AppProfile{ResourceAttributes{
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
	return &a
}

func (aps AppProfileService) Save(t *AppProfile) error {

	err := aps.ResourceService.Save(t)
	if err != nil {
		return err
	}

	return nil

}

func (aps AppProfileService) Get(domainName string) (*AppProfile, error) {

	data, err := aps.ResourceService.Get(domainName)

	if err != nil {
		return nil, err
	}

	newAppProfile, err := aps.fromJSON(data)

	if err != nil {
		return nil, err
	}

	return newAppProfile, nil
}

func (aps AppProfileService) GetAll() ([]*AppProfile, error) {
	var appProfiles []*AppProfile
	var errors error
	data, err := aps.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	fvAppProfiles, err := data.S("imdata").Children()
	if err != nil {
		return nil, err
	}

	// For each appProfile in the payload
	for _, fvAppProfile := range fvAppProfiles {

		newAppProfile, err := aps.fromJSON(fvAppProfile)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			appProfiles = append(appProfiles, newAppProfile)

		}
	}

	return appProfiles, err
}

func (aps AppProfileService) fromJSON(data *gabs.Container) (*AppProfile, error) {
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

	newAppProfile := aps.New(name, desc)
	return newAppProfile, nil
}
