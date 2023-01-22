# SIPEventType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SIPMethod** | Pointer to **string** |  | [optional] 
**EventHeader** | Pointer to **string** |  | [optional] 
**ExpiresHeader** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 

## Methods

### NewSIPEventType

`func NewSIPEventType() *SIPEventType`

NewSIPEventType instantiates a new SIPEventType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIPEventTypeWithDefaults

`func NewSIPEventTypeWithDefaults() *SIPEventType`

NewSIPEventTypeWithDefaults instantiates a new SIPEventType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSIPMethod

`func (o *SIPEventType) GetSIPMethod() string`

GetSIPMethod returns the SIPMethod field if non-nil, zero value otherwise.

### GetSIPMethodOk

`func (o *SIPEventType) GetSIPMethodOk() (*string, bool)`

GetSIPMethodOk returns a tuple with the SIPMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSIPMethod

`func (o *SIPEventType) SetSIPMethod(v string)`

SetSIPMethod sets SIPMethod field to given value.

### HasSIPMethod

`func (o *SIPEventType) HasSIPMethod() bool`

HasSIPMethod returns a boolean if a field has been set.

### GetEventHeader

`func (o *SIPEventType) GetEventHeader() string`

GetEventHeader returns the EventHeader field if non-nil, zero value otherwise.

### GetEventHeaderOk

`func (o *SIPEventType) GetEventHeaderOk() (*string, bool)`

GetEventHeaderOk returns a tuple with the EventHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHeader

`func (o *SIPEventType) SetEventHeader(v string)`

SetEventHeader sets EventHeader field to given value.

### HasEventHeader

`func (o *SIPEventType) HasEventHeader() bool`

HasEventHeader returns a boolean if a field has been set.

### GetExpiresHeader

`func (o *SIPEventType) GetExpiresHeader() int32`

GetExpiresHeader returns the ExpiresHeader field if non-nil, zero value otherwise.

### GetExpiresHeaderOk

`func (o *SIPEventType) GetExpiresHeaderOk() (*int32, bool)`

GetExpiresHeaderOk returns a tuple with the ExpiresHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresHeader

`func (o *SIPEventType) SetExpiresHeader(v int32)`

SetExpiresHeader sets ExpiresHeader field to given value.

### HasExpiresHeader

`func (o *SIPEventType) HasExpiresHeader() bool`

HasExpiresHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


