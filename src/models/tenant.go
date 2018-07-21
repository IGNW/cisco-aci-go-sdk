package models

const TENANT_RESOURCE_PREFIX = "tn"
const TENANT_OBJECT_CLASS = "fvTenant"

// Represents an ACI Tenant.
type Tenant struct {
	ResourceAttributes
	VRFs          []*VRF
	BridgeDomains []*BridgeDomain
	AppProfiles   []*AppProfile
	Contracts     []*Contract
	Filters       []*Filter
}

func (t *Tenant) GetObjectClass() string {
	return TENANT_OBJECT_CLASS
}

func (t *Tenant) GetResourcePrefix() string {
	return TENANT_RESOURCE_PREFIX
}

func (t *Tenant) HasParent() bool {
	return false
}

func (t *Tenant) ToMap() map[string]string {
	var model = t.ResourceAttributes.ToMap()
	return model
}

// NewTenant will construct a Tenant from a string map.
func NewTenant(model map[string]string) *Tenant {

	t := Tenant{NewResourceAttributes(model),
		nil,
		nil,
		nil,
		nil,
		nil,
	}

	return &t
}

// NewTenantMap will construct a string map from reading Tenant values that can be converted to the type.
func NewTenantMap() map[string]string {
	m := NewResourceAttributesMap()
	return m
}

// AddVRF adds a VRF to the Tenants VRF list and sets the Parent prop of the VRF to the Tenant it was called from
func (t *Tenant) AddVRF(v *VRF) *Tenant {
	v.SetParent(t)
	t.VRFs = append(t.VRFs, v)
	return t
}

// AddBridgeDomain adds a Domain to the Tenants BridgeDomain list and sets the Parent prop of the BridgeDomain to the Tenant it was called from
func (t *Tenant) AddBridgeDomain(bd *BridgeDomain) *Tenant {
	bd.SetParent(t)
	t.BridgeDomains = append(t.BridgeDomains, bd)

	return t
}

// AddAppProfile adds a Domain to the Tenants AppProfile list and sets the Parent prop of the AppProfile to the Tenant it was called from
func (t *Tenant) AddAppProfile(ap *AppProfile) *Tenant {
	ap.SetParent(t)
	t.AppProfiles = append(t.AppProfiles, ap)

	return t
}

// AddContract adds a Domain to the Tenants Contract list and sets the Parent prop of the Contract to the Tenant it was called from
func (t *Tenant) AddContract(c *Contract) *Tenant {
	c.SetParent(t)
	t.Contracts = append(t.Contracts, c)

	return t
}

// AddFilter adds a Domain to the Tenants Filter list and sets the Parent prop of the Filter to the Tenant it was called from
func (t *Tenant) AddFilter(f *Filter) *Tenant {
	f.SetParent(t)
	t.Filters = append(t.Filters, f)

	return t
}
