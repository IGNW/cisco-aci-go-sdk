package cage

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

/**ResourceAttributes represents the generic properties of *most*
ACI resources
*/
type ResourceAttributes struct {
	Name         string
	ResourceName string
	DomainName   string
	Status       string
	Tags         []Tag
	NameAlias    string
	Description  string
	ObjectClass  string
	Parent       ParentInterface
}

/** Defines the methods an object must have to be considered to have implemented the ResourceInterface,
which can be used as an arugment type in a method
*/
type ResourceInterface interface {
	SetResourceName() string
	SetDomainName() string
	CreateAPIPayload() gabs.Container
}

/** Defines the methods an object must have to be considered to have implemented the ParentInterface,
which can be used as an arugment type in a method
This is seperated from the ResourceInterface since not all resources can be parents to other resources
See also: squares vs rectangles
*/
type ParentInterface interface {
	AddChild(child ResourceInterface)
	AddChildren(children []ResourceInterface)
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
				"descr": "",
				"dn": "",
				"name": "",
				"nameAlias": "",
				"rn": "",
				"status": ""
			}
		}
	}`, r.ObjectClass))

	return gabs.ParseJSON(containerJSON)

}

func (r *ResourceAttributes) AddDefaultPropsToPayload(data *gabs.Container) {

	// set value -> key...
	data.Set(r.Name, r.ObjectClass, "attributes", "name")

	data.Set(r.NameAlias, r.ObjectClass, "attributes", "nameAlias")

	data.Set(r.Description, r.ObjectClass, "attributes", "descr")

	data.Set(r.ResourceName, r.ObjectClass, "attributes", "rn")

	data.Set(r.DomainName, r.ObjectClass, "attributes", "dn")

	data.Set(r.Status, r.ObjectClass, "attributes", "status")

	/** Create our empty children array.
	Cisco APIC expects this, or at least implments it in its GUI
	and it will save us a lof of 'if exists' checking later
	*/
	data.Array(r.ObjectClass, "children")
}
func (r *ResourceAttributes) AddTagsToPayload(data *gabs.Container) {
	for _, tag := range r.Tags {
		data.ArrayAppend(tag.AsPayLoadFormat(), r.ObjectClass, "children")
	}
}
