# ThresholdInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThresholdDirection** | Pointer to **string** |  | [optional] 
**ThresholdValue** | Pointer to [**ThresholdInfoThresholdValue**](ThresholdInfoThresholdValue.md) |  | [optional] 
**Hysteresis** | Pointer to [**ThresholdInfoHysteresis**](ThresholdInfoHysteresis.md) |  | [optional] 

## Methods

### NewThresholdInfo

`func NewThresholdInfo() *ThresholdInfo`

NewThresholdInfo instantiates a new ThresholdInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdInfoWithDefaults

`func NewThresholdInfoWithDefaults() *ThresholdInfo`

NewThresholdInfoWithDefaults instantiates a new ThresholdInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThresholdDirection

`func (o *ThresholdInfo) GetThresholdDirection() string`

GetThresholdDirection returns the ThresholdDirection field if non-nil, zero value otherwise.

### GetThresholdDirectionOk

`func (o *ThresholdInfo) GetThresholdDirectionOk() (*string, bool)`

GetThresholdDirectionOk returns a tuple with the ThresholdDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdDirection

`func (o *ThresholdInfo) SetThresholdDirection(v string)`

SetThresholdDirection sets ThresholdDirection field to given value.

### HasThresholdDirection

`func (o *ThresholdInfo) HasThresholdDirection() bool`

HasThresholdDirection returns a boolean if a field has been set.

### GetThresholdValue

`func (o *ThresholdInfo) GetThresholdValue() ThresholdInfoThresholdValue`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *ThresholdInfo) GetThresholdValueOk() (*ThresholdInfoThresholdValue, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *ThresholdInfo) SetThresholdValue(v ThresholdInfoThresholdValue)`

SetThresholdValue sets ThresholdValue field to given value.

### HasThresholdValue

`func (o *ThresholdInfo) HasThresholdValue() bool`

HasThresholdValue returns a boolean if a field has been set.

### GetHysteresis

`func (o *ThresholdInfo) GetHysteresis() ThresholdInfoHysteresis`

GetHysteresis returns the Hysteresis field if non-nil, zero value otherwise.

### GetHysteresisOk

`func (o *ThresholdInfo) GetHysteresisOk() (*ThresholdInfoHysteresis, bool)`

GetHysteresisOk returns a tuple with the Hysteresis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHysteresis

`func (o *ThresholdInfo) SetHysteresis(v ThresholdInfoHysteresis)`

SetHysteresis sets Hysteresis field to given value.

### HasHysteresis

`func (o *ThresholdInfo) HasHysteresis() bool`

HasHysteresis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


