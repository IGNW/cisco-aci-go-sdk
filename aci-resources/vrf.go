package cage

type VRF struct {
  *BaseAttributes
  NameAlias string
}

func NewVRF(name string, alias string, belongsTo string) VRF {
  v.Status = "created"
  v.Name = name
  v.NameAlias = alias
  v.DN = fmt.Sprintf("uni/%s/%s", belongsTo, name)
  v.RN = fmt.Sprintf("ctx-%s", name)
}
