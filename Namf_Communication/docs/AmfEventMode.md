# AmfEventMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trigger** | [**AmfEventTrigger**](AmfEventTrigger.md) |  | 
**MaxReports** | Pointer to **int32** |  | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RepPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**SampRatio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.   | [optional] 
**PartitioningCriteria** | Pointer to [**[]PartitioningCriteria**](PartitioningCriteria.md) |  | [optional] 
**NotifFlag** | Pointer to [**NotificationFlag**](NotificationFlag.md) |  | [optional] 

## Methods

### NewAmfEventMode

`func NewAmfEventMode(trigger AmfEventTrigger, ) *AmfEventMode`

NewAmfEventMode instantiates a new AmfEventMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfEventModeWithDefaults

`func NewAmfEventModeWithDefaults() *AmfEventMode`

NewAmfEventModeWithDefaults instantiates a new AmfEventMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrigger

`func (o *AmfEventMode) GetTrigger() AmfEventTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *AmfEventMode) GetTriggerOk() (*AmfEventTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *AmfEventMode) SetTrigger(v AmfEventTrigger)`

SetTrigger sets Trigger field to given value.


### GetMaxReports

`func (o *AmfEventMode) GetMaxReports() int32`

GetMaxReports returns the MaxReports field if non-nil, zero value otherwise.

### GetMaxReportsOk

`func (o *AmfEventMode) GetMaxReportsOk() (*int32, bool)`

GetMaxReportsOk returns a tuple with the MaxReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReports

`func (o *AmfEventMode) SetMaxReports(v int32)`

SetMaxReports sets MaxReports field to given value.

### HasMaxReports

`func (o *AmfEventMode) HasMaxReports() bool`

HasMaxReports returns a boolean if a field has been set.

### GetExpiry

`func (o *AmfEventMode) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *AmfEventMode) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *AmfEventMode) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *AmfEventMode) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetRepPeriod

`func (o *AmfEventMode) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *AmfEventMode) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *AmfEventMode) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *AmfEventMode) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.

### GetSampRatio

`func (o *AmfEventMode) GetSampRatio() int32`

GetSampRatio returns the SampRatio field if non-nil, zero value otherwise.

### GetSampRatioOk

`func (o *AmfEventMode) GetSampRatioOk() (*int32, bool)`

GetSampRatioOk returns a tuple with the SampRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampRatio

`func (o *AmfEventMode) SetSampRatio(v int32)`

SetSampRatio sets SampRatio field to given value.

### HasSampRatio

`func (o *AmfEventMode) HasSampRatio() bool`

HasSampRatio returns a boolean if a field has been set.

### GetPartitioningCriteria

`func (o *AmfEventMode) GetPartitioningCriteria() []PartitioningCriteria`

GetPartitioningCriteria returns the PartitioningCriteria field if non-nil, zero value otherwise.

### GetPartitioningCriteriaOk

`func (o *AmfEventMode) GetPartitioningCriteriaOk() (*[]PartitioningCriteria, bool)`

GetPartitioningCriteriaOk returns a tuple with the PartitioningCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioningCriteria

`func (o *AmfEventMode) SetPartitioningCriteria(v []PartitioningCriteria)`

SetPartitioningCriteria sets PartitioningCriteria field to given value.

### HasPartitioningCriteria

`func (o *AmfEventMode) HasPartitioningCriteria() bool`

HasPartitioningCriteria returns a boolean if a field has been set.

### GetNotifFlag

`func (o *AmfEventMode) GetNotifFlag() NotificationFlag`

GetNotifFlag returns the NotifFlag field if non-nil, zero value otherwise.

### GetNotifFlagOk

`func (o *AmfEventMode) GetNotifFlagOk() (*NotificationFlag, bool)`

GetNotifFlagOk returns a tuple with the NotifFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifFlag

`func (o *AmfEventMode) SetNotifFlag(v NotificationFlag)`

SetNotifFlag sets NotifFlag field to given value.

### HasNotifFlag

`func (o *AmfEventMode) HasNotifFlag() bool`

HasNotifFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


