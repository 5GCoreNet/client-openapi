# AmEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**AmEvent**](AmEvent.md) |  | 
**ImmRep** | Pointer to **bool** |  | [optional] 
**NotifMethod** | Pointer to [**NotificationMethod**](NotificationMethod.md) |  | [optional] 
**MaxReportNbr** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**MonDur** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RepPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 

## Methods

### NewAmEventData

`func NewAmEventData(event AmEvent, ) *AmEventData`

NewAmEventData instantiates a new AmEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmEventDataWithDefaults

`func NewAmEventDataWithDefaults() *AmEventData`

NewAmEventDataWithDefaults instantiates a new AmEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *AmEventData) GetEvent() AmEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AmEventData) GetEventOk() (*AmEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AmEventData) SetEvent(v AmEvent)`

SetEvent sets Event field to given value.


### GetImmRep

`func (o *AmEventData) GetImmRep() bool`

GetImmRep returns the ImmRep field if non-nil, zero value otherwise.

### GetImmRepOk

`func (o *AmEventData) GetImmRepOk() (*bool, bool)`

GetImmRepOk returns a tuple with the ImmRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmRep

`func (o *AmEventData) SetImmRep(v bool)`

SetImmRep sets ImmRep field to given value.

### HasImmRep

`func (o *AmEventData) HasImmRep() bool`

HasImmRep returns a boolean if a field has been set.

### GetNotifMethod

`func (o *AmEventData) GetNotifMethod() NotificationMethod`

GetNotifMethod returns the NotifMethod field if non-nil, zero value otherwise.

### GetNotifMethodOk

`func (o *AmEventData) GetNotifMethodOk() (*NotificationMethod, bool)`

GetNotifMethodOk returns a tuple with the NotifMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifMethod

`func (o *AmEventData) SetNotifMethod(v NotificationMethod)`

SetNotifMethod sets NotifMethod field to given value.

### HasNotifMethod

`func (o *AmEventData) HasNotifMethod() bool`

HasNotifMethod returns a boolean if a field has been set.

### GetMaxReportNbr

`func (o *AmEventData) GetMaxReportNbr() int32`

GetMaxReportNbr returns the MaxReportNbr field if non-nil, zero value otherwise.

### GetMaxReportNbrOk

`func (o *AmEventData) GetMaxReportNbrOk() (*int32, bool)`

GetMaxReportNbrOk returns a tuple with the MaxReportNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReportNbr

`func (o *AmEventData) SetMaxReportNbr(v int32)`

SetMaxReportNbr sets MaxReportNbr field to given value.

### HasMaxReportNbr

`func (o *AmEventData) HasMaxReportNbr() bool`

HasMaxReportNbr returns a boolean if a field has been set.

### GetMonDur

`func (o *AmEventData) GetMonDur() time.Time`

GetMonDur returns the MonDur field if non-nil, zero value otherwise.

### GetMonDurOk

`func (o *AmEventData) GetMonDurOk() (*time.Time, bool)`

GetMonDurOk returns a tuple with the MonDur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonDur

`func (o *AmEventData) SetMonDur(v time.Time)`

SetMonDur sets MonDur field to given value.

### HasMonDur

`func (o *AmEventData) HasMonDur() bool`

HasMonDur returns a boolean if a field has been set.

### GetRepPeriod

`func (o *AmEventData) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *AmEventData) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *AmEventData) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *AmEventData) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


