# ContextStatusNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsSessionId** | [**MbsSessionId**](MbsSessionId.md) |  | 
**AreaSessionId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 16-bit integer. | [optional] 
**N2MbsSmInfoList** | Pointer to [**[]N2MbsSmInfo**](N2MbsSmInfo.md) |  | [optional] 
**OperationEvents** | Pointer to [**[]OperationEvent**](OperationEvent.md) |  | [optional] 
**OperationStatus** | Pointer to [**OperationStatus**](OperationStatus.md) |  | [optional] 

## Methods

### NewContextStatusNotification

`func NewContextStatusNotification(mbsSessionId MbsSessionId, ) *ContextStatusNotification`

NewContextStatusNotification instantiates a new ContextStatusNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextStatusNotificationWithDefaults

`func NewContextStatusNotificationWithDefaults() *ContextStatusNotification`

NewContextStatusNotificationWithDefaults instantiates a new ContextStatusNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsSessionId

`func (o *ContextStatusNotification) GetMbsSessionId() MbsSessionId`

GetMbsSessionId returns the MbsSessionId field if non-nil, zero value otherwise.

### GetMbsSessionIdOk

`func (o *ContextStatusNotification) GetMbsSessionIdOk() (*MbsSessionId, bool)`

GetMbsSessionIdOk returns a tuple with the MbsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionId

`func (o *ContextStatusNotification) SetMbsSessionId(v MbsSessionId)`

SetMbsSessionId sets MbsSessionId field to given value.


### GetAreaSessionId

`func (o *ContextStatusNotification) GetAreaSessionId() int32`

GetAreaSessionId returns the AreaSessionId field if non-nil, zero value otherwise.

### GetAreaSessionIdOk

`func (o *ContextStatusNotification) GetAreaSessionIdOk() (*int32, bool)`

GetAreaSessionIdOk returns a tuple with the AreaSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaSessionId

`func (o *ContextStatusNotification) SetAreaSessionId(v int32)`

SetAreaSessionId sets AreaSessionId field to given value.

### HasAreaSessionId

`func (o *ContextStatusNotification) HasAreaSessionId() bool`

HasAreaSessionId returns a boolean if a field has been set.

### GetN2MbsSmInfoList

`func (o *ContextStatusNotification) GetN2MbsSmInfoList() []N2MbsSmInfo`

GetN2MbsSmInfoList returns the N2MbsSmInfoList field if non-nil, zero value otherwise.

### GetN2MbsSmInfoListOk

`func (o *ContextStatusNotification) GetN2MbsSmInfoListOk() (*[]N2MbsSmInfo, bool)`

GetN2MbsSmInfoListOk returns a tuple with the N2MbsSmInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2MbsSmInfoList

`func (o *ContextStatusNotification) SetN2MbsSmInfoList(v []N2MbsSmInfo)`

SetN2MbsSmInfoList sets N2MbsSmInfoList field to given value.

### HasN2MbsSmInfoList

`func (o *ContextStatusNotification) HasN2MbsSmInfoList() bool`

HasN2MbsSmInfoList returns a boolean if a field has been set.

### GetOperationEvents

`func (o *ContextStatusNotification) GetOperationEvents() []OperationEvent`

GetOperationEvents returns the OperationEvents field if non-nil, zero value otherwise.

### GetOperationEventsOk

`func (o *ContextStatusNotification) GetOperationEventsOk() (*[]OperationEvent, bool)`

GetOperationEventsOk returns a tuple with the OperationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationEvents

`func (o *ContextStatusNotification) SetOperationEvents(v []OperationEvent)`

SetOperationEvents sets OperationEvents field to given value.

### HasOperationEvents

`func (o *ContextStatusNotification) HasOperationEvents() bool`

HasOperationEvents returns a boolean if a field has been set.

### GetOperationStatus

`func (o *ContextStatusNotification) GetOperationStatus() OperationStatus`

GetOperationStatus returns the OperationStatus field if non-nil, zero value otherwise.

### GetOperationStatusOk

`func (o *ContextStatusNotification) GetOperationStatusOk() (*OperationStatus, bool)`

GetOperationStatusOk returns a tuple with the OperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStatus

`func (o *ContextStatusNotification) SetOperationStatus(v OperationStatus)`

SetOperationStatus sets OperationStatus field to given value.

### HasOperationStatus

`func (o *ContextStatusNotification) HasOperationStatus() bool`

HasOperationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


