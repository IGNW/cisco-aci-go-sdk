package service

import (
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var appProfileServiceInstance *AppProfileService

// AppProfileService is used to manage AppProfile resources.
type AppProfileService struct {
	ResourceService
}

// GetAppProfileService will construct or return the singleton for the AppProfileService.
func GetAppProfileService(client *Client) *AppProfileService {
	if appProfileServiceInstance == nil {
		appProfileServiceInstance = &AppProfileService{ResourceService{
			ObjectClass:        models.AP_OBJECT_CLASS,
			ResourceNamePrefix: models.AP_RESOURCE_PREFIX,
		}}
	}
	return appProfileServiceInstance
}

// New creates a new AppProfile with the appropriate default values.
func (aps AppProfileService) New(name string, description string) *models.AppProfile {

	a := models.AppProfile{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  models.AP_OBJECT_CLASS,
		ResourceName: aps.getResourceName(name),
	},
		nil,
	}
	//Do any additional construction logic here.
	return &a
}

// Save a new AppProfile or update an existing one.
func (aps AppProfileService) Save(t *models.AppProfile) (string, error) {

	dn, err := aps.ResourceService.Save(t)
	if err != nil {
		return "", err
	}

	return dn, nil

}

// Get will retrieve an AppProfile by it's domain name.
func (aps AppProfileService) Get(domainName string) (*models.AppProfile, error) {

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

// GetById will retrieve an AppProfile by it's unique identifier.
func (aps AppProfileService) GetById(id string) (*models.AppProfile, error) {

	data, err := aps.ResourceService.GetById(id)

	if err != nil {
		return nil, err
	}

	return aps.fromJSON(data)
}

// GetByName will retrieve AppProfile(s) by common name.
func (aps AppProfileService) GetByName(name string) ([]*models.AppProfile, error) {

	data, err := aps.ResourceService.GetByName(name)

	if err != nil {
		return nil, err
	}

	return aps.fromDataArray(data)
}

// GetByName will retrieve all AppProfile(s).
func (aps AppProfileService) GetAll() ([]*models.AppProfile, error) {

	data, err := aps.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	return aps.fromDataArray(data)
}

// fromDataArray will convert an array of gabs.Container (JSON) to AppProfile(s)
func (aps AppProfileService) fromDataArray(data []*gabs.Container) ([]*models.AppProfile, error) {
	var appProfiles []*models.AppProfile
	var err, errors error

	// For each appProfile in the payload
	for _, fvAppProfile := range data {

		newAppProfile, err := aps.fromJSON(fvAppProfile)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			appProfiles = append(appProfiles, newAppProfile)

		}
	}

	return appProfiles, err
}

// fromJSON will convert a gabs.Container (JSON) to AppProfile
func (aps AppProfileService) fromJSON(data *gabs.Container) (*models.AppProfile, error) {

	if data == nil {
		return nil, nil
	}

	mapped, err := aps.fromJSONToMap(models.NewAppProfileMap(), data)

	if err != nil {
		return nil, err
	}

	// TODO: process child collections
	return models.NewAppProfile(mapped), nil
}
