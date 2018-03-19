package main

var subnetServiceInstance *ResourceService

func GetSubnetService() *ResourceService {
	if subnetServiceInstance == nil {
		subnetServiceInstance = &ResourceService{
			ObjectClass: "fvSubnet",
			New:         NewSubnet,
			FromJSON:    SubnetFromJSON,
		}
	}
	return subnetServiceInstance
}
