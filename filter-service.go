package main

var filterServiceInstance *ResourceService

func GetFilterService() *ResourceService {
	if filterServiceInstance == nil {
		filterServiceInstance = &ResourceService{
			ObjectClass: "fvFilter",
			New:         NewFilter,
			FromJSON:    FilterFromJSON,
		}
	}
	return filterServiceInstance
}
