# CreateRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsSession** | Pointer to [**ExtMbsSession**](ExtMbsSession.md) |  | [optional] 
**EventList** | Pointer to [**MbsSessionEventReportList**](MbsSessionEventReportList.md) |  | [optional] 

## Methods

### NewCreateRspData

`func NewCreateRspData() *CreateRspData`

NewCreateRspData instantiates a new CreateRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRspDataWithDefaults

`func NewCreateRspDataWithDefaults() *CreateRspData`

NewCreateRspDataWithDefaults instantiates a new CreateRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsSession

`func (o *CreateRspData) GetMbsSession() ExtMbsSession`

GetMbsSession returns the MbsSession field if non-nil, zero value otherwise.

### GetMbsSessionOk

`func (o *CreateRspData) GetMbsSessionOk() (*ExtMbsSession, bool)`

GetMbsSessionOk returns a tuple with the MbsSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSession

`func (o *CreateRspData) SetMbsSession(v ExtMbsSession)`

SetMbsSession sets MbsSession field to given value.

### HasMbsSession

`func (o *CreateRspData) HasMbsSession() bool`

HasMbsSession returns a boolean if a field has been set.

### GetEventList

`func (o *CreateRspData) GetEventList() MbsSessionEventReportList`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *CreateRspData) GetEventListOk() (*MbsSessionEventReportList, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *CreateRspData) SetEventList(v MbsSessionEventReportList)`

SetEventList sets EventList field to given value.

### HasEventList

`func (o *CreateRspData) HasEventList() bool`

HasEventList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


