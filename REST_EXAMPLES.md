# Example REST calls for accessing ACI (APIC) resources

## Authentication
This snippet will autenticate using `curl` and cookies stored in `./.cookies`.

```bash
curl --insecure -c .cookies -X POST -H "Content-Type: application/json" -d '{"aaaUser":{"attributes":{"name":"admin", "pwd":"password"}}}' https://73.254.132.17:8443/api/aaaLogin.json
```

## Tenant

### Fetch All

```bash
curl --insecure -b .cookies -H "Content-Type: application/json" https://73.254.132.17:8443/api/class/fvTenant.json
```

## Fetch One

```bash
curl --insecure -b .cookies -H "Content-Type: application/json" https://73.254.132.17:8443/api/mo/uni/tn-IGNW.json
```

# Bridge Domain

## Fetch All

```bash
curl --insecure -b .cookies -H "Content-Type: application/json" https://73.254.132.17:8443/api/class/fvBD.json
```

## Fetch One

```bash
curl --insecure -b .cookies -H "Content-Type: application/json" https://73.254.132.17:8443/api/mo/uni/tn-infra/BD-default.json
```

# Pretty Printing

Install `jq` for the command line and pipe your `curl` to it, see example below.

```bash
curl --insecure -b .cookies -H "Content-Type: application/json" https://73.254.132.17:8443/api/mo/uni/tn-infra/BD-default.json | jq
```
