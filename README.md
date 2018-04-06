# Cisco ACI Go Enabler (main)

A limited functionality SDK derived from the APIC GUI with the primary goal of enabling a Terraform provider for Cisco ACI

## TODOS

- []    Create service providers and defintions for remaining resources.
        These should be able to be done programatically, but I was having some weird python behavior.
        See `scripts/new-resource.py`, maybe you'll have better luck getting it working. If not, copying from the Tenants service will work with some search and replace.

- []    Fill in object classes within resource and resource services
        Many resources have a "@TODO" for their object class (IE fvTenant for Tenants)

- []    Add required properties to resource objects
        All resource objects currently extend the ResourceAttributes struct which contains the common properties shared between resource types. 
        Unique properties should be added to the resource struct definition (IE tenant.go)

- []    Add properties from previous task to JSON encoder/decoder methods. All resource services have a `fromJSON` method to take response json 
        and turn it into and SDK object. Currently all the "convert to json" functionality is held within `ResourceServices.GetAPIPayload`.
        For semantic reasons this should be aligned with the json verbage IE `toJSON` 

- []    Refactor `fromJSON` method into a ResourceService method. Currently each service has a fully duplicated instance, and there's plenty of 
        opportunity to share.
        Individual services can overide this with their own `fromJSON` but still access it via ResourceService. 
        For example tenants service (ts) would be able to access via `ts.ResourceService.fromJSON` with it's own `ts.fromJSON` implementation
