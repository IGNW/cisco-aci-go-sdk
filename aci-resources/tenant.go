package cage

type Tenant struct {
  BaseAttributes
}
/* NewTenant creates a new Tenant with the appropriate default values */
func NewTenant(name string, alias string, descr string) *Tenant{
  resourceName := fmt.Sprintf("rn-%s", name)
  attrs := BaseAttributes{
    Name: name,
    NameAlias: alias,
    Description: descr,
    Status: "created",
    ObjectClass: "fvTenant",
    DN: fmt.Sprintf("uni/%s", resourceName),
    RN: resourceName
  }

  t := Tenant{BaseAttributes: attrs}
  //Do any additional construction logic here.
  return &t
}
