# NotifCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitoredAttributes** | Pointer to **[]string** |  | [optional] 
**UnmonitoredAttributes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNotifCondition

`func NewNotifCondition() *NotifCondition`

NewNotifCondition instantiates a new NotifCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifConditionWithDefaults

`func NewNotifConditionWithDefaults() *NotifCondition`

NewNotifConditionWithDefaults instantiates a new NotifCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitoredAttributes

`func (o *NotifCondition) GetMonitoredAttributes() []string`

GetMonitoredAttributes returns the MonitoredAttributes field if non-nil, zero value otherwise.

### GetMonitoredAttributesOk

`func (o *NotifCondition) GetMonitoredAttributesOk() (*[]string, bool)`

GetMonitoredAttributesOk returns a tuple with the MonitoredAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredAttributes

`func (o *NotifCondition) SetMonitoredAttributes(v []string)`

SetMonitoredAttributes sets MonitoredAttributes field to given value.

### HasMonitoredAttributes

`func (o *NotifCondition) HasMonitoredAttributes() bool`

HasMonitoredAttributes returns a boolean if a field has been set.

### GetUnmonitoredAttributes

`func (o *NotifCondition) GetUnmonitoredAttributes() []string`

GetUnmonitoredAttributes returns the UnmonitoredAttributes field if non-nil, zero value otherwise.

### GetUnmonitoredAttributesOk

`func (o *NotifCondition) GetUnmonitoredAttributesOk() (*[]string, bool)`

GetUnmonitoredAttributesOk returns a tuple with the UnmonitoredAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmonitoredAttributes

`func (o *NotifCondition) SetUnmonitoredAttributes(v []string)`

SetUnmonitoredAttributes sets UnmonitoredAttributes field to given value.

### HasUnmonitoredAttributes

`func (o *NotifCondition) HasUnmonitoredAttributes() bool`

HasUnmonitoredAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


