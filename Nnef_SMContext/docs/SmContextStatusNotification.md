# SmContextStatusNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**SmContextStatus**](SmContextStatus.md) |  | 
**SmContextId** | **string** | String providing an URI formatted according to RFC 3986. | 
**Cause** | Pointer to [**ReleaseCause**](ReleaseCause.md) |  | [optional] 
**SmallDataRateStatus** | Pointer to [**SmallDataRateStatus**](SmallDataRateStatus.md) |  | [optional] 
**ApnRateStatus** | Pointer to [**ApnRateStatus**](ApnRateStatus.md) |  | [optional] 

## Methods

### NewSmContextStatusNotification

`func NewSmContextStatusNotification(status SmContextStatus, smContextId string, ) *SmContextStatusNotification`

NewSmContextStatusNotification instantiates a new SmContextStatusNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextStatusNotificationWithDefaults

`func NewSmContextStatusNotificationWithDefaults() *SmContextStatusNotification`

NewSmContextStatusNotificationWithDefaults instantiates a new SmContextStatusNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SmContextStatusNotification) GetStatus() SmContextStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmContextStatusNotification) GetStatusOk() (*SmContextStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmContextStatusNotification) SetStatus(v SmContextStatus)`

SetStatus sets Status field to given value.


### GetSmContextId

`func (o *SmContextStatusNotification) GetSmContextId() string`

GetSmContextId returns the SmContextId field if non-nil, zero value otherwise.

### GetSmContextIdOk

`func (o *SmContextStatusNotification) GetSmContextIdOk() (*string, bool)`

GetSmContextIdOk returns a tuple with the SmContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextId

`func (o *SmContextStatusNotification) SetSmContextId(v string)`

SetSmContextId sets SmContextId field to given value.


### GetCause

`func (o *SmContextStatusNotification) GetCause() ReleaseCause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *SmContextStatusNotification) GetCauseOk() (*ReleaseCause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *SmContextStatusNotification) SetCause(v ReleaseCause)`

SetCause sets Cause field to given value.

### HasCause

`func (o *SmContextStatusNotification) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetSmallDataRateStatus

`func (o *SmContextStatusNotification) GetSmallDataRateStatus() SmallDataRateStatus`

GetSmallDataRateStatus returns the SmallDataRateStatus field if non-nil, zero value otherwise.

### GetSmallDataRateStatusOk

`func (o *SmContextStatusNotification) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool)`

GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallDataRateStatus

`func (o *SmContextStatusNotification) SetSmallDataRateStatus(v SmallDataRateStatus)`

SetSmallDataRateStatus sets SmallDataRateStatus field to given value.

### HasSmallDataRateStatus

`func (o *SmContextStatusNotification) HasSmallDataRateStatus() bool`

HasSmallDataRateStatus returns a boolean if a field has been set.

### GetApnRateStatus

`func (o *SmContextStatusNotification) GetApnRateStatus() ApnRateStatus`

GetApnRateStatus returns the ApnRateStatus field if non-nil, zero value otherwise.

### GetApnRateStatusOk

`func (o *SmContextStatusNotification) GetApnRateStatusOk() (*ApnRateStatus, bool)`

GetApnRateStatusOk returns a tuple with the ApnRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnRateStatus

`func (o *SmContextStatusNotification) SetApnRateStatus(v ApnRateStatus)`

SetApnRateStatus sets ApnRateStatus field to given value.

### HasApnRateStatus

`func (o *SmContextStatusNotification) HasApnRateStatus() bool`

HasApnRateStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


