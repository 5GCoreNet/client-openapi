# NnwdafDataManagementNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataNotification** | Pointer to [**DataNotification**](DataNotification.md) |  | [optional] 
**NotifCorrId** | **string** | Notification correlation identifier. | 
**TerminationReq** | Pointer to **string** | It indicates the termination of the data management subscription that requested by the NWDAF.  | [optional] 
**FetchInstruct** | Pointer to [**FetchInstruction**](FetchInstruction.md) |  | [optional] 
**NotifTimestamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 

## Methods

### NewNnwdafDataManagementNotif

`func NewNnwdafDataManagementNotif(notifCorrId string, notifTimestamp time.Time, ) *NnwdafDataManagementNotif`

NewNnwdafDataManagementNotif instantiates a new NnwdafDataManagementNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNnwdafDataManagementNotifWithDefaults

`func NewNnwdafDataManagementNotifWithDefaults() *NnwdafDataManagementNotif`

NewNnwdafDataManagementNotifWithDefaults instantiates a new NnwdafDataManagementNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataNotification

`func (o *NnwdafDataManagementNotif) GetDataNotification() DataNotification`

GetDataNotification returns the DataNotification field if non-nil, zero value otherwise.

### GetDataNotificationOk

`func (o *NnwdafDataManagementNotif) GetDataNotificationOk() (*DataNotification, bool)`

GetDataNotificationOk returns a tuple with the DataNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataNotification

`func (o *NnwdafDataManagementNotif) SetDataNotification(v DataNotification)`

SetDataNotification sets DataNotification field to given value.

### HasDataNotification

`func (o *NnwdafDataManagementNotif) HasDataNotification() bool`

HasDataNotification returns a boolean if a field has been set.

### GetNotifCorrId

`func (o *NnwdafDataManagementNotif) GetNotifCorrId() string`

GetNotifCorrId returns the NotifCorrId field if non-nil, zero value otherwise.

### GetNotifCorrIdOk

`func (o *NnwdafDataManagementNotif) GetNotifCorrIdOk() (*string, bool)`

GetNotifCorrIdOk returns a tuple with the NotifCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorrId

`func (o *NnwdafDataManagementNotif) SetNotifCorrId(v string)`

SetNotifCorrId sets NotifCorrId field to given value.


### GetTerminationReq

`func (o *NnwdafDataManagementNotif) GetTerminationReq() string`

GetTerminationReq returns the TerminationReq field if non-nil, zero value otherwise.

### GetTerminationReqOk

`func (o *NnwdafDataManagementNotif) GetTerminationReqOk() (*string, bool)`

GetTerminationReqOk returns a tuple with the TerminationReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationReq

`func (o *NnwdafDataManagementNotif) SetTerminationReq(v string)`

SetTerminationReq sets TerminationReq field to given value.

### HasTerminationReq

`func (o *NnwdafDataManagementNotif) HasTerminationReq() bool`

HasTerminationReq returns a boolean if a field has been set.

### GetFetchInstruct

`func (o *NnwdafDataManagementNotif) GetFetchInstruct() FetchInstruction`

GetFetchInstruct returns the FetchInstruct field if non-nil, zero value otherwise.

### GetFetchInstructOk

`func (o *NnwdafDataManagementNotif) GetFetchInstructOk() (*FetchInstruction, bool)`

GetFetchInstructOk returns a tuple with the FetchInstruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchInstruct

`func (o *NnwdafDataManagementNotif) SetFetchInstruct(v FetchInstruction)`

SetFetchInstruct sets FetchInstruct field to given value.

### HasFetchInstruct

`func (o *NnwdafDataManagementNotif) HasFetchInstruct() bool`

HasFetchInstruct returns a boolean if a field has been set.

### GetNotifTimestamp

`func (o *NnwdafDataManagementNotif) GetNotifTimestamp() time.Time`

GetNotifTimestamp returns the NotifTimestamp field if non-nil, zero value otherwise.

### GetNotifTimestampOk

`func (o *NnwdafDataManagementNotif) GetNotifTimestampOk() (*time.Time, bool)`

GetNotifTimestampOk returns a tuple with the NotifTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifTimestamp

`func (o *NnwdafDataManagementNotif) SetNotifTimestamp(v time.Time)`

SetNotifTimestamp sets NotifTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


