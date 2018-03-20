package main

import (
	"fmt"
	"net/http"

	"github.com/Jeffail/gabs"
)

// Services represents a collection of service objects within the client
// will acess.client().AppProfiles ...
type Services struct {
	AppProfiles   *ResourceService
	BridgeDomains *ResourceService
	Contracts     *ResourceService
	EPGs          *ResourceService
	Filters       *ResourceService
	Subjects      *ResourceService
	Subnets       *ResourceService
	Tenants       *ResourceService
}

// ServiceInterface defines the interface that each service object will expose
type ServiceInterface interface {
	Save(ResourceInterface) error
	GetAll() ([]*ResourceInterface, error)
	Get(*map[string]string) ([]*ResourceInterface, error)
	Delete(domainName string) (data *gabs.Container, response *http.Response, err error)
	New(name string, nameAlias string, descr string)
}

type ResourceGenerator func(string, string, string) ResourceInterface
type ResourceDecoder func(*gabs.Container) (ResourceInterface, error)

type ResourceService struct {
	ObjectClass string
	New         ResourceGenerator
	FromJSON    ResourceDecoder
}

func (s ResourceService) client() *Client {
	return GetClient()
}
func (s ResourceService) Save(r ResourceInterface) (data *gabs.Container, response *http.Response, err error) {

	data = r.GetAPIPayload()

	data.Set("created, modified", s.ObjectClass, "attributes", "status")

	json := r.GetAPIPayload()
	method := "POST"
	path := fmt.Sprintf("/api/node/mo/%s-%s.json", s.ObjectClass, r.getResourceName())

	req, err := s.client().newAuthdRequest(method, path, json)
	if err != nil {
		fmt.Printf("\nGot Error While Saving, Auth'd Request failed w/ %v", err)
		return nil, nil, err
	}

	return s.client().do(req)
}

func (s ResourceService) Get(params *map[string]string) ([]*ResourceInterface, error) {

	path := fmt.Sprintf("/api/node/mo/%s", s.ObjectClass)

	if params != nil {
		queryString := "?"
		paramCount := 0

		for key, value := range *params {
			if key != "" && value != "" {
				if paramCount > 0 {
					queryString += "&"
				}
				queryString += fmt.Sprintf("%s=%s", key, value)
				paramCount++
			}
		}

		if paramCount > 0 {
			path += queryString
		}
	}

	req, err := s.client().newAuthdRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	_, _, err = s.client().do(req)

	return nil, nil
}

func (s ResourceService) GetAll() ([]*ResourceInterface, error) {
	path := fmt.Sprintf("/api/node/class/%s", s.ObjectClass)
	req, err := s.client().newAuthdRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	_, _, err = s.client().do(req)

	return nil, nil
}

func (s ResourceService) Delete(dn string) (data *gabs.Container, response *http.Response, err error) {
	rawjson := fmt.Sprintf(`{"%s" : { "attributes" : {}}`, s.ObjectClass)

	data, err = gabs.Consume(rawjson)

	data.Set("deleted", s.ObjectClass, "attributes", "status")
	data.Set(dn, s.ObjectClass, "attributes", "dn")

	req, err := s.client().newAuthdRequest("POST", "/api/node/mo/uni", data)
	if err != nil {
		return nil, nil, err
	}
	data, response, err = s.client().do(req)

	return data, response, err
}
