# PfdData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalAppId** | **string** | Each element uniquely external application identifier | 
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**Pfds** | [**map[string]Pfd**](Pfd.md) | Contains the PFDs of the external application identifier. Each PFD is identified in the map via a key containing the PFD identifier. | 
**AllowedDelay** | Pointer to **NullableInt32** | Unsigned integer identifying a period of time in units of seconds with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**CachingTime** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds with \&quot;readOnly&#x3D;true\&quot; property. | [optional] [readonly] 

## Methods

### NewPfdData

`func NewPfdData(externalAppId string, pfds map[string]Pfd, ) *PfdData`

NewPfdData instantiates a new PfdData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPfdDataWithDefaults

`func NewPfdDataWithDefaults() *PfdData`

NewPfdDataWithDefaults instantiates a new PfdData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalAppId

`func (o *PfdData) GetExternalAppId() string`

GetExternalAppId returns the ExternalAppId field if non-nil, zero value otherwise.

### GetExternalAppIdOk

`func (o *PfdData) GetExternalAppIdOk() (*string, bool)`

GetExternalAppIdOk returns a tuple with the ExternalAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAppId

`func (o *PfdData) SetExternalAppId(v string)`

SetExternalAppId sets ExternalAppId field to given value.


### GetSelf

`func (o *PfdData) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PfdData) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PfdData) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PfdData) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetPfds

`func (o *PfdData) GetPfds() map[string]Pfd`

GetPfds returns the Pfds field if non-nil, zero value otherwise.

### GetPfdsOk

`func (o *PfdData) GetPfdsOk() (*map[string]Pfd, bool)`

GetPfdsOk returns a tuple with the Pfds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfds

`func (o *PfdData) SetPfds(v map[string]Pfd)`

SetPfds sets Pfds field to given value.


### GetAllowedDelay

`func (o *PfdData) GetAllowedDelay() int32`

GetAllowedDelay returns the AllowedDelay field if non-nil, zero value otherwise.

### GetAllowedDelayOk

`func (o *PfdData) GetAllowedDelayOk() (*int32, bool)`

GetAllowedDelayOk returns a tuple with the AllowedDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDelay

`func (o *PfdData) SetAllowedDelay(v int32)`

SetAllowedDelay sets AllowedDelay field to given value.

### HasAllowedDelay

`func (o *PfdData) HasAllowedDelay() bool`

HasAllowedDelay returns a boolean if a field has been set.

### SetAllowedDelayNil

`func (o *PfdData) SetAllowedDelayNil(b bool)`

 SetAllowedDelayNil sets the value for AllowedDelay to be an explicit nil

### UnsetAllowedDelay
`func (o *PfdData) UnsetAllowedDelay()`

UnsetAllowedDelay ensures that no value is present for AllowedDelay, not even an explicit nil
### GetCachingTime

`func (o *PfdData) GetCachingTime() int32`

GetCachingTime returns the CachingTime field if non-nil, zero value otherwise.

### GetCachingTimeOk

`func (o *PfdData) GetCachingTimeOk() (*int32, bool)`

GetCachingTimeOk returns a tuple with the CachingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingTime

`func (o *PfdData) SetCachingTime(v int32)`

SetCachingTime sets CachingTime field to given value.

### HasCachingTime

`func (o *PfdData) HasCachingTime() bool`

HasCachingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


