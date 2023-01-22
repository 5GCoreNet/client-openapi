# MbsSessionEventReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**MbsSessionEventType**](MbsSessionEventType.md) |  | 
**TimeStamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**IngressTunAddrInfo** | Pointer to [**IngressTunAddrInfo**](IngressTunAddrInfo.md) |  | [optional] 
**BroadcastDelStatus** | Pointer to [**BroadcastDeliveryStatus**](BroadcastDeliveryStatus.md) |  | [optional] 

## Methods

### NewMbsSessionEventReport

`func NewMbsSessionEventReport(eventType MbsSessionEventType, ) *MbsSessionEventReport`

NewMbsSessionEventReport instantiates a new MbsSessionEventReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsSessionEventReportWithDefaults

`func NewMbsSessionEventReportWithDefaults() *MbsSessionEventReport`

NewMbsSessionEventReportWithDefaults instantiates a new MbsSessionEventReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *MbsSessionEventReport) GetEventType() MbsSessionEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MbsSessionEventReport) GetEventTypeOk() (*MbsSessionEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MbsSessionEventReport) SetEventType(v MbsSessionEventType)`

SetEventType sets EventType field to given value.


### GetTimeStamp

`func (o *MbsSessionEventReport) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *MbsSessionEventReport) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *MbsSessionEventReport) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *MbsSessionEventReport) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetIngressTunAddrInfo

`func (o *MbsSessionEventReport) GetIngressTunAddrInfo() IngressTunAddrInfo`

GetIngressTunAddrInfo returns the IngressTunAddrInfo field if non-nil, zero value otherwise.

### GetIngressTunAddrInfoOk

`func (o *MbsSessionEventReport) GetIngressTunAddrInfoOk() (*IngressTunAddrInfo, bool)`

GetIngressTunAddrInfoOk returns a tuple with the IngressTunAddrInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressTunAddrInfo

`func (o *MbsSessionEventReport) SetIngressTunAddrInfo(v IngressTunAddrInfo)`

SetIngressTunAddrInfo sets IngressTunAddrInfo field to given value.

### HasIngressTunAddrInfo

`func (o *MbsSessionEventReport) HasIngressTunAddrInfo() bool`

HasIngressTunAddrInfo returns a boolean if a field has been set.

### GetBroadcastDelStatus

`func (o *MbsSessionEventReport) GetBroadcastDelStatus() BroadcastDeliveryStatus`

GetBroadcastDelStatus returns the BroadcastDelStatus field if non-nil, zero value otherwise.

### GetBroadcastDelStatusOk

`func (o *MbsSessionEventReport) GetBroadcastDelStatusOk() (*BroadcastDeliveryStatus, bool)`

GetBroadcastDelStatusOk returns a tuple with the BroadcastDelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastDelStatus

`func (o *MbsSessionEventReport) SetBroadcastDelStatus(v BroadcastDeliveryStatus)`

SetBroadcastDelStatus sets BroadcastDelStatus field to given value.

### HasBroadcastDelStatus

`func (o *MbsSessionEventReport) HasBroadcastDelStatus() bool`

HasBroadcastDelStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


