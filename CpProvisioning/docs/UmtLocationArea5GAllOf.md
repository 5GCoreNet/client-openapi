# UmtLocationArea5GAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UmtTime** | Pointer to **string** | String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC). | [optional] 
**UmtDuration** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 

## Methods

### NewUmtLocationArea5GAllOf

`func NewUmtLocationArea5GAllOf() *UmtLocationArea5GAllOf`

NewUmtLocationArea5GAllOf instantiates a new UmtLocationArea5GAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmtLocationArea5GAllOfWithDefaults

`func NewUmtLocationArea5GAllOfWithDefaults() *UmtLocationArea5GAllOf`

NewUmtLocationArea5GAllOfWithDefaults instantiates a new UmtLocationArea5GAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUmtTime

`func (o *UmtLocationArea5GAllOf) GetUmtTime() string`

GetUmtTime returns the UmtTime field if non-nil, zero value otherwise.

### GetUmtTimeOk

`func (o *UmtLocationArea5GAllOf) GetUmtTimeOk() (*string, bool)`

GetUmtTimeOk returns a tuple with the UmtTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmtTime

`func (o *UmtLocationArea5GAllOf) SetUmtTime(v string)`

SetUmtTime sets UmtTime field to given value.

### HasUmtTime

`func (o *UmtLocationArea5GAllOf) HasUmtTime() bool`

HasUmtTime returns a boolean if a field has been set.

### GetUmtDuration

`func (o *UmtLocationArea5GAllOf) GetUmtDuration() int32`

GetUmtDuration returns the UmtDuration field if non-nil, zero value otherwise.

### GetUmtDurationOk

`func (o *UmtLocationArea5GAllOf) GetUmtDurationOk() (*int32, bool)`

GetUmtDurationOk returns a tuple with the UmtDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmtDuration

`func (o *UmtLocationArea5GAllOf) SetUmtDuration(v int32)`

SetUmtDuration sets UmtDuration field to given value.

### HasUmtDuration

`func (o *UmtLocationArea5GAllOf) HasUmtDuration() bool`

HasUmtDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


