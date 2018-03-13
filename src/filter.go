package cage

type Filter struct {
	ResourceAttributes
	Subjects []*Subjects
	Parent   *Tenant
}
