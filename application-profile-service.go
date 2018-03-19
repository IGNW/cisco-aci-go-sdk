package main

var appProfileServiceInstance *ResourceService

func GetAppProfileService() *ResourceService {
	if appProfileServiceInstance == nil {
		appProfileServiceInstance = &ResourceService{
			ObjectClass: "fvAppProfile",
			New:         NewAppProfile,
			FromJSON:    AppProfileFromJSON,
		}
	}
	return appProfileServiceInstance
}
