# NFService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceInstanceId** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**InterPlmnFqdn** | Pointer to **string** |  | [optional] 
**IpEndPoints** | Pointer to [**[]IpEndPoint**](IpEndPoint.md) |  | [optional] 
**ApiPrfix** | Pointer to **string** |  | [optional] 
**AllowedPlmns** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**AllowedNfTypes** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**AllowedNssais** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 

## Methods

### NewNFService

`func NewNFService() *NFService`

NewNFService instantiates a new NFService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFServiceWithDefaults

`func NewNFServiceWithDefaults() *NFService`

NewNFServiceWithDefaults instantiates a new NFService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceInstanceId

`func (o *NFService) GetServiceInstanceId() string`

GetServiceInstanceId returns the ServiceInstanceId field if non-nil, zero value otherwise.

### GetServiceInstanceIdOk

`func (o *NFService) GetServiceInstanceIdOk() (*string, bool)`

GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstanceId

`func (o *NFService) SetServiceInstanceId(v string)`

SetServiceInstanceId sets ServiceInstanceId field to given value.

### HasServiceInstanceId

`func (o *NFService) HasServiceInstanceId() bool`

HasServiceInstanceId returns a boolean if a field has been set.

### GetServiceName

`func (o *NFService) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *NFService) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *NFService) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *NFService) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetVersion

`func (o *NFService) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NFService) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NFService) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NFService) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSchema

`func (o *NFService) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *NFService) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *NFService) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *NFService) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetFqdn

`func (o *NFService) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *NFService) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *NFService) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *NFService) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetInterPlmnFqdn

`func (o *NFService) GetInterPlmnFqdn() string`

GetInterPlmnFqdn returns the InterPlmnFqdn field if non-nil, zero value otherwise.

### GetInterPlmnFqdnOk

`func (o *NFService) GetInterPlmnFqdnOk() (*string, bool)`

GetInterPlmnFqdnOk returns a tuple with the InterPlmnFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterPlmnFqdn

`func (o *NFService) SetInterPlmnFqdn(v string)`

SetInterPlmnFqdn sets InterPlmnFqdn field to given value.

### HasInterPlmnFqdn

`func (o *NFService) HasInterPlmnFqdn() bool`

HasInterPlmnFqdn returns a boolean if a field has been set.

### GetIpEndPoints

`func (o *NFService) GetIpEndPoints() []IpEndPoint`

GetIpEndPoints returns the IpEndPoints field if non-nil, zero value otherwise.

### GetIpEndPointsOk

`func (o *NFService) GetIpEndPointsOk() (*[]IpEndPoint, bool)`

GetIpEndPointsOk returns a tuple with the IpEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpEndPoints

`func (o *NFService) SetIpEndPoints(v []IpEndPoint)`

SetIpEndPoints sets IpEndPoints field to given value.

### HasIpEndPoints

`func (o *NFService) HasIpEndPoints() bool`

HasIpEndPoints returns a boolean if a field has been set.

### GetApiPrfix

`func (o *NFService) GetApiPrfix() string`

GetApiPrfix returns the ApiPrfix field if non-nil, zero value otherwise.

### GetApiPrfixOk

`func (o *NFService) GetApiPrfixOk() (*string, bool)`

GetApiPrfixOk returns a tuple with the ApiPrfix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPrfix

`func (o *NFService) SetApiPrfix(v string)`

SetApiPrfix sets ApiPrfix field to given value.

### HasApiPrfix

`func (o *NFService) HasApiPrfix() bool`

HasApiPrfix returns a boolean if a field has been set.

### GetAllowedPlmns

`func (o *NFService) GetAllowedPlmns() PlmnId`

GetAllowedPlmns returns the AllowedPlmns field if non-nil, zero value otherwise.

### GetAllowedPlmnsOk

`func (o *NFService) GetAllowedPlmnsOk() (*PlmnId, bool)`

GetAllowedPlmnsOk returns a tuple with the AllowedPlmns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPlmns

`func (o *NFService) SetAllowedPlmns(v PlmnId)`

SetAllowedPlmns sets AllowedPlmns field to given value.

### HasAllowedPlmns

`func (o *NFService) HasAllowedPlmns() bool`

HasAllowedPlmns returns a boolean if a field has been set.

### GetAllowedNfTypes

`func (o *NFService) GetAllowedNfTypes() []NFType`

GetAllowedNfTypes returns the AllowedNfTypes field if non-nil, zero value otherwise.

### GetAllowedNfTypesOk

`func (o *NFService) GetAllowedNfTypesOk() (*[]NFType, bool)`

GetAllowedNfTypesOk returns a tuple with the AllowedNfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNfTypes

`func (o *NFService) SetAllowedNfTypes(v []NFType)`

SetAllowedNfTypes sets AllowedNfTypes field to given value.

### HasAllowedNfTypes

`func (o *NFService) HasAllowedNfTypes() bool`

HasAllowedNfTypes returns a boolean if a field has been set.

### GetAllowedNssais

`func (o *NFService) GetAllowedNssais() []Snssai`

GetAllowedNssais returns the AllowedNssais field if non-nil, zero value otherwise.

### GetAllowedNssaisOk

`func (o *NFService) GetAllowedNssaisOk() (*[]Snssai, bool)`

GetAllowedNssaisOk returns a tuple with the AllowedNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNssais

`func (o *NFService) SetAllowedNssais(v []Snssai)`

SetAllowedNssais sets AllowedNssais field to given value.

### HasAllowedNssais

`func (o *NFService) HasAllowedNssais() bool`

HasAllowedNssais returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


