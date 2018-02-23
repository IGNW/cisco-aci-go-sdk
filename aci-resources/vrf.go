package cage

type VRF struct {
  BaseAttributes
}

func NewVRF(name string, alias string, descr, string, belongsTo BaseAttributes) *VRF {

  resourceName := fmt.Sprintf("ctx-%s", name)
  attrs := BaseAttributes{
    Name: name,
    NameAlias: alias,
    Description: descr,
    Status: "created",
    ObjectClass: "fvCtx",
    DN: fmt.Sprintf("%s/%s", belongsTo.DN, resourceName),
    RN: resourceName
  }

  v := VRF{BaseAttributes: attrs}
  //Do any additional construction logic here.
  return &v
}
