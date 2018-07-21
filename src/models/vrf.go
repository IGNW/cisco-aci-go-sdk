package models

const VRF_RESOURCE_PREFIX = "ctx"
const VRF_OBJECT_CLASS = "fvCtx"

// Represents an ACI VRF.
// See: https://pubhub.devnetcloud.com/media/apic-mim-ref-311/docs/MO-fvCtx.html
type VRF struct {
	ResourceAttributes
	Enforce              string `unenforced enforced`
	EnforcementDirection string `ingress egress`
	BridgeDomains        []*BridgeDomain
}

func (v *VRF) GetObjectClass() string {
	return VRF_OBJECT_CLASS
}

func (v *VRF) GetResourcePrefix() string {
	return VRF_RESOURCE_PREFIX
}

func (v *VRF) HasParent() bool {
	return true
}

func (v *VRF) ToMap() map[string]string {
	var model = v.ResourceAttributes.ToMap()

	// Treated as virtual IP address. Used in case of BD extended to multiple sites.
	model["pcEnfPref"] = v.Enforce
	model["pcEnfDir"] = v.EnforcementDirection

	return model
}

// NewVRF will construct a VRF from a string map.
func NewVRF(model map[string]string) *VRF {

	v := VRF{NewResourceAttributes(model),
		"",
		"",
		nil,
	}

	// The subnet control state. The control can be specific protocols applied to the subnet such as IGMP Snooping.
	v.Enforce = model["pcEnfPref"]
	v.EnforcementDirection = model["pcEnfDir"]

	return &v
}

// NewVRFMap will construct a string map from reading VRF values that can be converted to the type.
func NewVRFMap() map[string]string {

	m := NewResourceAttributesMap()

	m["pcEnfPref"] = "unenforced"
	m["pcEnfDir"] = "ingress"

	return m
}

// AddBridgeDomain adds a BridgeDomain to the VRF BridgeDomain list
func (v *VRF) AddBridgeDomain(bd *BridgeDomain) *VRF {
	v.BridgeDomains = append(v.BridgeDomains, bd)

	return v
}
