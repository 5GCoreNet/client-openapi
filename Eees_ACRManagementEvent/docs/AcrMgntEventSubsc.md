# AcrMgntEventSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**AcrMgntEvent**](AcrMgntEvent.md) |  | 
**EventFilter** | Pointer to [**AcrMgntEventFilter**](AcrMgntEventFilter.md) |  | [optional] 
**EvtReq** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**TgtUeId** | Pointer to [**TargetUeIdentification**](TargetUeIdentification.md) |  | [optional] 
**DnaiChgType** | Pointer to [**DnaiChangeType**](DnaiChangeType.md) |  | [optional] 
**EasAckInd** | Pointer to **bool** |  | [optional] 
**EasChars** | Pointer to [**[]EasCharacteristics**](EasCharacteristics.md) | A list of EAS characteristics. | [optional] 

## Methods

### NewAcrMgntEventSubsc

`func NewAcrMgntEventSubsc(event AcrMgntEvent, ) *AcrMgntEventSubsc`

NewAcrMgntEventSubsc instantiates a new AcrMgntEventSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcrMgntEventSubscWithDefaults

`func NewAcrMgntEventSubscWithDefaults() *AcrMgntEventSubsc`

NewAcrMgntEventSubscWithDefaults instantiates a new AcrMgntEventSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *AcrMgntEventSubsc) GetEvent() AcrMgntEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AcrMgntEventSubsc) GetEventOk() (*AcrMgntEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AcrMgntEventSubsc) SetEvent(v AcrMgntEvent)`

SetEvent sets Event field to given value.


### GetEventFilter

`func (o *AcrMgntEventSubsc) GetEventFilter() AcrMgntEventFilter`

GetEventFilter returns the EventFilter field if non-nil, zero value otherwise.

### GetEventFilterOk

`func (o *AcrMgntEventSubsc) GetEventFilterOk() (*AcrMgntEventFilter, bool)`

GetEventFilterOk returns a tuple with the EventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilter

`func (o *AcrMgntEventSubsc) SetEventFilter(v AcrMgntEventFilter)`

SetEventFilter sets EventFilter field to given value.

### HasEventFilter

`func (o *AcrMgntEventSubsc) HasEventFilter() bool`

HasEventFilter returns a boolean if a field has been set.

### GetEvtReq

`func (o *AcrMgntEventSubsc) GetEvtReq() ReportingInformation`

GetEvtReq returns the EvtReq field if non-nil, zero value otherwise.

### GetEvtReqOk

`func (o *AcrMgntEventSubsc) GetEvtReqOk() (*ReportingInformation, bool)`

GetEvtReqOk returns a tuple with the EvtReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvtReq

`func (o *AcrMgntEventSubsc) SetEvtReq(v ReportingInformation)`

SetEvtReq sets EvtReq field to given value.

### HasEvtReq

`func (o *AcrMgntEventSubsc) HasEvtReq() bool`

HasEvtReq returns a boolean if a field has been set.

### GetTgtUeId

`func (o *AcrMgntEventSubsc) GetTgtUeId() TargetUeIdentification`

GetTgtUeId returns the TgtUeId field if non-nil, zero value otherwise.

### GetTgtUeIdOk

`func (o *AcrMgntEventSubsc) GetTgtUeIdOk() (*TargetUeIdentification, bool)`

GetTgtUeIdOk returns a tuple with the TgtUeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUeId

`func (o *AcrMgntEventSubsc) SetTgtUeId(v TargetUeIdentification)`

SetTgtUeId sets TgtUeId field to given value.

### HasTgtUeId

`func (o *AcrMgntEventSubsc) HasTgtUeId() bool`

HasTgtUeId returns a boolean if a field has been set.

### GetDnaiChgType

`func (o *AcrMgntEventSubsc) GetDnaiChgType() DnaiChangeType`

GetDnaiChgType returns the DnaiChgType field if non-nil, zero value otherwise.

### GetDnaiChgTypeOk

`func (o *AcrMgntEventSubsc) GetDnaiChgTypeOk() (*DnaiChangeType, bool)`

GetDnaiChgTypeOk returns a tuple with the DnaiChgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiChgType

`func (o *AcrMgntEventSubsc) SetDnaiChgType(v DnaiChangeType)`

SetDnaiChgType sets DnaiChgType field to given value.

### HasDnaiChgType

`func (o *AcrMgntEventSubsc) HasDnaiChgType() bool`

HasDnaiChgType returns a boolean if a field has been set.

### GetEasAckInd

`func (o *AcrMgntEventSubsc) GetEasAckInd() bool`

GetEasAckInd returns the EasAckInd field if non-nil, zero value otherwise.

### GetEasAckIndOk

`func (o *AcrMgntEventSubsc) GetEasAckIndOk() (*bool, bool)`

GetEasAckIndOk returns a tuple with the EasAckInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasAckInd

`func (o *AcrMgntEventSubsc) SetEasAckInd(v bool)`

SetEasAckInd sets EasAckInd field to given value.

### HasEasAckInd

`func (o *AcrMgntEventSubsc) HasEasAckInd() bool`

HasEasAckInd returns a boolean if a field has been set.

### GetEasChars

`func (o *AcrMgntEventSubsc) GetEasChars() []EasCharacteristics`

GetEasChars returns the EasChars field if non-nil, zero value otherwise.

### GetEasCharsOk

`func (o *AcrMgntEventSubsc) GetEasCharsOk() (*[]EasCharacteristics, bool)`

GetEasCharsOk returns a tuple with the EasChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasChars

`func (o *AcrMgntEventSubsc) SetEasChars(v []EasCharacteristics)`

SetEasChars sets EasChars field to given value.

### HasEasChars

`func (o *AcrMgntEventSubsc) HasEasChars() bool`

HasEasChars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


