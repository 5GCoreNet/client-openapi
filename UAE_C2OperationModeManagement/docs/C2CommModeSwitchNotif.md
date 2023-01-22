# C2CommModeSwitchNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UaeServerId** | **string** | string providing an URI formatted according to IETF RFC 3986. | 
**UasId** | [**UasId**](UasId.md) |  | 
**C2CommModeSwitchType** | [**C2CommModeSwitching**](C2CommModeSwitching.md) |  | 
**SwitchingCause** | Pointer to [**C2SwitchingCause**](C2SwitchingCause.md) |  | [optional] 

## Methods

### NewC2CommModeSwitchNotif

`func NewC2CommModeSwitchNotif(uaeServerId string, uasId UasId, c2CommModeSwitchType C2CommModeSwitching, ) *C2CommModeSwitchNotif`

NewC2CommModeSwitchNotif instantiates a new C2CommModeSwitchNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC2CommModeSwitchNotifWithDefaults

`func NewC2CommModeSwitchNotifWithDefaults() *C2CommModeSwitchNotif`

NewC2CommModeSwitchNotifWithDefaults instantiates a new C2CommModeSwitchNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUaeServerId

`func (o *C2CommModeSwitchNotif) GetUaeServerId() string`

GetUaeServerId returns the UaeServerId field if non-nil, zero value otherwise.

### GetUaeServerIdOk

`func (o *C2CommModeSwitchNotif) GetUaeServerIdOk() (*string, bool)`

GetUaeServerIdOk returns a tuple with the UaeServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUaeServerId

`func (o *C2CommModeSwitchNotif) SetUaeServerId(v string)`

SetUaeServerId sets UaeServerId field to given value.


### GetUasId

`func (o *C2CommModeSwitchNotif) GetUasId() UasId`

GetUasId returns the UasId field if non-nil, zero value otherwise.

### GetUasIdOk

`func (o *C2CommModeSwitchNotif) GetUasIdOk() (*UasId, bool)`

GetUasIdOk returns a tuple with the UasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUasId

`func (o *C2CommModeSwitchNotif) SetUasId(v UasId)`

SetUasId sets UasId field to given value.


### GetC2CommModeSwitchType

`func (o *C2CommModeSwitchNotif) GetC2CommModeSwitchType() C2CommModeSwitching`

GetC2CommModeSwitchType returns the C2CommModeSwitchType field if non-nil, zero value otherwise.

### GetC2CommModeSwitchTypeOk

`func (o *C2CommModeSwitchNotif) GetC2CommModeSwitchTypeOk() (*C2CommModeSwitching, bool)`

GetC2CommModeSwitchTypeOk returns a tuple with the C2CommModeSwitchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC2CommModeSwitchType

`func (o *C2CommModeSwitchNotif) SetC2CommModeSwitchType(v C2CommModeSwitching)`

SetC2CommModeSwitchType sets C2CommModeSwitchType field to given value.


### GetSwitchingCause

`func (o *C2CommModeSwitchNotif) GetSwitchingCause() C2SwitchingCause`

GetSwitchingCause returns the SwitchingCause field if non-nil, zero value otherwise.

### GetSwitchingCauseOk

`func (o *C2CommModeSwitchNotif) GetSwitchingCauseOk() (*C2SwitchingCause, bool)`

GetSwitchingCauseOk returns a tuple with the SwitchingCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchingCause

`func (o *C2CommModeSwitchNotif) SetSwitchingCause(v C2SwitchingCause)`

SetSwitchingCause sets SwitchingCause field to given value.

### HasSwitchingCause

`func (o *C2CommModeSwitchNotif) HasSwitchingCause() bool`

HasSwitchingCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


