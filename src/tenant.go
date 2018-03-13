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
	AppProfile      []*AppProfile
	Contracts       []*Contract
	Filters         []*Filter
}

/* New creates a new Tenant with the appropriate default values */
func (Tenant) New(name string, alias string, descr string) Tenant {
	resourceName := fmt.Sprintf("tn-%s", name)

	t := Tenant{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "fvTenant",
		DomainName:   fmt.Sprintf("uni/%s", resourceName),
		ResourceName: resourceName,
	}}
	//Do any additional construction logic here.
	return t
}

func (t *Tenant) CreateAPIPayload() *gabs.Container {
	return t.CreateDefaultPayload()
}

// AddVRF adds a VRF to the Tenants VRF list and sets the Parent prop of the VRF to the Tenant it was called from
func (t *Tenant) AddVRF(v *VRF) {
	v.Parent = &t
	t.VRFs = append(t.VRFs, v)
}

// AddBridgeDomain adds a Domain to the Tenants BridgeDomain list and sets the Parent prop of the BridgeDomain to the Tenant it was called from
func (t *Tenant) AddBridgeDomain(bd *BridgeDomain) {
	bd.Parent = &t
	t.VRFs = append(t.BridgeDomains, bd)
}

// AddAppProfile adds a Domain to the Tenants AppProfile list and sets the Parent prop of the AppProfile to the Tenant it was called from
func (t *Tenant) AddAppProfile(ap *AppProfile) {
	ap.Parent = &t
	t.AppProfiles = append(t.AppProfiles, ap)
}

// AddContract adds a Domain to the Tenants Contract list and sets the Parent prop of the Contract to the Tenant it was called from
func (t *Tenant) AddContract(c *Contract) {
	c.Parent = &t
	t.Contracts = append(t.Contracts, c)
}

// AddFilter adds a Domain to the Tenants Filter list and sets the Parent prop of the Filter to the Tenant it was called from
func (t *Tenant) AddFilter(f *Filter) {
	f.Parent = &t
	t.Filters = append(t.Filters, f)
}
