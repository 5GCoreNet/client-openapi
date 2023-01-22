# ACRInfoNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubId** | **string** | String identifying the Individual ACR events subscription for which the ACT Information notification is delivered.  | 
**EasId** | **string** | Application identifier of the EAS. | 
**AcId** | Pointer to **string** | Identity of the AC. | [optional] 
**EventId** | [**ACREventIDs**](ACREventIDs.md) |  | 
**TrgtInfo** | Pointer to [**TargetInfo**](TargetInfo.md) |  | [optional] 
**AcrStatus** | Pointer to [**ACRCompleteEventInfo**](ACRCompleteEventInfo.md) |  | [optional] 
**FailReason** | Pointer to **string** | Indicates the cause information for the failure. | [optional] 
**EecCtxtReloc** | Pointer to [**EecCtxtRelocStatus**](EecCtxtRelocStatus.md) |  | [optional] 

## Methods

### NewACRInfoNotification

`func NewACRInfoNotification(subId string, easId string, eventId ACREventIDs, ) *ACRInfoNotification`

NewACRInfoNotification instantiates a new ACRInfoNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACRInfoNotificationWithDefaults

`func NewACRInfoNotificationWithDefaults() *ACRInfoNotification`

NewACRInfoNotificationWithDefaults instantiates a new ACRInfoNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubId

`func (o *ACRInfoNotification) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *ACRInfoNotification) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *ACRInfoNotification) SetSubId(v string)`

SetSubId sets SubId field to given value.


### GetEasId

`func (o *ACRInfoNotification) GetEasId() string`

GetEasId returns the EasId field if non-nil, zero value otherwise.

### GetEasIdOk

`func (o *ACRInfoNotification) GetEasIdOk() (*string, bool)`

GetEasIdOk returns a tuple with the EasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasId

`func (o *ACRInfoNotification) SetEasId(v string)`

SetEasId sets EasId field to given value.


### GetAcId

`func (o *ACRInfoNotification) GetAcId() string`

GetAcId returns the AcId field if non-nil, zero value otherwise.

### GetAcIdOk

`func (o *ACRInfoNotification) GetAcIdOk() (*string, bool)`

GetAcIdOk returns a tuple with the AcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcId

`func (o *ACRInfoNotification) SetAcId(v string)`

SetAcId sets AcId field to given value.

### HasAcId

`func (o *ACRInfoNotification) HasAcId() bool`

HasAcId returns a boolean if a field has been set.

### GetEventId

`func (o *ACRInfoNotification) GetEventId() ACREventIDs`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ACRInfoNotification) GetEventIdOk() (*ACREventIDs, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ACRInfoNotification) SetEventId(v ACREventIDs)`

SetEventId sets EventId field to given value.


### GetTrgtInfo

`func (o *ACRInfoNotification) GetTrgtInfo() TargetInfo`

GetTrgtInfo returns the TrgtInfo field if non-nil, zero value otherwise.

### GetTrgtInfoOk

`func (o *ACRInfoNotification) GetTrgtInfoOk() (*TargetInfo, bool)`

GetTrgtInfoOk returns a tuple with the TrgtInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrgtInfo

`func (o *ACRInfoNotification) SetTrgtInfo(v TargetInfo)`

SetTrgtInfo sets TrgtInfo field to given value.

### HasTrgtInfo

`func (o *ACRInfoNotification) HasTrgtInfo() bool`

HasTrgtInfo returns a boolean if a field has been set.

### GetAcrStatus

`func (o *ACRInfoNotification) GetAcrStatus() ACRCompleteEventInfo`

GetAcrStatus returns the AcrStatus field if non-nil, zero value otherwise.

### GetAcrStatusOk

`func (o *ACRInfoNotification) GetAcrStatusOk() (*ACRCompleteEventInfo, bool)`

GetAcrStatusOk returns a tuple with the AcrStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrStatus

`func (o *ACRInfoNotification) SetAcrStatus(v ACRCompleteEventInfo)`

SetAcrStatus sets AcrStatus field to given value.

### HasAcrStatus

`func (o *ACRInfoNotification) HasAcrStatus() bool`

HasAcrStatus returns a boolean if a field has been set.

### GetFailReason

`func (o *ACRInfoNotification) GetFailReason() string`

GetFailReason returns the FailReason field if non-nil, zero value otherwise.

### GetFailReasonOk

`func (o *ACRInfoNotification) GetFailReasonOk() (*string, bool)`

GetFailReasonOk returns a tuple with the FailReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailReason

`func (o *ACRInfoNotification) SetFailReason(v string)`

SetFailReason sets FailReason field to given value.

### HasFailReason

`func (o *ACRInfoNotification) HasFailReason() bool`

HasFailReason returns a boolean if a field has been set.

### GetEecCtxtReloc

`func (o *ACRInfoNotification) GetEecCtxtReloc() EecCtxtRelocStatus`

GetEecCtxtReloc returns the EecCtxtReloc field if non-nil, zero value otherwise.

### GetEecCtxtRelocOk

`func (o *ACRInfoNotification) GetEecCtxtRelocOk() (*EecCtxtRelocStatus, bool)`

GetEecCtxtRelocOk returns a tuple with the EecCtxtReloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecCtxtReloc

`func (o *ACRInfoNotification) SetEecCtxtReloc(v EecCtxtRelocStatus)`

SetEecCtxtReloc sets EecCtxtReloc field to given value.

### HasEecCtxtReloc

`func (o *ACRInfoNotification) HasEecCtxtReloc() bool`

HasEecCtxtReloc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


