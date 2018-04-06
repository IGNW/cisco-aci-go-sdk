package models

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

/**ResourceAttributes represents the generic properties of *most*
ACI resources
*/
type ResourceAttributes struct {
	// client       *Client
	Name         string
	ResourceName string
	DomainName   string
	Status       string
	Tags         []Tag
	NameAlias    string
	Description  string
	ObjectClass  string
	Parent       ResourceInterface
}

/** Defines the methods an object must have to be considered to have implemented the ResourceInterface,
which can be used as an arugment type in a method
*/
type ResourceInterface interface {
	GetAPIPayload() *gabs.Container
	GetResourceName() string
	GetParent() ResourceInterface
}

func (r ResourceAttributes) GetAPIPayload() *gabs.Container {
	return r.CreateDefaultPayload()
}

func (r ResourceAttributes) GetResourceName() string {
	return r.ResourceName
}

func (r *ResourceAttributes) CreateDefaultPayload() *gabs.Container {

	payloadContainer, _ := r.CreateEmptyJSONContainer()

	r.AddDefaultPropsToPayload(payloadContainer)

	r.AddTagsToPayload(payloadContainer)

	return payloadContainer
}

func (r *ResourceAttributes) CreateEmptyJSONContainer() (*gabs.Container, error) {
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
			}
		}
	}`, r.ObjectClass))

	return gabs.ParseJSON(containerJSON)

}

func (r *ResourceAttributes) AddDefaultPropsToPayload(data *gabs.Container) {

	// set value -> key...
	data.Set(r.Name, r.ObjectClass, "attributes", "name")

	//data.Set(r.NameAlias, r.ObjectClass, "attributes", "nameAlias")

	data.Set(r.Description, r.ObjectClass, "attributes", "descr")

	data.Set(r.ResourceName, r.ObjectClass, "attributes", "rn")

	data.Set(r.Status, r.ObjectClass, "attributes", "status")

	/** Create our empty children array.
	Cisco APIC expects this, or at least implments it in its GUI
	and it will save us a lof of 'if exists' checking later
	*/
	data.Array(r.ObjectClass, "children")

	//fmt.Printf("Payload:\n%s\n\n", data.String())
}
func (r *ResourceAttributes) AddTagsToPayload(data *gabs.Container) {
	for _, tag := range r.Tags {
		data.ArrayAppend(tag.AsPayLoadFormat(), r.ObjectClass, "children")
	}
}

// AddTag adds a tag to a given resource. Returns bool for status, and err if any was encountered
func (r *ResourceAttributes) AddTag(name string) {
	r.Tags = append(r.Tags, NewTag(name))
}

func (r *ResourceAttributes) GetParent() ResourceInterface {
	return r.Parent
}

func (r *ResourceAttributes) SetParent(parent ResourceInterface) {
	r.Parent = parent
}
