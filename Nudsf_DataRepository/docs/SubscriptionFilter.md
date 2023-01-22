# SubscriptionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitoredResourceUris** | Pointer to **[]string** | list of resources applicable to the subscription | [optional] 
**Operations** | Pointer to [**[]RecordOperation**](RecordOperation.md) | list of resources applicable to the subscription | [optional] 

## Methods

### NewSubscriptionFilter

`func NewSubscriptionFilter() *SubscriptionFilter`

NewSubscriptionFilter instantiates a new SubscriptionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionFilterWithDefaults

`func NewSubscriptionFilterWithDefaults() *SubscriptionFilter`

NewSubscriptionFilterWithDefaults instantiates a new SubscriptionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitoredResourceUris

`func (o *SubscriptionFilter) GetMonitoredResourceUris() []string`

GetMonitoredResourceUris returns the MonitoredResourceUris field if non-nil, zero value otherwise.

### GetMonitoredResourceUrisOk

`func (o *SubscriptionFilter) GetMonitoredResourceUrisOk() (*[]string, bool)`

GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredResourceUris

`func (o *SubscriptionFilter) SetMonitoredResourceUris(v []string)`

SetMonitoredResourceUris sets MonitoredResourceUris field to given value.

### HasMonitoredResourceUris

`func (o *SubscriptionFilter) HasMonitoredResourceUris() bool`

HasMonitoredResourceUris returns a boolean if a field has been set.

### GetOperations

`func (o *SubscriptionFilter) GetOperations() []RecordOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *SubscriptionFilter) GetOperationsOk() (*[]RecordOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *SubscriptionFilter) SetOperations(v []RecordOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *SubscriptionFilter) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


