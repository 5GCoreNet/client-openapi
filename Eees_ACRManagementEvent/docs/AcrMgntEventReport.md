# AcrMgntEventReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**AcrMgntEvent**](AcrMgntEvent.md) |  | 
**TimeStamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**UpPathChgInfo** | Pointer to [**UpPathChangeInfo**](UpPathChangeInfo.md) |  | [optional] 
**EasEndPoint** | Pointer to [**EndPoint**](EndPoint.md) |  | [optional] 
**ActStatus** | Pointer to [**ActStatus**](ActStatus.md) |  | [optional] 

## Methods

### NewAcrMgntEventReport

`func NewAcrMgntEventReport(event AcrMgntEvent, ) *AcrMgntEventReport`

NewAcrMgntEventReport instantiates a new AcrMgntEventReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcrMgntEventReportWithDefaults

`func NewAcrMgntEventReportWithDefaults() *AcrMgntEventReport`

NewAcrMgntEventReportWithDefaults instantiates a new AcrMgntEventReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *AcrMgntEventReport) GetEvent() AcrMgntEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AcrMgntEventReport) GetEventOk() (*AcrMgntEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AcrMgntEventReport) SetEvent(v AcrMgntEvent)`

SetEvent sets Event field to given value.


### GetTimeStamp

`func (o *AcrMgntEventReport) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *AcrMgntEventReport) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *AcrMgntEventReport) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *AcrMgntEventReport) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetUpPathChgInfo

`func (o *AcrMgntEventReport) GetUpPathChgInfo() UpPathChangeInfo`

GetUpPathChgInfo returns the UpPathChgInfo field if non-nil, zero value otherwise.

### GetUpPathChgInfoOk

`func (o *AcrMgntEventReport) GetUpPathChgInfoOk() (*UpPathChangeInfo, bool)`

GetUpPathChgInfoOk returns a tuple with the UpPathChgInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPathChgInfo

`func (o *AcrMgntEventReport) SetUpPathChgInfo(v UpPathChangeInfo)`

SetUpPathChgInfo sets UpPathChgInfo field to given value.

### HasUpPathChgInfo

`func (o *AcrMgntEventReport) HasUpPathChgInfo() bool`

HasUpPathChgInfo returns a boolean if a field has been set.

### GetEasEndPoint

`func (o *AcrMgntEventReport) GetEasEndPoint() EndPoint`

GetEasEndPoint returns the EasEndPoint field if non-nil, zero value otherwise.

### GetEasEndPointOk

`func (o *AcrMgntEventReport) GetEasEndPointOk() (*EndPoint, bool)`

GetEasEndPointOk returns a tuple with the EasEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasEndPoint

`func (o *AcrMgntEventReport) SetEasEndPoint(v EndPoint)`

SetEasEndPoint sets EasEndPoint field to given value.

### HasEasEndPoint

`func (o *AcrMgntEventReport) HasEasEndPoint() bool`

HasEasEndPoint returns a boolean if a field has been set.

### GetActStatus

`func (o *AcrMgntEventReport) GetActStatus() ActStatus`

GetActStatus returns the ActStatus field if non-nil, zero value otherwise.

### GetActStatusOk

`func (o *AcrMgntEventReport) GetActStatusOk() (*ActStatus, bool)`

GetActStatusOk returns a tuple with the ActStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActStatus

`func (o *AcrMgntEventReport) SetActStatus(v ActStatus)`

SetActStatus sets ActStatus field to given value.

### HasActStatus

`func (o *AcrMgntEventReport) HasActStatus() bool`

HasActStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


