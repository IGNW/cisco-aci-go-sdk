package cage

type Filter struct {
	ResourceAttributes
	Subjects []*Subject
	Parent   *Tenant
}
