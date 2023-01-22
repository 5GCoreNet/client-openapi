# Timer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimerId** | Pointer to **string** | Represents the identifier of a timer. | [optional] 
**Expires** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**MetaTags** | Pointer to **map[string][]string** | A map (list of key-value pairs where a tagName of type string serves as key) of tagValue lists | [optional] 
**CallbackReference** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**DeleteAfter** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewTimer

`func NewTimer(expires time.Time, ) *Timer`

NewTimer instantiates a new Timer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimerWithDefaults

`func NewTimerWithDefaults() *Timer`

NewTimerWithDefaults instantiates a new Timer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimerId

`func (o *Timer) GetTimerId() string`

GetTimerId returns the TimerId field if non-nil, zero value otherwise.

### GetTimerIdOk

`func (o *Timer) GetTimerIdOk() (*string, bool)`

GetTimerIdOk returns a tuple with the TimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerId

`func (o *Timer) SetTimerId(v string)`

SetTimerId sets TimerId field to given value.

### HasTimerId

`func (o *Timer) HasTimerId() bool`

HasTimerId returns a boolean if a field has been set.

### GetExpires

`func (o *Timer) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Timer) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Timer) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### GetMetaTags

`func (o *Timer) GetMetaTags() map[string][]string`

GetMetaTags returns the MetaTags field if non-nil, zero value otherwise.

### GetMetaTagsOk

`func (o *Timer) GetMetaTagsOk() (*map[string][]string, bool)`

GetMetaTagsOk returns a tuple with the MetaTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTags

`func (o *Timer) SetMetaTags(v map[string][]string)`

SetMetaTags sets MetaTags field to given value.

### HasMetaTags

`func (o *Timer) HasMetaTags() bool`

HasMetaTags returns a boolean if a field has been set.

### GetCallbackReference

`func (o *Timer) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *Timer) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *Timer) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.

### HasCallbackReference

`func (o *Timer) HasCallbackReference() bool`

HasCallbackReference returns a boolean if a field has been set.

### GetDeleteAfter

`func (o *Timer) GetDeleteAfter() int32`

GetDeleteAfter returns the DeleteAfter field if non-nil, zero value otherwise.

### GetDeleteAfterOk

`func (o *Timer) GetDeleteAfterOk() (*int32, bool)`

GetDeleteAfterOk returns a tuple with the DeleteAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfter

`func (o *Timer) SetDeleteAfter(v int32)`

SetDeleteAfter sets DeleteAfter field to given value.

### HasDeleteAfter

`func (o *Timer) HasDeleteAfter() bool`

HasDeleteAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


