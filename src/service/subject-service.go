package service

import (
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

const SJ_RESOURCE_NAME_PREFIX = "subj"
const SJ_OBJECT_CLASS = "vzSubj"

var subjectServiceInstance *SubjectService

type SubjectService struct {
	ResourceService
}

func GetSubjectService(client *Client) *SubjectService {
	if subjectServiceInstance == nil {
		subjectServiceInstance = &SubjectService{ResourceService{
			ObjectClass:        SJ_OBJECT_CLASS,
			ResourceNamePrefix: SJ_RESOURCE_NAME_PREFIX,
			HasParent:          true,
		}}
	}
	return subjectServiceInstance
}

/* New creates a new Subject with the appropriate default values */
func (ss SubjectService) New(name string, description string) *models.Subject {

	s := models.Subject{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  SJ_OBJECT_CLASS,
		ResourceName: ss.getResourceName(name),
	}}

	//Do any additional construction logic here.
	return &s
}

func (ss SubjectService) Save(s *models.Subject) error {

	err := ss.ResourceService.Save(s)
	if err != nil {
		return err
	}

	return nil

}

func (ss SubjectService) Get(domainName string) (*models.Subject, error) {

	data, err := ss.ResourceService.Get(domainName)

	if err != nil {
		return nil, err
	}

	newSubject, err := ss.fromJSON(data)

	if err != nil {
		return nil, err
	}

	return newSubject, nil
}

func (ss SubjectService) GetByName(name string) ([]*models.Subject, error) {

	data, err := ss.ResourceService.GetByName(name)
	if err != nil {
		return nil, err
	}

	return ss.fromDataArray(data)
}

func (ss SubjectService) GetAll() ([]*models.Subject, error) {

	data, err := ss.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	return ss.fromDataArray(data)
}

func (ss SubjectService) fromDataArray(data []*gabs.Container) ([]*models.Subject, error) {
	var epgs []*models.Subject
	var err, errors error

	// For each epg in the payload
	for _, fvSubject := range data {

		newSubject, err := ss.fromJSON(fvSubject)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			epgs = append(epgs, newSubject)

		}
	}

	return epgs, err
}

func (ss SubjectService) fromJSON(data *gabs.Container) (*models.Subject, error) {

	resourceAttributes, err := ss.fromJSONToAttributes(ss.ObjectClass, data)

	if err != nil {
		return nil, err
	}

	return &models.Subject{
		resourceAttributes,
	}, nil
}
