# TmgiAllocated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TmgiList** | [**[]Tmgi**](Tmgi.md) | One or more TMGI values | 
**ExpirationTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 

## Methods

### NewTmgiAllocated

`func NewTmgiAllocated(tmgiList []Tmgi, expirationTime time.Time, ) *TmgiAllocated`

NewTmgiAllocated instantiates a new TmgiAllocated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTmgiAllocatedWithDefaults

`func NewTmgiAllocatedWithDefaults() *TmgiAllocated`

NewTmgiAllocatedWithDefaults instantiates a new TmgiAllocated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTmgiList

`func (o *TmgiAllocated) GetTmgiList() []Tmgi`

GetTmgiList returns the TmgiList field if non-nil, zero value otherwise.

### GetTmgiListOk

`func (o *TmgiAllocated) GetTmgiListOk() (*[]Tmgi, bool)`

GetTmgiListOk returns a tuple with the TmgiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmgiList

`func (o *TmgiAllocated) SetTmgiList(v []Tmgi)`

SetTmgiList sets TmgiList field to given value.


### GetExpirationTime

`func (o *TmgiAllocated) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *TmgiAllocated) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *TmgiAllocated) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


