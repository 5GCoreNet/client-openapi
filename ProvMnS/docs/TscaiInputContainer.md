# TscaiInputContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Periodicity** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**BurstArrivalTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SurTimeInNumMsg** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**SurTimeInTime** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewTscaiInputContainer

`func NewTscaiInputContainer() *TscaiInputContainer`

NewTscaiInputContainer instantiates a new TscaiInputContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTscaiInputContainerWithDefaults

`func NewTscaiInputContainerWithDefaults() *TscaiInputContainer`

NewTscaiInputContainerWithDefaults instantiates a new TscaiInputContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodicity

`func (o *TscaiInputContainer) GetPeriodicity() int32`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *TscaiInputContainer) GetPeriodicityOk() (*int32, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *TscaiInputContainer) SetPeriodicity(v int32)`

SetPeriodicity sets Periodicity field to given value.

### HasPeriodicity

`func (o *TscaiInputContainer) HasPeriodicity() bool`

HasPeriodicity returns a boolean if a field has been set.

### GetBurstArrivalTime

`func (o *TscaiInputContainer) GetBurstArrivalTime() time.Time`

GetBurstArrivalTime returns the BurstArrivalTime field if non-nil, zero value otherwise.

### GetBurstArrivalTimeOk

`func (o *TscaiInputContainer) GetBurstArrivalTimeOk() (*time.Time, bool)`

GetBurstArrivalTimeOk returns a tuple with the BurstArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstArrivalTime

`func (o *TscaiInputContainer) SetBurstArrivalTime(v time.Time)`

SetBurstArrivalTime sets BurstArrivalTime field to given value.

### HasBurstArrivalTime

`func (o *TscaiInputContainer) HasBurstArrivalTime() bool`

HasBurstArrivalTime returns a boolean if a field has been set.

### GetSurTimeInNumMsg

`func (o *TscaiInputContainer) GetSurTimeInNumMsg() int32`

GetSurTimeInNumMsg returns the SurTimeInNumMsg field if non-nil, zero value otherwise.

### GetSurTimeInNumMsgOk

`func (o *TscaiInputContainer) GetSurTimeInNumMsgOk() (*int32, bool)`

GetSurTimeInNumMsgOk returns a tuple with the SurTimeInNumMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurTimeInNumMsg

`func (o *TscaiInputContainer) SetSurTimeInNumMsg(v int32)`

SetSurTimeInNumMsg sets SurTimeInNumMsg field to given value.

### HasSurTimeInNumMsg

`func (o *TscaiInputContainer) HasSurTimeInNumMsg() bool`

HasSurTimeInNumMsg returns a boolean if a field has been set.

### GetSurTimeInTime

`func (o *TscaiInputContainer) GetSurTimeInTime() int32`

GetSurTimeInTime returns the SurTimeInTime field if non-nil, zero value otherwise.

### GetSurTimeInTimeOk

`func (o *TscaiInputContainer) GetSurTimeInTimeOk() (*int32, bool)`

GetSurTimeInTimeOk returns a tuple with the SurTimeInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurTimeInTime

`func (o *TscaiInputContainer) SetSurTimeInTime(v int32)`

SetSurTimeInTime sets SurTimeInTime field to given value.

### HasSurTimeInTime

`func (o *TscaiInputContainer) HasSurTimeInTime() bool`

HasSurTimeInTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


