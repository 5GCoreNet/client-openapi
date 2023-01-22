# SEALEventDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | [**SEALEvent**](SEALEvent.md) |  | 
**LmInfos** | Pointer to [**[]LMInformation**](LMInformation.md) |  | [optional] 
**ValGroupDocuments** | Pointer to [**[]VALGroupDocument**](VALGroupDocument.md) | The VAL groups documents with modified membership and configuration information.  | [optional] 
**ProfileDocs** | Pointer to [**[]ProfileDoc**](ProfileDoc.md) | Updated profile information associated with VAL Users or VAL UEs. | [optional] 
**MsgFltrs** | Pointer to [**[]MessageFilter**](MessageFilter.md) | The message filter information for various member VAL User or UEs of the VAL group.  | [optional] 
**MonRep** | Pointer to [**[]MonitorEventsReport**](MonitorEventsReport.md) | The events reports with details of the events related to the VAL UE(s). | [optional] 
**LocAdhr** | Pointer to [**[]LocationDevMonReport**](LocationDevMonReport.md) | The location deviation information for the interested VAL User ID or UE IDs in a given location.  | [optional] 
**TempGroupInfo** | Pointer to [**TempGroupInfo**](TempGroupInfo.md) |  | [optional] 
**LocAreaMonRep** | Pointer to [**[]LocationAreaMonReport**](LocationAreaMonReport.md) | The location area monitoring of the given area of interest. | [optional] 

## Methods

### NewSEALEventDetail

`func NewSEALEventDetail(eventId SEALEvent, ) *SEALEventDetail`

NewSEALEventDetail instantiates a new SEALEventDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSEALEventDetailWithDefaults

`func NewSEALEventDetailWithDefaults() *SEALEventDetail`

NewSEALEventDetailWithDefaults instantiates a new SEALEventDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *SEALEventDetail) GetEventId() SEALEvent`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SEALEventDetail) GetEventIdOk() (*SEALEvent, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SEALEventDetail) SetEventId(v SEALEvent)`

SetEventId sets EventId field to given value.


### GetLmInfos

`func (o *SEALEventDetail) GetLmInfos() []LMInformation`

GetLmInfos returns the LmInfos field if non-nil, zero value otherwise.

### GetLmInfosOk

`func (o *SEALEventDetail) GetLmInfosOk() (*[]LMInformation, bool)`

GetLmInfosOk returns a tuple with the LmInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmInfos

`func (o *SEALEventDetail) SetLmInfos(v []LMInformation)`

SetLmInfos sets LmInfos field to given value.

### HasLmInfos

`func (o *SEALEventDetail) HasLmInfos() bool`

HasLmInfos returns a boolean if a field has been set.

### GetValGroupDocuments

`func (o *SEALEventDetail) GetValGroupDocuments() []VALGroupDocument`

GetValGroupDocuments returns the ValGroupDocuments field if non-nil, zero value otherwise.

### GetValGroupDocumentsOk

`func (o *SEALEventDetail) GetValGroupDocumentsOk() (*[]VALGroupDocument, bool)`

GetValGroupDocumentsOk returns a tuple with the ValGroupDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValGroupDocuments

`func (o *SEALEventDetail) SetValGroupDocuments(v []VALGroupDocument)`

SetValGroupDocuments sets ValGroupDocuments field to given value.

### HasValGroupDocuments

`func (o *SEALEventDetail) HasValGroupDocuments() bool`

HasValGroupDocuments returns a boolean if a field has been set.

### GetProfileDocs

`func (o *SEALEventDetail) GetProfileDocs() []ProfileDoc`

GetProfileDocs returns the ProfileDocs field if non-nil, zero value otherwise.

### GetProfileDocsOk

`func (o *SEALEventDetail) GetProfileDocsOk() (*[]ProfileDoc, bool)`

GetProfileDocsOk returns a tuple with the ProfileDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileDocs

`func (o *SEALEventDetail) SetProfileDocs(v []ProfileDoc)`

SetProfileDocs sets ProfileDocs field to given value.

### HasProfileDocs

`func (o *SEALEventDetail) HasProfileDocs() bool`

HasProfileDocs returns a boolean if a field has been set.

### GetMsgFltrs

`func (o *SEALEventDetail) GetMsgFltrs() []MessageFilter`

GetMsgFltrs returns the MsgFltrs field if non-nil, zero value otherwise.

### GetMsgFltrsOk

`func (o *SEALEventDetail) GetMsgFltrsOk() (*[]MessageFilter, bool)`

GetMsgFltrsOk returns a tuple with the MsgFltrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgFltrs

`func (o *SEALEventDetail) SetMsgFltrs(v []MessageFilter)`

SetMsgFltrs sets MsgFltrs field to given value.

### HasMsgFltrs

`func (o *SEALEventDetail) HasMsgFltrs() bool`

HasMsgFltrs returns a boolean if a field has been set.

### GetMonRep

`func (o *SEALEventDetail) GetMonRep() []MonitorEventsReport`

GetMonRep returns the MonRep field if non-nil, zero value otherwise.

### GetMonRepOk

`func (o *SEALEventDetail) GetMonRepOk() (*[]MonitorEventsReport, bool)`

GetMonRepOk returns a tuple with the MonRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonRep

`func (o *SEALEventDetail) SetMonRep(v []MonitorEventsReport)`

SetMonRep sets MonRep field to given value.

### HasMonRep

`func (o *SEALEventDetail) HasMonRep() bool`

HasMonRep returns a boolean if a field has been set.

### GetLocAdhr

`func (o *SEALEventDetail) GetLocAdhr() []LocationDevMonReport`

GetLocAdhr returns the LocAdhr field if non-nil, zero value otherwise.

### GetLocAdhrOk

`func (o *SEALEventDetail) GetLocAdhrOk() (*[]LocationDevMonReport, bool)`

GetLocAdhrOk returns a tuple with the LocAdhr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocAdhr

`func (o *SEALEventDetail) SetLocAdhr(v []LocationDevMonReport)`

SetLocAdhr sets LocAdhr field to given value.

### HasLocAdhr

`func (o *SEALEventDetail) HasLocAdhr() bool`

HasLocAdhr returns a boolean if a field has been set.

### GetTempGroupInfo

`func (o *SEALEventDetail) GetTempGroupInfo() TempGroupInfo`

GetTempGroupInfo returns the TempGroupInfo field if non-nil, zero value otherwise.

### GetTempGroupInfoOk

`func (o *SEALEventDetail) GetTempGroupInfoOk() (*TempGroupInfo, bool)`

GetTempGroupInfoOk returns a tuple with the TempGroupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempGroupInfo

`func (o *SEALEventDetail) SetTempGroupInfo(v TempGroupInfo)`

SetTempGroupInfo sets TempGroupInfo field to given value.

### HasTempGroupInfo

`func (o *SEALEventDetail) HasTempGroupInfo() bool`

HasTempGroupInfo returns a boolean if a field has been set.

### GetLocAreaMonRep

`func (o *SEALEventDetail) GetLocAreaMonRep() []LocationAreaMonReport`

GetLocAreaMonRep returns the LocAreaMonRep field if non-nil, zero value otherwise.

### GetLocAreaMonRepOk

`func (o *SEALEventDetail) GetLocAreaMonRepOk() (*[]LocationAreaMonReport, bool)`

GetLocAreaMonRepOk returns a tuple with the LocAreaMonRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocAreaMonRep

`func (o *SEALEventDetail) SetLocAreaMonRep(v []LocationAreaMonReport)`

SetLocAreaMonRep sets LocAreaMonRep field to given value.

### HasLocAreaMonRep

`func (o *SEALEventDetail) HasLocAreaMonRep() bool`

HasLocAreaMonRep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


