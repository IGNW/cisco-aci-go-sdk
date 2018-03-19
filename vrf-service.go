package main

var vrfServiceInstance *ResourceService

func GetVRFService() *ResourceService {
	if vrfServiceInstance == nil {
		vrfServiceInstance = &ResourceService{
			ObjectClass: "fvVRF",
			New:         NewVRF,
			FromJSON:    VRFFromJSON,
		}
	}
	return vrfServiceInstance
}
