# NetworkPerfExposure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocArea** | [**LocationArea5G**](LocationArea5G.md) |  | 
**NwPerfType** | [**NetworkPerfType**](NetworkPerfType.md) |  | 
**RelativeRatio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.   | [optional] 
**AbsoluteNum** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewNetworkPerfExposure

`func NewNetworkPerfExposure(locArea LocationArea5G, nwPerfType NetworkPerfType, ) *NetworkPerfExposure`

NewNetworkPerfExposure instantiates a new NetworkPerfExposure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPerfExposureWithDefaults

`func NewNetworkPerfExposureWithDefaults() *NetworkPerfExposure`

NewNetworkPerfExposureWithDefaults instantiates a new NetworkPerfExposure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocArea

`func (o *NetworkPerfExposure) GetLocArea() LocationArea5G`

GetLocArea returns the LocArea field if non-nil, zero value otherwise.

### GetLocAreaOk

`func (o *NetworkPerfExposure) GetLocAreaOk() (*LocationArea5G, bool)`

GetLocAreaOk returns a tuple with the LocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocArea

`func (o *NetworkPerfExposure) SetLocArea(v LocationArea5G)`

SetLocArea sets LocArea field to given value.


### GetNwPerfType

`func (o *NetworkPerfExposure) GetNwPerfType() NetworkPerfType`

GetNwPerfType returns the NwPerfType field if non-nil, zero value otherwise.

### GetNwPerfTypeOk

`func (o *NetworkPerfExposure) GetNwPerfTypeOk() (*NetworkPerfType, bool)`

GetNwPerfTypeOk returns a tuple with the NwPerfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfType

`func (o *NetworkPerfExposure) SetNwPerfType(v NetworkPerfType)`

SetNwPerfType sets NwPerfType field to given value.


### GetRelativeRatio

`func (o *NetworkPerfExposure) GetRelativeRatio() int32`

GetRelativeRatio returns the RelativeRatio field if non-nil, zero value otherwise.

### GetRelativeRatioOk

`func (o *NetworkPerfExposure) GetRelativeRatioOk() (*int32, bool)`

GetRelativeRatioOk returns a tuple with the RelativeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeRatio

`func (o *NetworkPerfExposure) SetRelativeRatio(v int32)`

SetRelativeRatio sets RelativeRatio field to given value.

### HasRelativeRatio

`func (o *NetworkPerfExposure) HasRelativeRatio() bool`

HasRelativeRatio returns a boolean if a field has been set.

### GetAbsoluteNum

`func (o *NetworkPerfExposure) GetAbsoluteNum() int32`

GetAbsoluteNum returns the AbsoluteNum field if non-nil, zero value otherwise.

### GetAbsoluteNumOk

`func (o *NetworkPerfExposure) GetAbsoluteNumOk() (*int32, bool)`

GetAbsoluteNumOk returns a tuple with the AbsoluteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteNum

`func (o *NetworkPerfExposure) SetAbsoluteNum(v int32)`

SetAbsoluteNum sets AbsoluteNum field to given value.

### HasAbsoluteNum

`func (o *NetworkPerfExposure) HasAbsoluteNum() bool`

HasAbsoluteNum returns a boolean if a field has been set.

### GetConfidence

`func (o *NetworkPerfExposure) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *NetworkPerfExposure) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *NetworkPerfExposure) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *NetworkPerfExposure) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


