package models

import (
	"github.com/Jeffail/gabs"
)

// Represents the base attributes of all ACI models.
type ResourceAttributes struct {
	Name         string
	ResourceName string
	DomainName   string
	Status       string
	Tags         []Tag
	// NameAlias    string
	Description string
	ObjectClass string
	Parent      ResourceInterface
}

/** Defines the methods an object must have to be considered to have implemented the ResourceInterface,
which can be used as an argument type in a method
*/
type ResourceInterface interface {
	GetResourceName() string
	GetParent() ResourceInterface
	ToMap() map[string]string
}

// Convert ResourceAttributes model to a string map.
func (r ResourceAttributes) ToMap() map[string]string {
	return map[string]string{"dn": r.DomainName, "name": r.Name, "descr": r.Description, "status": r.Status}
}

func NewResourceAttributes(model map[string]string) ResourceAttributes {
	newModel := ResourceAttributes{
		DomainName: model["dn"],
		Name:       model["name"],
		// NameAlias:	  model["nameAlias"],
		Description:  model["descr"],
		Status:       model["status"],
		ResourceName: model["rn"],
		ObjectClass:  model["objectClass"],
	}

	return newModel
}

// NewResourceAttributesMap will construct a string map from reading ResourceAttributes values that can be converted to the type.
func NewResourceAttributesMap() map[string]string {
	return map[string]string{"dn": "", "name": "", "descr": "", "status": ""}
}

// Convert string map to ResourceAttributes.
func (r ResourceAttributes) FromMap(model map[string]string) ResourceInterface {

	newModel := ResourceAttributes{
		DomainName: model["dn"],
		Name:       model["name"],
		// NameAlias:    model["nameAlias"],
		Description:  model["descr"],
		Status:       model["status"],
		ResourceName: model["rn"],
		ObjectClass:  model["objectClass"],
	}

	return &newModel
}

// Convert JSON string map to ResourceAttributes model.
func (r ResourceAttributes) FromJSONMap(attributes map[string]string) ResourceAttributes {

	return ResourceAttributes{
		DomainName: attributes["dn"],
		Name:       attributes["name"],
		// NameAlias:    attributes["nameAlias"],
		Description:  attributes["descr"],
		Status:       attributes["status"],
		ResourceName: attributes["rn"],
		ObjectClass:  attributes["objectClass"],
	}
}

func (r ResourceAttributes) GetResourceName() string {
	return r.ResourceName
}

// Parse ACI boolean values (yes/no)
func (r ResourceAttributes) ParseBool(value string) bool {

	switch value {
	case "yes", "Yes", "YES":
		return true
	case "no", "No", "NO":
		return false
	}

	return false
}

// Format ACI boolean values (yes/no)
func (r ResourceAttributes) FormatBool(value bool) string {
	if value {
		return "yes"
	} else {
		return "no"
	}
}

/*
func (r *ResourceAttributes) CreateDefaultPayload() *gabs.Container {

	payloadContainer, _ := r.CreateEmptyJSONContainer()

	r.AddDefaultPropsToPayload(payloadContainer)

	r.AddTagsToPayload(payloadContainer)

	return payloadContainer
}
*/

func (r *ResourceAttributes) AddDefaultPropsToPayload(data *gabs.Container) {

	// set value -> key...
	data.Set(r.Name, r.ObjectClass, "attributes", "name")

	// data.Set(r.NameAlias, r.ObjectClass, "attributes", "nameAlias")

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
