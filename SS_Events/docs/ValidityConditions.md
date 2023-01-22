# ValidityConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocArea** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**TmWdws** | Pointer to [**[]TimeWindow**](TimeWindow.md) | Time window validity conditions. | [optional] 

## Methods

### NewValidityConditions

`func NewValidityConditions() *ValidityConditions`

NewValidityConditions instantiates a new ValidityConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidityConditionsWithDefaults

`func NewValidityConditionsWithDefaults() *ValidityConditions`

NewValidityConditionsWithDefaults instantiates a new ValidityConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocArea

`func (o *ValidityConditions) GetLocArea() LocationArea5G`

GetLocArea returns the LocArea field if non-nil, zero value otherwise.

### GetLocAreaOk

`func (o *ValidityConditions) GetLocAreaOk() (*LocationArea5G, bool)`

GetLocAreaOk returns a tuple with the LocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocArea

`func (o *ValidityConditions) SetLocArea(v LocationArea5G)`

SetLocArea sets LocArea field to given value.

### HasLocArea

`func (o *ValidityConditions) HasLocArea() bool`

HasLocArea returns a boolean if a field has been set.

### GetTmWdws

`func (o *ValidityConditions) GetTmWdws() []TimeWindow`

GetTmWdws returns the TmWdws field if non-nil, zero value otherwise.

### GetTmWdwsOk

`func (o *ValidityConditions) GetTmWdwsOk() (*[]TimeWindow, bool)`

GetTmWdwsOk returns a tuple with the TmWdws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmWdws

`func (o *ValidityConditions) SetTmWdws(v []TimeWindow)`

SetTmWdws sets TmWdws field to given value.

### HasTmWdws

`func (o *ValidityConditions) HasTmWdws() bool`

HasTmWdws returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


