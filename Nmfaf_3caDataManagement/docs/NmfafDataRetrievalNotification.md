# NmfafDataRetrievalNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorreId** | **string** | Represents the Analytics Consumer Notification Correlation ID or Data Consumer Notification Correlation ID. It shall be set to the same value as the \&quot;correId\&quot; attribute of MessageConfiguration data type.  | 
**DataAnaNotif** | Pointer to [**NmfafDataAnaNotification**](NmfafDataAnaNotification.md) |  | [optional] 
**FetchInstruction** | Pointer to [**FetchInstruction**](FetchInstruction.md) |  | [optional] 

## Methods

### NewNmfafDataRetrievalNotification

`func NewNmfafDataRetrievalNotification(correId string, ) *NmfafDataRetrievalNotification`

NewNmfafDataRetrievalNotification instantiates a new NmfafDataRetrievalNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNmfafDataRetrievalNotificationWithDefaults

`func NewNmfafDataRetrievalNotificationWithDefaults() *NmfafDataRetrievalNotification`

NewNmfafDataRetrievalNotificationWithDefaults instantiates a new NmfafDataRetrievalNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorreId

`func (o *NmfafDataRetrievalNotification) GetCorreId() string`

GetCorreId returns the CorreId field if non-nil, zero value otherwise.

### GetCorreIdOk

`func (o *NmfafDataRetrievalNotification) GetCorreIdOk() (*string, bool)`

GetCorreIdOk returns a tuple with the CorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorreId

`func (o *NmfafDataRetrievalNotification) SetCorreId(v string)`

SetCorreId sets CorreId field to given value.


### GetDataAnaNotif

`func (o *NmfafDataRetrievalNotification) GetDataAnaNotif() NmfafDataAnaNotification`

GetDataAnaNotif returns the DataAnaNotif field if non-nil, zero value otherwise.

### GetDataAnaNotifOk

`func (o *NmfafDataRetrievalNotification) GetDataAnaNotifOk() (*NmfafDataAnaNotification, bool)`

GetDataAnaNotifOk returns a tuple with the DataAnaNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAnaNotif

`func (o *NmfafDataRetrievalNotification) SetDataAnaNotif(v NmfafDataAnaNotification)`

SetDataAnaNotif sets DataAnaNotif field to given value.

### HasDataAnaNotif

`func (o *NmfafDataRetrievalNotification) HasDataAnaNotif() bool`

HasDataAnaNotif returns a boolean if a field has been set.

### GetFetchInstruction

`func (o *NmfafDataRetrievalNotification) GetFetchInstruction() FetchInstruction`

GetFetchInstruction returns the FetchInstruction field if non-nil, zero value otherwise.

### GetFetchInstructionOk

`func (o *NmfafDataRetrievalNotification) GetFetchInstructionOk() (*FetchInstruction, bool)`

GetFetchInstructionOk returns a tuple with the FetchInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchInstruction

`func (o *NmfafDataRetrievalNotification) SetFetchInstruction(v FetchInstruction)`

SetFetchInstruction sets FetchInstruction field to given value.

### HasFetchInstruction

`func (o *NmfafDataRetrievalNotification) HasFetchInstruction() bool`

HasFetchInstruction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


