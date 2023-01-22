# NwdafData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NwdafInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**NwdafEvents** | Pointer to [**[]NwdafEvent**](NwdafEvent.md) |  | [optional] 

## Methods

### NewNwdafData

`func NewNwdafData(nwdafInstanceId string, ) *NwdafData`

NewNwdafData instantiates a new NwdafData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNwdafDataWithDefaults

`func NewNwdafDataWithDefaults() *NwdafData`

NewNwdafDataWithDefaults instantiates a new NwdafData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNwdafInstanceId

`func (o *NwdafData) GetNwdafInstanceId() string`

GetNwdafInstanceId returns the NwdafInstanceId field if non-nil, zero value otherwise.

### GetNwdafInstanceIdOk

`func (o *NwdafData) GetNwdafInstanceIdOk() (*string, bool)`

GetNwdafInstanceIdOk returns a tuple with the NwdafInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafInstanceId

`func (o *NwdafData) SetNwdafInstanceId(v string)`

SetNwdafInstanceId sets NwdafInstanceId field to given value.


### GetNwdafEvents

`func (o *NwdafData) GetNwdafEvents() []NwdafEvent`

GetNwdafEvents returns the NwdafEvents field if non-nil, zero value otherwise.

### GetNwdafEventsOk

`func (o *NwdafData) GetNwdafEventsOk() (*[]NwdafEvent, bool)`

GetNwdafEventsOk returns a tuple with the NwdafEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafEvents

`func (o *NwdafData) SetNwdafEvents(v []NwdafEvent)`

SetNwdafEvents sets NwdafEvents field to given value.

### HasNwdafEvents

`func (o *NwdafData) HasNwdafEvents() bool`

HasNwdafEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


