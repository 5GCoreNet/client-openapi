# RTUavStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UavId** | Pointer to [**UavId**](UavId.md) |  | [optional] 
**UavNetConnStatus** | Pointer to [**UavNetConnStatus**](UavNetConnStatus.md) |  | [optional] 
**UavLocInfo** | Pointer to [**LocationInfo**](LocationInfo.md) |  | [optional] 

## Methods

### NewRTUavStatus

`func NewRTUavStatus() *RTUavStatus`

NewRTUavStatus instantiates a new RTUavStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRTUavStatusWithDefaults

`func NewRTUavStatusWithDefaults() *RTUavStatus`

NewRTUavStatusWithDefaults instantiates a new RTUavStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUavId

`func (o *RTUavStatus) GetUavId() UavId`

GetUavId returns the UavId field if non-nil, zero value otherwise.

### GetUavIdOk

`func (o *RTUavStatus) GetUavIdOk() (*UavId, bool)`

GetUavIdOk returns a tuple with the UavId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUavId

`func (o *RTUavStatus) SetUavId(v UavId)`

SetUavId sets UavId field to given value.

### HasUavId

`func (o *RTUavStatus) HasUavId() bool`

HasUavId returns a boolean if a field has been set.

### GetUavNetConnStatus

`func (o *RTUavStatus) GetUavNetConnStatus() UavNetConnStatus`

GetUavNetConnStatus returns the UavNetConnStatus field if non-nil, zero value otherwise.

### GetUavNetConnStatusOk

`func (o *RTUavStatus) GetUavNetConnStatusOk() (*UavNetConnStatus, bool)`

GetUavNetConnStatusOk returns a tuple with the UavNetConnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUavNetConnStatus

`func (o *RTUavStatus) SetUavNetConnStatus(v UavNetConnStatus)`

SetUavNetConnStatus sets UavNetConnStatus field to given value.

### HasUavNetConnStatus

`func (o *RTUavStatus) HasUavNetConnStatus() bool`

HasUavNetConnStatus returns a boolean if a field has been set.

### GetUavLocInfo

`func (o *RTUavStatus) GetUavLocInfo() LocationInfo`

GetUavLocInfo returns the UavLocInfo field if non-nil, zero value otherwise.

### GetUavLocInfoOk

`func (o *RTUavStatus) GetUavLocInfoOk() (*LocationInfo, bool)`

GetUavLocInfoOk returns a tuple with the UavLocInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUavLocInfo

`func (o *RTUavStatus) SetUavLocInfo(v LocationInfo)`

SetUavLocInfo sets UavLocInfo field to given value.

### HasUavLocInfo

`func (o *RTUavStatus) HasUavLocInfo() bool`

HasUavLocInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


