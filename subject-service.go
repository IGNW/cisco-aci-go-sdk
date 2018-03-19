package main

var subjectServiceInstance *ResourceService

func GetSubjectService() *ResourceService {
	if subjectServiceInstance == nil {
		subjectServiceInstance = &ResourceService{
			ObjectClass: "fvSubject",
			New:         NewSubject,
			FromJSON:    SubjectFromJSON,
		}
	}
	return subjectServiceInstance
}
