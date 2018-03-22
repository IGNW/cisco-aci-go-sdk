package main

import (
	"fmt"
	"net/http"

	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
)

// Services represents a collection of service objects within the client
// will acess.client().AppProfiles ...
type Services struct {
	// AppProfiles   *ResourceService
	// BridgeDomains *ResourceService
	// Contracts     *ResourceService
	// EPGs          *ResourceService
	// Filters       *ResourceService
	// Subjects      *ResourceService
	// Subnets       *ResourceService
	Tenants *TenantService
}

type ResourceGenerator func(string, string, string) ResourceInterface
type ResourceDecoder func(*gabs.Container) (ResourceInterface, error)

type ResourceService struct {
	ObjectClass string
}

func (s ResourceService) client() *Client {
	return GetClient()
}
func (s ResourceService) Save(r ResourceInterface) (err error) {

	data := r.GetAPIPayload()

	data.Set("created, modified", s.ObjectClass, "attributes", "status")

	json := r.GetAPIPayload()
	method := "POST"
	path := fmt.Sprintf("/api/node/mo/uni/%s.json", r.getResourceName())

	req, err := s.client().newAuthdRequest(method, path, json)
	if err != nil {
		fmt.Printf("\nGot Error While Saving, Auth'd Request failed w/ %v", err)
		return err
	}

	data, response, err := s.client().do(req)

	fmt.Printf("RESP: %#v\n\n", response)
	fmt.Printf("DATA: %#v\n\n", data)

	if err != nil {
		return err
	}

	if err = s.getResponseError(data); err != nil {
		return err
	}

	return nil
}

func (s ResourceService) Get(domainName string) (*gabs.Container, error) {

	path := fmt.Sprintf("api/mo/%s.json", s.ObjectClass)

	req, err := s.client().newAuthdRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	data, response, err := s.client().do(req)

	err = s.combineErrors(data, response, err)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func (s ResourceService) GetAll() (*gabs.Container, error) {
	path := fmt.Sprintf("/api//class/%s.json", s.ObjectClass)
	req, err := s.client().newAuthdRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	data, response, err := s.client().do(req)

	err = s.combineErrors(data, response, err)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s ResourceService) Delete(dn string) error {
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
			}
		}
	}`, s.ObjectClass))

	data, err := gabs.ParseJSON(containerJSON)

	if err != nil {
		fmt.Println(err)
	}

	_, err = data.Set("deleted", s.ObjectClass, "attributes", "status")
	if err != nil {
		fmt.Println(err)
	}

	_, err = data.Set(dn, s.ObjectClass, "attributes", "dn")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Payload: %s", data.String())

	path := fmt.Sprintf("/api/node/mo/%s.json", dn)

	req, err := s.client().newAuthdRequest("POST", path, data)

	if err != nil {
		return err
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

func (s ResourceService) getGabsValue(data *gabs.Container, valuePath string) string {
	// Not sure if this is Cisco or Gabs, but wow.
	// @TODO find a better way to extract values from gabs containers
	return data.Path(valuePath).Data().([]interface{})[0].(string)
}

//@TODO rename this, response should always refer to an http.Response for clarity, this is a response body dijested as a gabs Container
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
