package main

var epgServiceInstance *ResourceService

func GetEPGService() *ResourceService {
	if epgServiceInstance == nil {
		epgServiceInstance = &ResourceService{
			ObjectClass: "fvCTx",
			New:         NewEPG,
			FromJSON:    EPGFromJSON,
		}
	}
	return epgServiceInstance
}
