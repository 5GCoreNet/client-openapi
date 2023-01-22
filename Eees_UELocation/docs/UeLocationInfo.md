# UeLocationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Loc** | [**LocationArea5G**](LocationArea5G.md) |  | 
**Ratio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.   | [optional] 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewUeLocationInfo

`func NewUeLocationInfo(loc LocationArea5G, ) *UeLocationInfo`

NewUeLocationInfo instantiates a new UeLocationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeLocationInfoWithDefaults

`func NewUeLocationInfoWithDefaults() *UeLocationInfo`

NewUeLocationInfoWithDefaults instantiates a new UeLocationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoc

`func (o *UeLocationInfo) GetLoc() LocationArea5G`

GetLoc returns the Loc field if non-nil, zero value otherwise.

### GetLocOk

`func (o *UeLocationInfo) GetLocOk() (*LocationArea5G, bool)`

GetLocOk returns a tuple with the Loc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoc

`func (o *UeLocationInfo) SetLoc(v LocationArea5G)`

SetLoc sets Loc field to given value.


### GetRatio

`func (o *UeLocationInfo) GetRatio() int32`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *UeLocationInfo) GetRatioOk() (*int32, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *UeLocationInfo) SetRatio(v int32)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *UeLocationInfo) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetConfidence

`func (o *UeLocationInfo) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *UeLocationInfo) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *UeLocationInfo) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *UeLocationInfo) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


