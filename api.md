


# vmware-rest-proxy
  

## Informations

### Version

0.1.0

### Contact

DO!DevOps info@dodevops.io http://dodevops.io

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## Access control

### Security Schemes

#### BasicAuth



> **Type**: basic

## All endpoints

###  datastore

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /datastores | [get datastores](#get-datastores) | Retrieve a list of datastores |
  


###  host

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /hosts | [get hosts](#get-hosts) | Retrieve a list of ESXi hosts |
  


###  operations

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /status | [get status](#get-status) | Checks whether the service is running |
  


###  vm

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /vms | [get vms](#get-vms) | Retrieve a list of all vms |
| GET | /vms/{id}/fqdn | [get vms ID fqdn](#get-vms-id-fqdn) | Get fqdn of VM |
| GET | /vms/{id}/info | [get vms ID info](#get-vms-id-info) | Get informational data about a VM |
| GET | /vms/{id}/tags | [get vms ID tags](#get-vms-id-tags) | Retrieve tags |
  


## Paths

### <span id="get-datastores"></span> Retrieve a list of datastores (*GetDatastores*)

```
GET /datastores
```

Fetches a list of registered datastores in the vCenter

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-datastores-200) | OK | OK |  | [schema](#get-datastores-200-schema) |
| [400](#get-datastores-400) | Bad Request | Invalid request |  | [schema](#get-datastores-400-schema) |
| [401](#get-datastores-401) | Unauthorized | Authorization is required |  | [schema](#get-datastores-401-schema) |

#### Responses


##### <span id="get-datastores-200"></span> 200 - OK
Status: OK

###### <span id="get-datastores-200-schema"></span> Schema
   
  

[EndpointsDatastores](#endpoints-datastores)

##### <span id="get-datastores-400"></span> 400 - Invalid request
Status: Bad Request

###### <span id="get-datastores-400-schema"></span> Schema

##### <span id="get-datastores-401"></span> 401 - Authorization is required
Status: Unauthorized

###### <span id="get-datastores-401-schema"></span> Schema

### <span id="get-hosts"></span> Retrieve a list of ESXi hosts (*GetHosts*)

```
GET /hosts
```

Fetches a list of registered ESXi hosts in the vCenter

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-hosts-200) | OK | OK |  | [schema](#get-hosts-200-schema) |
| [400](#get-hosts-400) | Bad Request | Invalid request |  | [schema](#get-hosts-400-schema) |
| [401](#get-hosts-401) | Unauthorized | Authorization is required |  | [schema](#get-hosts-401-schema) |

#### Responses


##### <span id="get-hosts-200"></span> 200 - OK
Status: OK

###### <span id="get-hosts-200-schema"></span> Schema
   
  

[EndpointsHosts](#endpoints-hosts)

##### <span id="get-hosts-400"></span> 400 - Invalid request
Status: Bad Request

###### <span id="get-hosts-400-schema"></span> Schema

##### <span id="get-hosts-401"></span> 401 - Authorization is required
Status: Unauthorized

###### <span id="get-hosts-401-schema"></span> Schema

### <span id="get-status"></span> Checks whether the service is running (*GetStatus*)

```
GET /status
```

Just responses with a 200 to signal that the service is running

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-status-200) | OK | OK |  | [schema](#get-status-200-schema) |

#### Responses


##### <span id="get-status-200"></span> 200 - OK
Status: OK

###### <span id="get-status-200-schema"></span> Schema
   
  

[EndpointsStatus](#endpoints-status)

### <span id="get-vms"></span> Retrieve a list of all vms (*GetVms*)

```
GET /vms
```

Fetches a list of vms from the vCenter

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-vms-200) | OK | OK |  | [schema](#get-vms-200-schema) |
| [400](#get-vms-400) | Bad Request | Invalid request |  | [schema](#get-vms-400-schema) |
| [401](#get-vms-401) | Unauthorized | Authorization is required |  | [schema](#get-vms-401-schema) |

#### Responses


##### <span id="get-vms-200"></span> 200 - OK
Status: OK

###### <span id="get-vms-200-schema"></span> Schema
   
  

[EndpointsVMS](#endpoints-vm-s)

##### <span id="get-vms-400"></span> 400 - Invalid request
Status: Bad Request

###### <span id="get-vms-400-schema"></span> Schema

##### <span id="get-vms-401"></span> 401 - Authorization is required
Status: Unauthorized

###### <span id="get-vms-401-schema"></span> Schema

### <span id="get-vms-id-fqdn"></span> Get fqdn of VM (*GetVmsIDFqdn*)

```
GET /vms/{id}/fqdn
```

Try to find out the fqdn of the given VM using the guest tools

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | ID of VM |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-vms-id-fqdn-200) | OK | OK |  | [schema](#get-vms-id-fqdn-200-schema) |
| [400](#get-vms-id-fqdn-400) | Bad Request | Invalid request |  | [schema](#get-vms-id-fqdn-400-schema) |
| [401](#get-vms-id-fqdn-401) | Unauthorized | Authorization is required |  | [schema](#get-vms-id-fqdn-401-schema) |

#### Responses


##### <span id="get-vms-id-fqdn-200"></span> 200 - OK
Status: OK

###### <span id="get-vms-id-fqdn-200-schema"></span> Schema
   
  

[EndpointsFQDN](#endpoints-f-q-d-n)

##### <span id="get-vms-id-fqdn-400"></span> 400 - Invalid request
Status: Bad Request

###### <span id="get-vms-id-fqdn-400-schema"></span> Schema

##### <span id="get-vms-id-fqdn-401"></span> 401 - Authorization is required
Status: Unauthorized

###### <span id="get-vms-id-fqdn-401-schema"></span> Schema

### <span id="get-vms-id-info"></span> Get informational data about a VM (*GetVmsIDInfo*)

```
GET /vms/{id}/info
```

Find out some information about a VM and return them

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | ID of VM |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-vms-id-info-200) | OK | OK |  | [schema](#get-vms-id-info-200-schema) |
| [400](#get-vms-id-info-400) | Bad Request | Invalid request |  | [schema](#get-vms-id-info-400-schema) |
| [401](#get-vms-id-info-401) | Unauthorized | Authorization is required |  | [schema](#get-vms-id-info-401-schema) |

#### Responses


##### <span id="get-vms-id-info-200"></span> 200 - OK
Status: OK

###### <span id="get-vms-id-info-200-schema"></span> Schema
   
  

[APIVMInfo](#api-vm-info)

##### <span id="get-vms-id-info-400"></span> 400 - Invalid request
Status: Bad Request

###### <span id="get-vms-id-info-400-schema"></span> Schema

##### <span id="get-vms-id-info-401"></span> 401 - Authorization is required
Status: Unauthorized

###### <span id="get-vms-id-info-401-schema"></span> Schema

### <span id="get-vms-id-tags"></span> Retrieve tags (*GetVmsIDTags*)

```
GET /vms/{id}/tags
```

Retrieve tags  and their categories for a vm

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | ID of VM |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-vms-id-tags-200) | OK | OK |  | [schema](#get-vms-id-tags-200-schema) |
| [400](#get-vms-id-tags-400) | Bad Request | Invalid request |  | [schema](#get-vms-id-tags-400-schema) |
| [401](#get-vms-id-tags-401) | Unauthorized | Authorization is required |  | [schema](#get-vms-id-tags-401-schema) |

#### Responses


##### <span id="get-vms-id-tags-200"></span> 200 - OK
Status: OK

###### <span id="get-vms-id-tags-200-schema"></span> Schema
   
  

[EndpointsTags](#endpoints-tags)

##### <span id="get-vms-id-tags-400"></span> 400 - Invalid request
Status: Bad Request

###### <span id="get-vms-id-tags-400-schema"></span> Schema

##### <span id="get-vms-id-tags-401"></span> 401 - Authorization is required
Status: Unauthorized

###### <span id="get-vms-id-tags-401-schema"></span> Schema

## Models

### <span id="api-datastore"></span> api.Datastore


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| capacity | integer| `int64` |  | |  |  |
| datastore | string| `string` |  | |  |  |
| free_space | integer| `int64` |  | |  |  |
| name | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |



### <span id="api-host"></span> api.Host


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| connection_state | string| `string` |  | |  |  |
| host | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| power_state | string| `string` |  | |  |  |



### <span id="api-vm"></span> api.VM


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | |  |  |
| power_state | string| `string` |  | |  |  |
| vm | string| `string` |  | |  |  |



### <span id="api-vm-info"></span> api.VMInfo


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cpu_cores | integer| `int64` |  | |  |  |
| name | string| `string` |  | |  |  |
| provisioned_ram | integer| `int64` |  | |  |  |
| provisioned_storage | integer| `int64` |  | |  |  |
| used_storage | integer| `int64` |  | |  |  |



### <span id="api-vm-tag"></span> api.VMTag


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| category | string| `string` |  | | Category holds the tag category |  |
| value | string| `string` |  | | Value holds the value of the tag |  |



### <span id="endpoints-datastores"></span> endpoints.Datastores


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| datastores | [EndpointsDatastoresResult](#endpoints-datastores-result)| `EndpointsDatastoresResult` |  | |  |  |



### <span id="endpoints-datastores-result"></span> endpoints.DatastoresResult


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| count | integer| `int64` |  | |  |  |
| datastores | [][APIDatastore](#api-datastore)| `[]*APIDatastore` |  | |  |  |



### <span id="endpoints-f-q-d-n"></span> endpoints.FQDN


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| fqdn | string| `string` |  | |  |  |



### <span id="endpoints-hosts"></span> endpoints.Hosts


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| hosts | [EndpointsHostsResult](#endpoints-hosts-result)| `EndpointsHostsResult` |  | |  |  |



### <span id="endpoints-hosts-result"></span> endpoints.HostsResult


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| count | integer| `int64` |  | |  |  |
| hosts | [][APIHost](#api-host)| `[]*APIHost` |  | |  |  |



### <span id="endpoints-status"></span> endpoints.Status


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="endpoints-tags"></span> endpoints.Tags


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| tags | [EndpointsTagsResult](#endpoints-tags-result)| `EndpointsTagsResult` |  | |  |  |



### <span id="endpoints-tags-result"></span> endpoints.TagsResult


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| count | integer| `int64` |  | |  |  |
| tags | [][APIVMTag](#api-vm-tag)| `[]*APIVMTag` |  | |  |  |



### <span id="endpoints-vm-s"></span> endpoints.VMS


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| vms | [EndpointsVMSResult](#endpoints-vm-s-result)| `EndpointsVMSResult` |  | |  |  |



### <span id="endpoints-vm-s-result"></span> endpoints.VMSResult


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| count | integer| `int64` |  | |  |  |
| vms | [][APIVM](#api-vm)| `[]*APIVM` |  | |  |  |


