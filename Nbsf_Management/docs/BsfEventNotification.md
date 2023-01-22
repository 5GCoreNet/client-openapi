# BsfEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**BsfEvent**](BsfEvent.md) |  | 
**PcfForUeInfo** | Pointer to [**PcfForUeInfo**](PcfForUeInfo.md) |  | [optional] 
**PcfForPduSessInfos** | Pointer to [**[]PcfForPduSessionInfo**](PcfForPduSessionInfo.md) | The information of the PCF for a PDU session. | [optional] 
**MatchSnssaiDnns** | Pointer to [**[]SnssaiDnnPair**](SnssaiDnnPair.md) | Matching S-NSSAI and DNN pairs. | [optional] 

## Methods

### NewBsfEventNotification

`func NewBsfEventNotification(event BsfEvent, ) *BsfEventNotification`

NewBsfEventNotification instantiates a new BsfEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBsfEventNotificationWithDefaults

`func NewBsfEventNotificationWithDefaults() *BsfEventNotification`

NewBsfEventNotificationWithDefaults instantiates a new BsfEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *BsfEventNotification) GetEvent() BsfEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *BsfEventNotification) GetEventOk() (*BsfEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *BsfEventNotification) SetEvent(v BsfEvent)`

SetEvent sets Event field to given value.


### GetPcfForUeInfo

`func (o *BsfEventNotification) GetPcfForUeInfo() PcfForUeInfo`

GetPcfForUeInfo returns the PcfForUeInfo field if non-nil, zero value otherwise.

### GetPcfForUeInfoOk

`func (o *BsfEventNotification) GetPcfForUeInfoOk() (*PcfForUeInfo, bool)`

GetPcfForUeInfoOk returns a tuple with the PcfForUeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfForUeInfo

`func (o *BsfEventNotification) SetPcfForUeInfo(v PcfForUeInfo)`

SetPcfForUeInfo sets PcfForUeInfo field to given value.

### HasPcfForUeInfo

`func (o *BsfEventNotification) HasPcfForUeInfo() bool`

HasPcfForUeInfo returns a boolean if a field has been set.

### GetPcfForPduSessInfos

`func (o *BsfEventNotification) GetPcfForPduSessInfos() []PcfForPduSessionInfo`

GetPcfForPduSessInfos returns the PcfForPduSessInfos field if non-nil, zero value otherwise.

### GetPcfForPduSessInfosOk

`func (o *BsfEventNotification) GetPcfForPduSessInfosOk() (*[]PcfForPduSessionInfo, bool)`

GetPcfForPduSessInfosOk returns a tuple with the PcfForPduSessInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfForPduSessInfos

`func (o *BsfEventNotification) SetPcfForPduSessInfos(v []PcfForPduSessionInfo)`

SetPcfForPduSessInfos sets PcfForPduSessInfos field to given value.

### HasPcfForPduSessInfos

`func (o *BsfEventNotification) HasPcfForPduSessInfos() bool`

HasPcfForPduSessInfos returns a boolean if a field has been set.

### GetMatchSnssaiDnns

`func (o *BsfEventNotification) GetMatchSnssaiDnns() []SnssaiDnnPair`

GetMatchSnssaiDnns returns the MatchSnssaiDnns field if non-nil, zero value otherwise.

### GetMatchSnssaiDnnsOk

`func (o *BsfEventNotification) GetMatchSnssaiDnnsOk() (*[]SnssaiDnnPair, bool)`

GetMatchSnssaiDnnsOk returns a tuple with the MatchSnssaiDnns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSnssaiDnns

`func (o *BsfEventNotification) SetMatchSnssaiDnns(v []SnssaiDnnPair)`

SetMatchSnssaiDnns sets MatchSnssaiDnns field to given value.

### HasMatchSnssaiDnns

`func (o *BsfEventNotification) HasMatchSnssaiDnns() bool`

HasMatchSnssaiDnns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


