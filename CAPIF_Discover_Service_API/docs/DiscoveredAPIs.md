# DiscoveredAPIs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAPIDescriptions** | Pointer to [**[]ServiceAPIDescription**](ServiceAPIDescription.md) | Description of the service API as published by the service. Each service API description shall include AEF profiles matching the filter criteria.  | [optional] 

## Methods

### NewDiscoveredAPIs

`func NewDiscoveredAPIs() *DiscoveredAPIs`

NewDiscoveredAPIs instantiates a new DiscoveredAPIs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredAPIsWithDefaults

`func NewDiscoveredAPIsWithDefaults() *DiscoveredAPIs`

NewDiscoveredAPIsWithDefaults instantiates a new DiscoveredAPIs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAPIDescriptions

`func (o *DiscoveredAPIs) GetServiceAPIDescriptions() []ServiceAPIDescription`

GetServiceAPIDescriptions returns the ServiceAPIDescriptions field if non-nil, zero value otherwise.

### GetServiceAPIDescriptionsOk

`func (o *DiscoveredAPIs) GetServiceAPIDescriptionsOk() (*[]ServiceAPIDescription, bool)`

GetServiceAPIDescriptionsOk returns a tuple with the ServiceAPIDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAPIDescriptions

`func (o *DiscoveredAPIs) SetServiceAPIDescriptions(v []ServiceAPIDescription)`

SetServiceAPIDescriptions sets ServiceAPIDescriptions field to given value.

### HasServiceAPIDescriptions

`func (o *DiscoveredAPIs) HasServiceAPIDescriptions() bool`

HasServiceAPIDescriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


