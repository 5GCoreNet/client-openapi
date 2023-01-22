# SelectedC2CommModeNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UasId** | [**UasId**](UasId.md) |  | 
**SelPrimaryC2CommMode** | [**C2CommMode**](C2CommMode.md) |  | 
**SelSecondaryC2CommMode** | Pointer to [**C2CommMode**](C2CommMode.md) |  | [optional] 

## Methods

### NewSelectedC2CommModeNotif

`func NewSelectedC2CommModeNotif(uasId UasId, selPrimaryC2CommMode C2CommMode, ) *SelectedC2CommModeNotif`

NewSelectedC2CommModeNotif instantiates a new SelectedC2CommModeNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectedC2CommModeNotifWithDefaults

`func NewSelectedC2CommModeNotifWithDefaults() *SelectedC2CommModeNotif`

NewSelectedC2CommModeNotifWithDefaults instantiates a new SelectedC2CommModeNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUasId

`func (o *SelectedC2CommModeNotif) GetUasId() UasId`

GetUasId returns the UasId field if non-nil, zero value otherwise.

### GetUasIdOk

`func (o *SelectedC2CommModeNotif) GetUasIdOk() (*UasId, bool)`

GetUasIdOk returns a tuple with the UasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUasId

`func (o *SelectedC2CommModeNotif) SetUasId(v UasId)`

SetUasId sets UasId field to given value.


### GetSelPrimaryC2CommMode

`func (o *SelectedC2CommModeNotif) GetSelPrimaryC2CommMode() C2CommMode`

GetSelPrimaryC2CommMode returns the SelPrimaryC2CommMode field if non-nil, zero value otherwise.

### GetSelPrimaryC2CommModeOk

`func (o *SelectedC2CommModeNotif) GetSelPrimaryC2CommModeOk() (*C2CommMode, bool)`

GetSelPrimaryC2CommModeOk returns a tuple with the SelPrimaryC2CommMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelPrimaryC2CommMode

`func (o *SelectedC2CommModeNotif) SetSelPrimaryC2CommMode(v C2CommMode)`

SetSelPrimaryC2CommMode sets SelPrimaryC2CommMode field to given value.


### GetSelSecondaryC2CommMode

`func (o *SelectedC2CommModeNotif) GetSelSecondaryC2CommMode() C2CommMode`

GetSelSecondaryC2CommMode returns the SelSecondaryC2CommMode field if non-nil, zero value otherwise.

### GetSelSecondaryC2CommModeOk

`func (o *SelectedC2CommModeNotif) GetSelSecondaryC2CommModeOk() (*C2CommMode, bool)`

GetSelSecondaryC2CommModeOk returns a tuple with the SelSecondaryC2CommMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelSecondaryC2CommMode

`func (o *SelectedC2CommModeNotif) SetSelSecondaryC2CommMode(v C2CommMode)`

SetSelSecondaryC2CommMode sets SelSecondaryC2CommMode field to given value.

### HasSelSecondaryC2CommMode

`func (o *SelectedC2CommModeNotif) HasSelSecondaryC2CommMode() bool`

HasSelSecondaryC2CommMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


