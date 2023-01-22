# ACRCompleteEventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcrRes** | **bool** | Indicates whether the ACR is successful or failure. | 
**TEasEndpoint** | [**EndPoint**](EndPoint.md) |  | 

## Methods

### NewACRCompleteEventInfo

`func NewACRCompleteEventInfo(acrRes bool, tEasEndpoint EndPoint, ) *ACRCompleteEventInfo`

NewACRCompleteEventInfo instantiates a new ACRCompleteEventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACRCompleteEventInfoWithDefaults

`func NewACRCompleteEventInfoWithDefaults() *ACRCompleteEventInfo`

NewACRCompleteEventInfoWithDefaults instantiates a new ACRCompleteEventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcrRes

`func (o *ACRCompleteEventInfo) GetAcrRes() bool`

GetAcrRes returns the AcrRes field if non-nil, zero value otherwise.

### GetAcrResOk

`func (o *ACRCompleteEventInfo) GetAcrResOk() (*bool, bool)`

GetAcrResOk returns a tuple with the AcrRes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrRes

`func (o *ACRCompleteEventInfo) SetAcrRes(v bool)`

SetAcrRes sets AcrRes field to given value.


### GetTEasEndpoint

`func (o *ACRCompleteEventInfo) GetTEasEndpoint() EndPoint`

GetTEasEndpoint returns the TEasEndpoint field if non-nil, zero value otherwise.

### GetTEasEndpointOk

`func (o *ACRCompleteEventInfo) GetTEasEndpointOk() (*EndPoint, bool)`

GetTEasEndpointOk returns a tuple with the TEasEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEasEndpoint

`func (o *ACRCompleteEventInfo) SetTEasEndpoint(v EndPoint)`

SetTEasEndpoint sets TEasEndpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


