package main

var contractServiceInstance *ResourceService

func GetContractService() *ResourceService {
	if contractServiceInstance == nil {
		contractServiceInstance = &ResourceService{
			ObjectClass: "fvContract",
			New:         NewContract,
			FromJSON:    ContractFromJSON,
		}
	}
	return contractServiceInstance
}
