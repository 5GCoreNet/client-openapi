# AsTimeDistributionParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsTimeDisEnabled** | Pointer to **bool** | When this attribute is included and set to true, it indicates that the access stratum time distribution via Uu reference point is activated.  | [optional] 
**TimeSyncErrBdgt** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**TempValidity** | Pointer to [**TemporalValidity**](TemporalValidity.md) |  | [optional] 

## Methods

### NewAsTimeDistributionParam

`func NewAsTimeDistributionParam() *AsTimeDistributionParam`

NewAsTimeDistributionParam instantiates a new AsTimeDistributionParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsTimeDistributionParamWithDefaults

`func NewAsTimeDistributionParamWithDefaults() *AsTimeDistributionParam`

NewAsTimeDistributionParamWithDefaults instantiates a new AsTimeDistributionParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsTimeDisEnabled

`func (o *AsTimeDistributionParam) GetAsTimeDisEnabled() bool`

GetAsTimeDisEnabled returns the AsTimeDisEnabled field if non-nil, zero value otherwise.

### GetAsTimeDisEnabledOk

`func (o *AsTimeDistributionParam) GetAsTimeDisEnabledOk() (*bool, bool)`

GetAsTimeDisEnabledOk returns a tuple with the AsTimeDisEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTimeDisEnabled

`func (o *AsTimeDistributionParam) SetAsTimeDisEnabled(v bool)`

SetAsTimeDisEnabled sets AsTimeDisEnabled field to given value.

### HasAsTimeDisEnabled

`func (o *AsTimeDistributionParam) HasAsTimeDisEnabled() bool`

HasAsTimeDisEnabled returns a boolean if a field has been set.

### GetTimeSyncErrBdgt

`func (o *AsTimeDistributionParam) GetTimeSyncErrBdgt() int32`

GetTimeSyncErrBdgt returns the TimeSyncErrBdgt field if non-nil, zero value otherwise.

### GetTimeSyncErrBdgtOk

`func (o *AsTimeDistributionParam) GetTimeSyncErrBdgtOk() (*int32, bool)`

GetTimeSyncErrBdgtOk returns a tuple with the TimeSyncErrBdgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSyncErrBdgt

`func (o *AsTimeDistributionParam) SetTimeSyncErrBdgt(v int32)`

SetTimeSyncErrBdgt sets TimeSyncErrBdgt field to given value.

### HasTimeSyncErrBdgt

`func (o *AsTimeDistributionParam) HasTimeSyncErrBdgt() bool`

HasTimeSyncErrBdgt returns a boolean if a field has been set.

### GetTempValidity

`func (o *AsTimeDistributionParam) GetTempValidity() TemporalValidity`

GetTempValidity returns the TempValidity field if non-nil, zero value otherwise.

### GetTempValidityOk

`func (o *AsTimeDistributionParam) GetTempValidityOk() (*TemporalValidity, bool)`

GetTempValidityOk returns a tuple with the TempValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempValidity

`func (o *AsTimeDistributionParam) SetTempValidity(v TemporalValidity)`

SetTempValidity sets TempValidity field to given value.

### HasTempValidity

`func (o *AsTimeDistributionParam) HasTempValidity() bool`

HasTempValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


