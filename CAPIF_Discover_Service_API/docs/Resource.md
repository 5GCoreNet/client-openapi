# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | Resource name | 
**CommType** | [**CommunicationType**](CommunicationType.md) |  | 
**Uri** | **string** | Relative URI of the API resource, it is set as {apiSpecificSuffixes} part of the URI structure as defined in clause 5.2.4 of 3GPP TS 29.122.  | 
**CustOpName** | Pointer to **string** | it is set as {custOpName} part of the URI structure for a custom operation associated with a resource as defined in clause 5.2.4 of 3GPP TS 29.122.  | [optional] 
**Operations** | Pointer to [**[]Operation**](Operation.md) | Supported HTTP methods for the API resource. Only applicable when the protocol in AefProfile indicates HTTP.  | [optional] 
**Description** | Pointer to **string** | Text description of the API resource | [optional] 

## Methods

### NewResource

`func NewResource(resourceName string, commType CommunicationType, uri string, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *Resource) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *Resource) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *Resource) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetCommType

`func (o *Resource) GetCommType() CommunicationType`

GetCommType returns the CommType field if non-nil, zero value otherwise.

### GetCommTypeOk

`func (o *Resource) GetCommTypeOk() (*CommunicationType, bool)`

GetCommTypeOk returns a tuple with the CommType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommType

`func (o *Resource) SetCommType(v CommunicationType)`

SetCommType sets CommType field to given value.


### GetUri

`func (o *Resource) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Resource) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Resource) SetUri(v string)`

SetUri sets Uri field to given value.


### GetCustOpName

`func (o *Resource) GetCustOpName() string`

GetCustOpName returns the CustOpName field if non-nil, zero value otherwise.

### GetCustOpNameOk

`func (o *Resource) GetCustOpNameOk() (*string, bool)`

GetCustOpNameOk returns a tuple with the CustOpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustOpName

`func (o *Resource) SetCustOpName(v string)`

SetCustOpName sets CustOpName field to given value.

### HasCustOpName

`func (o *Resource) HasCustOpName() bool`

HasCustOpName returns a boolean if a field has been set.

### GetOperations

`func (o *Resource) GetOperations() []Operation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Resource) GetOperationsOk() (*[]Operation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *Resource) SetOperations(v []Operation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *Resource) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetDescription

`func (o *Resource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Resource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Resource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Resource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


