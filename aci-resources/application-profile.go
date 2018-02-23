package cage

type AppProfile struct {
  BaseAttributes
}
/* NewTenant creates a new Tenant with the appropriate default values */
func NewAppProfile(name string, alias string, descr string, belongsTo BaseAttributes) *AppProfile {

  resourceName := fmt.Sprintf("ap-%s", name)
  attrs := BaseAttributes{
    Name: name,
    NameAlias: alias,
    Description: descr,
    Status: "created",
    ObjectClass: "fvAp",
    Status: "created",
    DN: fmt.Sprintf("%s/%s", belongsTo.DN, resourceName),
    RN: resourceName
  }

  a := AppProfile{BaseAttributes: attrs}
  //Do any additional construction logic specific to the EPG here
  return &a
}
