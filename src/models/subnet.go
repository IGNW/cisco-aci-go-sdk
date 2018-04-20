package models

import "strings"

const SUBNET_RESOURCE_NAME_PREFIX = "subnet"
const SUBNET_OBJECT_CLASS = "fvSubnet"

// Represents an ACI Bridge Domain Subnet.
// See: https://pubhub.devnetcloud.com/media/apic-mim-ref-311/docs/MO-fvSubnet.html
type Subnet struct {
	ResourceAttributes
	Control   string `oneof=unspecified querier nd no-default-gateway`
	IpAddress string
	Preferred bool
	Scope     []string `oneof=public private shared`
	Virtual   bool
}

func (s *Subnet) GetObjectClass() string {
	return SUBNET_OBJECT_CLASS
}

func (s *Subnet) GetResourcePrefix() string {
	return SUBNET_RESOURCE_NAME_PREFIX
}

func (s *Subnet) HasParent() bool {
	return true
}

func (s *Subnet) ToMap() map[string]string {
	var model = s.ResourceAttributes.ToMap()

	// The subnet control state. The control can be specific protocols applied to the subnet such as IGMP Snooping.
	model["ctrl"] = s.Control

	// The IP address and mask of the default gateway.
	model["ip"] = s.IpAddress

	// Indicates if the subnet is preferred (primary) over the available alternatives. Only one preferred subnet is allowed.
	model["preferred"] = s.FormatBool(s.Preferred)

	model["scope"] = strings.Join(s.Scope, ",")

	// Treated as virtual IP address. Used in case of BD extended to multiple sites.
	model["virtual"] = s.FormatBool(s.Virtual)

	return model
}

// NewSubnet will construct a Subnet from a string map.
func NewSubnet(model map[string]string) *Subnet {

	s := Subnet{NewResourceAttributes(model),
		"",
		"",
		false,
		nil,
		false,
	}

	// The subnet control state. The control can be specific protocols applied to the subnet such as IGMP Snooping.
	s.Control = model["ctrl"]

	// The IP address and mask of the default gateway.
	s.IpAddress = model["ip"]

	// Indicates if the subnet is preferred (primary) over the available alternatives. Only one preferred subnet is allowed.
	s.Preferred = s.ParseBool(model["preferred"])

	s.Scope = strings.Split(model["scope"], ",")

	// Treated as virtual IP address. Used in case of BD extended to multiple sites.
	s.Virtual = s.ParseBool(model["virtual"])

	return &s
}

// NewSubnetMap will construct a string map from reading Subnet values that can be converted to the type.
func NewSubnetMap() map[string]string {

	m := NewResourceAttributesMap()

	m["ctrl"] = "nd"
	m["ip"] = ""
	m["preferred"] = "no"
	m["scope"] = "private"

	return m
}

/*
"fvSubnet": {
	"attributes": {
		"childAction": "",
		"ctrl": "nd",
		"descr": "",
		"extMngdBy": "",
		"ip": "192.168.101.1/24",
		"lcOwn": "local",
		"modTs": "2018-02-23T04:36:17.500+00:00",
		"monPolDn": "uni/tn-common/monepg-default",
		"name": "VLAN",
		"nameAlias": "",
		"preferred": "no",
		"rn": "subnet-[192.168.101.1/24]",
		"scope": "private",
		"status": "",
		"uid": "15374",
		"virtual": "no"
	}
}
*/
