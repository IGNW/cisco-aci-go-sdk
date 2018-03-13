package cage

import "fmt"

// import "fmt"

type VRF struct {
	ResourceAttributes
	BridgeDomains []*BridgeDomain
}

func NewVRF(name string, alias string, descr, string, belongsTo BaseAttributes) *VRF {

	resourceName := fmt.Sprintf("ctx-%s", name)
	attrs := ResourceAttributes{
		Name:        name,
		NameAlias:   alias,
		Description: descr,
		Status:      "created",
		ObjectClass: "fvCtx",
		DN:          fmt.Sprintf("%s/%s", belongsTo.DN, resourceName),
		RN:          resourceName,
	}

	v := VRF{Attributes: attrs}
	//Do any additional construction logic here.
	return &v
}

// AddBridgeDomain adds a BridgeDomain to the VRF BridgeDomain list
func (v *VRF) AddBridgeDomain(bd *BridgeDomain) {
	v.BridgeDomains = append(v.BridgeDomains, bd)
}
