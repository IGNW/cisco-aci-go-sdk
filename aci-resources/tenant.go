package cage

type Tenant struct {
  *BaseAttributes
  NameAlias string
  Description string
}
/* NewTenant creates a new Tenant with the appropriate default values */
func NewTenant(name string, alias string, descr string) {
  t = new(Tenant)
  t.Name = name
  t.NameAlias =  alias
  t.Description = descr
  t.Status = "created"
  t.DN = fmt.Sprintf("uni/%s", name)
  t.RN = fmt.Sprintf("tn-%s", name)
  return t
}
