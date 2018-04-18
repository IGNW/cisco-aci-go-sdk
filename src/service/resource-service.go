package service

import (
	"fmt"
	"github.com/Jeffail/gabs"
	log "github.com/golang/glog"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"net/http"
	"strings"
)

// Services represents a collection of service objects within the client
// will acess.client().AppProfiles ...
type Services struct {
	AppProfiles   *AppProfileService
	BridgeDomains *BridgeDomainService
	Contracts     *ContractService
	VRFs          *VRFService
	EPGs          *EPGService
	Filters       *FilterService
	Subjects      *SubjectService
	Subnets       *SubnetService
	Tenants       *TenantService
}

type ResourceGenerator func(string, string, string) models.ResourceInterface
type ResourceDecoder func(*gabs.Container) (models.ResourceInterface, error)

type ResourceService struct {
	ObjectClass        string
	ResourceNamePrefix string
	HasParent          bool
}

func (s ResourceService) client() *Client {
	return GetClient()
}

func (s ResourceService) Save(r models.ResourceInterface) (err error) {
	var path string
	var parent models.ResourceInterface

	// perform base validation
	err = s.validate(r)
	if err != nil {
		return fmt.Errorf("\nGot Error While Validating, Auth'd Request failed w/ %v", err)
	}

	json, err := s.toJSON(r)

	json.Set("created, modified", s.ObjectClass, "attributes", "status")

	method := "POST"

	// TODO: refactor to getResourcePath()
	path = s.getResourcePath(r, "")

	parent = r.GetParent()

	if parent != nil {
		path = fmt.Sprintf("/api/node/mo/uni/%s/%s.json", parent.GetResourceName(), r.GetResourceName())
	} else {
		path = fmt.Sprintf("/api/node/mo/uni/%s.json", r.GetResourceName())
	}

	// END Refactor

	req, err := s.client().newAuthdRequest(method, path, json)
	if err != nil {
		return fmt.Errorf("\nGot Error While Saving, Auth'd Request failed w/ %v", err)
	}

	data, response, err := s.client().do(req)

	log.Infof("RESP: %#v\n\n", response)
	log.Infof("DATA: %#v\n\n", data)

	if err != nil {
		return err
	}

	if err = s.getResponseError(data); err != nil {
		return err
	}

	return nil
}

func (s ResourceService) Get(domainName string) (*gabs.Container, error) {

	// TODO: refactor to use domain name field

	path := fmt.Sprintf("api/mo/%s.json", domainName)

	req, err := s.client().newAuthdRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	data, response, err := s.client().do(req)

	err = s.combineErrors(data, response, err)

	if err != nil {
		return nil, err
	}

	return s.getChild(data)

}

func (s ResourceService) GetById(id string) (*gabs.Container, error) {

	path := fmt.Sprintf("/api/node/class/%s.json?query-target-filter=eq(%s.id,\"%s\")", s.ObjectClass, s.ObjectClass, id)

	req, err := s.client().newAuthdRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	data, response, err := s.client().do(req)

	err = s.combineErrors(data, response, err)

	if err != nil {
		return nil, err
	}

	return s.getChild(data)
}

func (s ResourceService) GetByName(name string) ([]*gabs.Container, error) {
	path := fmt.Sprintf("/api/node/class/%s.json?query-target-filter=eq(%s.name,\"%s\")", s.ObjectClass, s.ObjectClass, name)

	req, err := s.client().newAuthdRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	data, response, err := s.client().do(req)

	err = s.combineErrors(data, response, err)

	if err != nil {
		return nil, err
	}

	return s.getChildren(data)
}

func (s ResourceService) GetAll() ([]*gabs.Container, error) {
	path := fmt.Sprintf("/api/node/class/%s.json", s.ObjectClass)
	req, err := s.client().newAuthdRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	data, response, err := s.client().do(req)

	err = s.combineErrors(data, response, err)

	if err != nil {
		return nil, err
	}

	return s.getChildren(data)
}

func (s ResourceService) Delete(domainName string) error {
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
			}
		}
	}`, s.ObjectClass))

	data, err := gabs.ParseJSON(containerJSON)

	if err != nil {
		return fmt.Errorf("Error parsing JSON.\n%#v", err)
	}

	_, err = data.Set("deleted", s.ObjectClass, "attributes", "status")
	if err != nil {
		return fmt.Errorf("Error setting deleted flag on ACI model.\n%#v", err)
	}

	_, err = data.Set(domainName, s.ObjectClass, "attributes", "dn")
	if err != nil {
		return fmt.Errorf("Error setting attributes on ACI model.\n%#v", err)
	}

	log.Infof("Payload: %s", data.String())

	path := fmt.Sprintf("/api/node/mo/%s.json", domainName)

	req, err := s.client().newAuthdRequest("POST", path, data)

	if err != nil {
		return fmt.Errorf("Error with delete request to ACI: POST %s\n%s", path, err)
	}

	data, _, err = s.client().do(req)

	if err != nil {
		return err
	}

	if err = s.getResponseError(data); err != nil {
		return err
	}

	return nil
}

func (s ResourceService) getChild(data *gabs.Container) (*gabs.Container, error) {
	items, err := data.S("imdata").Children()

	if err != nil {
		return nil, err
	}

	return items[0], nil
}

func (s ResourceService) getChildren(data *gabs.Container) ([]*gabs.Container, error) {
	return data.S("imdata").Children()
}

// TODO: refactor the create json function from ResourceAttributes to ResourceService
func (s ResourceService) toJSON(model models.ResourceInterface) (*gabs.Container, error) {
	var data *gabs.Container
	var err error

	data, err = s.CreateEmptyJSONContainer()

	if err != nil {
		return nil, err
	}

	for key, value := range model.ToMap() {
		data.Set(value, s.ObjectClass, "attributes", key)
		//TODO: add error checking?
	}

	data.Array(s.ObjectClass, "children")

	return data, nil
}

// Create an empty JSON container compatible with ACI object model.
func (s ResourceService) CreateEmptyJSONContainer() (*gabs.Container, error) {
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
			}
		}
	}`, s.ObjectClass))

	return gabs.ParseJSON(containerJSON)
}

// TODO: Deprecate
// Convert JSON container to ResourceAttributes.
func (s ResourceService) fromJSONToAttributes(objectClass string, data *gabs.Container) (models.ResourceAttributes, error) {
	var errors error
	var path, value, errMsg string
	var ok bool

	// TODO: tags
	var attributes = map[string]string{"dn": "", "name": "", "descr": "", "status": ""}

	errMsg = "Could not find value '%s' within child of imdata"

	for key, _ := range attributes {
		path = objectClass + ".attributes." + key
		if value, ok = data.Path(path).Data().(string); !ok {
			errors = multierror.Append(errors, fmt.Errorf(errMsg, path))
		}
		attributes[key] = value
	}

	if errors != nil {
		return models.ResourceAttributes{}, errors
	}

	paths := strings.Split(attributes["dn"], "/")

	return models.ResourceAttributes{
		DomainName:   attributes["dn"],
		Name:         attributes["name"],
		Description:  attributes["descr"],
		Status:       attributes["status"],
		ResourceName: paths[len(paths)-1],
		ObjectClass:  objectClass,
	}, nil

}

// fromJSONToMap converts the Gabs container into a string map based on the supplied model template
func (s ResourceService) fromJSONToMap(template map[string]string, data *gabs.Container) (map[string]string, error) {
	var errors error
	var path, value, errMsg string
	var ok bool

	errMsg = "Could not find value '%s' within child of imdata"

	for key, _ := range template {
		path = s.ObjectClass + ".attributes." + key
		if value, ok = data.Path(path).Data().(string); !ok {
			errors = multierror.Append(errors, fmt.Errorf(errMsg, path))
		}
		template[key] = value
	}

	if errors != nil {
		return nil, errors
	}

	paths := strings.Split(template["dn"], "/")

	template["rn"] = paths[len(paths)-1]
	template["objectClass"] = s.ObjectClass

	return template, nil
}

func (s ResourceService) getResourceName(name string) string {
	resourceName := fmt.Sprintf("%s-%s", s.ResourceNamePrefix, name)
	return resourceName
}

func (s ResourceService) getResourcePath(model models.ResourceInterface, path string) string {
	const basePath = "/api/node/mo/uni/"
	var parent models.ResourceInterface

	if path == "" {
		path = basePath
	}

	parent = model.GetParent()

	if parent == nil {
		path += model.GetResourceName() + ".json"
		return path
	} else {
		path += model.GetResourceName() + "/"
		return s.getResourcePath(parent, path)
	}
}

func (s ResourceService) validate(model models.ResourceInterface) error {
	var err error

	if s.HasParent && model.GetParent() == nil {
		err = fmt.Errorf("Models of type '%s' require a parent to be set", s.ObjectClass)
	}

	return err
}

func (s ResourceService) getGabsValue(data *gabs.Container, valuePath string) string {
	// Not sure if this is Cisco or Gabs, but wow.
	// @TODO find a better way to extract values from gabs containers
	return data.Path(valuePath).Data().([]interface{})[0].(string)
}

//@TODO rename this, response should always refer to an http.Response for clarity, this is a response body digested as a gabs Container
func (s ResourceService) getResponseError(responseData *gabs.Container) error {
	valpath := "imdata.error.attributes.text"

	if exists := responseData.ExistsP("imdata.error.attributes.text"); exists {
		err := s.getGabsValue(responseData, valpath)

		if err != "" {
			return fmt.Errorf(err)
		}
	}

	return nil
}

func (s ResourceService) combineErrors(data *gabs.Container, response *http.Response, err error) error {

	var errors *multierror.Error
	var newErr error

	// err will be set if there was an error making the request
	if err != nil {
		newErr = fmt.Errorf("Got Error Making Call: %s", err)
		errors = multierror.Append(errors, err)
	}

	// Then we check if the request worked but the API returned an error
	if response.StatusCode >= 400 {
		newErr = fmt.Errorf("Got Request Error:\n    StatusCode: %v\n    Status: %s", response.StatusCode, response.Status)
		errors = multierror.Append(errors, newErr)
	}

	if newErr = s.getResponseError(data); newErr != nil {
		errors = multierror.Append(errors, fmt.Errorf("Response Body contained an error message:\n   %s", newErr))
	}

	/**TODO change multi error reporting to be semantically correct,
	right now a 400 response w/ error message in body shows as 2 errors when it is really just 1
	See multierror README under "Customizing the formatting of the errors" */

	return errors.ErrorOrNil()
}
