package service

import (
	"fmt"
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

var subjectServiceInstance *SubjectService

type SubjectService struct {
	ResourceService
}

func GetSubjectService(client *Client) *SubjectService {
	if subjectServiceInstance == nil {
		subjectServiceInstance = &SubjectService{ResourceService{
			ObjectClass: "@TODO",
		}}
	}
	return subjectServiceInstance
}

/* New creates a new Subject with the appropriate default values */
func (ss SubjectService) New(name string, description string) *models.Subject {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	s := models.Subject{models.ResourceAttributes{
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

func (ss SubjectService) GetAll() ([]*models.Subject, error) {
	var epgs []*models.Subject
	var errors error
	data, err := ss.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	fvSubjects, err := data.S("imdata").Children()
	if err != nil {
		return nil, err
	}

	// For each epg in the payload
	for _, fvSubject := range fvSubjects {

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

	newSubject := ss.New(name, desc)
	return newSubject, nil
}
