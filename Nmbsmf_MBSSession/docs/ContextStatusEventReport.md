# ContextStatusEventReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**ContextStatusEventType**](ContextStatusEventType.md) |  | 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**QosInfo** | Pointer to [**QosInfo**](QosInfo.md) |  | [optional] 
**StatusInfo** | Pointer to [**MbsSessionActivityStatus**](MbsSessionActivityStatus.md) |  | [optional] 
**MbsServiceArea** | Pointer to [**MbsServiceArea**](MbsServiceArea.md) |  | [optional] 
**MbsServiceAreaInfoList** | Pointer to [**map[string]MbsServiceAreaInfo**](MbsServiceAreaInfo.md) | A map (list of key-value pairs) where the key identifies an areaSessionId  | [optional] 
**MulticastTransAddInfo** | Pointer to [**MulticastTransportAddressChangeInfo**](MulticastTransportAddressChangeInfo.md) |  | [optional] 
**MbsSecurityContext** | Pointer to [**MbsSecurityContext**](MbsSecurityContext.md) |  | [optional] 

## Methods

### NewContextStatusEventReport

`func NewContextStatusEventReport(eventType ContextStatusEventType, timeStamp time.Time, ) *ContextStatusEventReport`

NewContextStatusEventReport instantiates a new ContextStatusEventReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextStatusEventReportWithDefaults

`func NewContextStatusEventReportWithDefaults() *ContextStatusEventReport`

NewContextStatusEventReportWithDefaults instantiates a new ContextStatusEventReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ContextStatusEventReport) GetEventType() ContextStatusEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ContextStatusEventReport) GetEventTypeOk() (*ContextStatusEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ContextStatusEventReport) SetEventType(v ContextStatusEventType)`

SetEventType sets EventType field to given value.


### GetTimeStamp

`func (o *ContextStatusEventReport) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *ContextStatusEventReport) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *ContextStatusEventReport) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetQosInfo

`func (o *ContextStatusEventReport) GetQosInfo() QosInfo`

GetQosInfo returns the QosInfo field if non-nil, zero value otherwise.

### GetQosInfoOk

`func (o *ContextStatusEventReport) GetQosInfoOk() (*QosInfo, bool)`

GetQosInfoOk returns a tuple with the QosInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosInfo

`func (o *ContextStatusEventReport) SetQosInfo(v QosInfo)`

SetQosInfo sets QosInfo field to given value.

### HasQosInfo

`func (o *ContextStatusEventReport) HasQosInfo() bool`

HasQosInfo returns a boolean if a field has been set.

### GetStatusInfo

`func (o *ContextStatusEventReport) GetStatusInfo() MbsSessionActivityStatus`

GetStatusInfo returns the StatusInfo field if non-nil, zero value otherwise.

### GetStatusInfoOk

`func (o *ContextStatusEventReport) GetStatusInfoOk() (*MbsSessionActivityStatus, bool)`

GetStatusInfoOk returns a tuple with the StatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusInfo

`func (o *ContextStatusEventReport) SetStatusInfo(v MbsSessionActivityStatus)`

SetStatusInfo sets StatusInfo field to given value.

### HasStatusInfo

`func (o *ContextStatusEventReport) HasStatusInfo() bool`

HasStatusInfo returns a boolean if a field has been set.

### GetMbsServiceArea

`func (o *ContextStatusEventReport) GetMbsServiceArea() MbsServiceArea`

GetMbsServiceArea returns the MbsServiceArea field if non-nil, zero value otherwise.

### GetMbsServiceAreaOk

`func (o *ContextStatusEventReport) GetMbsServiceAreaOk() (*MbsServiceArea, bool)`

GetMbsServiceAreaOk returns a tuple with the MbsServiceArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceArea

`func (o *ContextStatusEventReport) SetMbsServiceArea(v MbsServiceArea)`

SetMbsServiceArea sets MbsServiceArea field to given value.

### HasMbsServiceArea

`func (o *ContextStatusEventReport) HasMbsServiceArea() bool`

HasMbsServiceArea returns a boolean if a field has been set.

### GetMbsServiceAreaInfoList

`func (o *ContextStatusEventReport) GetMbsServiceAreaInfoList() map[string]MbsServiceAreaInfo`

GetMbsServiceAreaInfoList returns the MbsServiceAreaInfoList field if non-nil, zero value otherwise.

### GetMbsServiceAreaInfoListOk

`func (o *ContextStatusEventReport) GetMbsServiceAreaInfoListOk() (*map[string]MbsServiceAreaInfo, bool)`

GetMbsServiceAreaInfoListOk returns a tuple with the MbsServiceAreaInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceAreaInfoList

`func (o *ContextStatusEventReport) SetMbsServiceAreaInfoList(v map[string]MbsServiceAreaInfo)`

SetMbsServiceAreaInfoList sets MbsServiceAreaInfoList field to given value.

### HasMbsServiceAreaInfoList

`func (o *ContextStatusEventReport) HasMbsServiceAreaInfoList() bool`

HasMbsServiceAreaInfoList returns a boolean if a field has been set.

### GetMulticastTransAddInfo

`func (o *ContextStatusEventReport) GetMulticastTransAddInfo() MulticastTransportAddressChangeInfo`

GetMulticastTransAddInfo returns the MulticastTransAddInfo field if non-nil, zero value otherwise.

### GetMulticastTransAddInfoOk

`func (o *ContextStatusEventReport) GetMulticastTransAddInfoOk() (*MulticastTransportAddressChangeInfo, bool)`

GetMulticastTransAddInfoOk returns a tuple with the MulticastTransAddInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastTransAddInfo

`func (o *ContextStatusEventReport) SetMulticastTransAddInfo(v MulticastTransportAddressChangeInfo)`

SetMulticastTransAddInfo sets MulticastTransAddInfo field to given value.

### HasMulticastTransAddInfo

`func (o *ContextStatusEventReport) HasMulticastTransAddInfo() bool`

HasMulticastTransAddInfo returns a boolean if a field has been set.

### GetMbsSecurityContext

`func (o *ContextStatusEventReport) GetMbsSecurityContext() MbsSecurityContext`

GetMbsSecurityContext returns the MbsSecurityContext field if non-nil, zero value otherwise.

### GetMbsSecurityContextOk

`func (o *ContextStatusEventReport) GetMbsSecurityContextOk() (*MbsSecurityContext, bool)`

GetMbsSecurityContextOk returns a tuple with the MbsSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSecurityContext

`func (o *ContextStatusEventReport) SetMbsSecurityContext(v MbsSecurityContext)`

SetMbsSecurityContext sets MbsSecurityContext field to given value.

### HasMbsSecurityContext

`func (o *ContextStatusEventReport) HasMbsSecurityContext() bool`

HasMbsSecurityContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


