package main

var bridgeDomainServiceInstance *ResourceService

func GetBridgeDomainService() *ResourceService {
	if bridgeDomainServiceInstance == nil {
		bridgeDomainServiceInstance = &ResourceService{
			ObjectClass: "fvBridgeDomain",
			New:         NewBridgeDomain,
			FromJSON:    BridgeDomainFromJSON,
		}
	}
	return bridgeDomainServiceInstance
}
