package cage

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

type Tenant struct {
	ResourceAttributes
	L3NetIdentifier string
	VRFs            []*VRF
	BridgeDomains   []*BridgeDomain
	AppProfiles     []*AppProfile
	Contracts       []*Contract
	Filters         []*Filter
}

/* New creates a new Tenant with the appropriate default values */
func NewTenant(name string, alias string, descr string) Tenant {
	resourceName := fmt.Sprintf("tn-%s", name)

	t := Tenant{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "fvTenant",
		DomainName:   fmt.Sprintf("uni/%s", resourceName),
		ResourceName: resourceName,
	},
		"",
		nil,
		nil,
		nil,
		nil,
		nil,
	}
	//Do any additional construction logic here.
	return t
}

func (t *Tenant) CreateAPIPayload() *gabs.Container {
	return t.CreateDefaultPayload()
}

// AddVRF adds a VRF to the Tenants VRF list and sets the Parent prop of the VRF to the Tenant it was called from
func (t *Tenant) AddVRF(v *VRF) {
	v.SetParent(t)
	t.VRFs = append(t.VRFs, v)
}

// AddBridgeDomain adds a Domain to the Tenants BridgeDomain list and sets the Parent prop of the BridgeDomain to the Tenant it was called from
func (t *Tenant) AddBridgeDomain(bd *BridgeDomain) {
	bd.SetParent(t)
	t.BridgeDomains = append(t.BridgeDomains, bd)
}

// AddAppProfile adds a Domain to the Tenants AppProfile list and sets the Parent prop of the AppProfile to the Tenant it was called from
func (t *Tenant) AddAppProfile(ap *AppProfile) {
	ap.SetParent(t)
	t.AppProfiles = append(t.AppProfiles, ap)
}

// AddContract adds a Domain to the Tenants Contract list and sets the Parent prop of the Contract to the Tenant it was called from
func (t *Tenant) AddContract(c *Contract) {
	c.SetParent(t)
	t.Contracts = append(t.Contracts, c)
}

// AddFilter adds a Domain to the Tenants Filter list and sets the Parent prop of the Filter to the Tenant it was called from
func (t *Tenant) AddFilter(f *Filter) {
	f.SetParent(t)
	t.Filters = append(t.Filters, f)
}
