# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API major version in URI (e.g. v1) | 
**Expiry** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**Resources** | Pointer to [**[]Resource**](Resource.md) | Resources supported by the API. | [optional] 
**CustOperations** | Pointer to [**[]CustomOperation**](CustomOperation.md) | Custom operations without resource association. | [optional] 

## Methods

### NewVersion

`func NewVersion(apiVersion string, ) *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Version) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Version) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Version) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetExpiry

`func (o *Version) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *Version) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *Version) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *Version) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetResources

`func (o *Version) GetResources() []Resource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Version) GetResourcesOk() (*[]Resource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Version) SetResources(v []Resource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Version) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetCustOperations

`func (o *Version) GetCustOperations() []CustomOperation`

GetCustOperations returns the CustOperations field if non-nil, zero value otherwise.

### GetCustOperationsOk

`func (o *Version) GetCustOperationsOk() (*[]CustomOperation, bool)`

GetCustOperationsOk returns a tuple with the CustOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustOperations

`func (o *Version) SetCustOperations(v []CustomOperation)`

SetCustOperations sets CustOperations field to given value.

### HasCustOperations

`func (o *Version) HasCustOperations() bool`

HasCustOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


