# FailureCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BssgpCause** | Pointer to **int32** | Identifies a non-transparent copy of the BSSGP cause code. Refer to 3GPP TS 29.128. | [optional] 
**CauseType** | Pointer to **int32** | Identify the type of the S1AP-Cause. Refer to 3GPP TS 29.128. | [optional] 
**GmmCause** | Pointer to **int32** | Identifies a non-transparent copy of the GMM cause code. Refer to 3GPP TS 29.128. | [optional] 
**RanapCause** | Pointer to **int32** | Identifies a non-transparent copy of the RANAP cause code. Refer to 3GPP TS 29.128. | [optional] 
**RanNasCause** | Pointer to **string** | Indicates RAN and/or NAS release cause code information, TWAN release cause code information or untrusted WLAN release cause code information. Refer to 3GPP TS 29.214. | [optional] 
**S1ApCause** | Pointer to **int32** | Identifies a non-transparent copy of the S1AP cause code. Refer to 3GPP TS 29.128. | [optional] 
**SmCause** | Pointer to **int32** | Identifies a non-transparent copy of the SM cause code. Refer to 3GPP TS 29.128. | [optional] 

## Methods

### NewFailureCause

`func NewFailureCause() *FailureCause`

NewFailureCause instantiates a new FailureCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureCauseWithDefaults

`func NewFailureCauseWithDefaults() *FailureCause`

NewFailureCauseWithDefaults instantiates a new FailureCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBssgpCause

`func (o *FailureCause) GetBssgpCause() int32`

GetBssgpCause returns the BssgpCause field if non-nil, zero value otherwise.

### GetBssgpCauseOk

`func (o *FailureCause) GetBssgpCauseOk() (*int32, bool)`

GetBssgpCauseOk returns a tuple with the BssgpCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssgpCause

`func (o *FailureCause) SetBssgpCause(v int32)`

SetBssgpCause sets BssgpCause field to given value.

### HasBssgpCause

`func (o *FailureCause) HasBssgpCause() bool`

HasBssgpCause returns a boolean if a field has been set.

### GetCauseType

`func (o *FailureCause) GetCauseType() int32`

GetCauseType returns the CauseType field if non-nil, zero value otherwise.

### GetCauseTypeOk

`func (o *FailureCause) GetCauseTypeOk() (*int32, bool)`

GetCauseTypeOk returns a tuple with the CauseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseType

`func (o *FailureCause) SetCauseType(v int32)`

SetCauseType sets CauseType field to given value.

### HasCauseType

`func (o *FailureCause) HasCauseType() bool`

HasCauseType returns a boolean if a field has been set.

### GetGmmCause

`func (o *FailureCause) GetGmmCause() int32`

GetGmmCause returns the GmmCause field if non-nil, zero value otherwise.

### GetGmmCauseOk

`func (o *FailureCause) GetGmmCauseOk() (*int32, bool)`

GetGmmCauseOk returns a tuple with the GmmCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmmCause

`func (o *FailureCause) SetGmmCause(v int32)`

SetGmmCause sets GmmCause field to given value.

### HasGmmCause

`func (o *FailureCause) HasGmmCause() bool`

HasGmmCause returns a boolean if a field has been set.

### GetRanapCause

`func (o *FailureCause) GetRanapCause() int32`

GetRanapCause returns the RanapCause field if non-nil, zero value otherwise.

### GetRanapCauseOk

`func (o *FailureCause) GetRanapCauseOk() (*int32, bool)`

GetRanapCauseOk returns a tuple with the RanapCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanapCause

`func (o *FailureCause) SetRanapCause(v int32)`

SetRanapCause sets RanapCause field to given value.

### HasRanapCause

`func (o *FailureCause) HasRanapCause() bool`

HasRanapCause returns a boolean if a field has been set.

### GetRanNasCause

`func (o *FailureCause) GetRanNasCause() string`

GetRanNasCause returns the RanNasCause field if non-nil, zero value otherwise.

### GetRanNasCauseOk

`func (o *FailureCause) GetRanNasCauseOk() (*string, bool)`

GetRanNasCauseOk returns a tuple with the RanNasCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanNasCause

`func (o *FailureCause) SetRanNasCause(v string)`

SetRanNasCause sets RanNasCause field to given value.

### HasRanNasCause

`func (o *FailureCause) HasRanNasCause() bool`

HasRanNasCause returns a boolean if a field has been set.

### GetS1ApCause

`func (o *FailureCause) GetS1ApCause() int32`

GetS1ApCause returns the S1ApCause field if non-nil, zero value otherwise.

### GetS1ApCauseOk

`func (o *FailureCause) GetS1ApCauseOk() (*int32, bool)`

GetS1ApCauseOk returns a tuple with the S1ApCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS1ApCause

`func (o *FailureCause) SetS1ApCause(v int32)`

SetS1ApCause sets S1ApCause field to given value.

### HasS1ApCause

`func (o *FailureCause) HasS1ApCause() bool`

HasS1ApCause returns a boolean if a field has been set.

### GetSmCause

`func (o *FailureCause) GetSmCause() int32`

GetSmCause returns the SmCause field if non-nil, zero value otherwise.

### GetSmCauseOk

`func (o *FailureCause) GetSmCauseOk() (*int32, bool)`

GetSmCauseOk returns a tuple with the SmCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmCause

`func (o *FailureCause) SetSmCause(v int32)`

SetSmCause sets SmCause field to given value.

### HasSmCause

`func (o *FailureCause) HasSmCause() bool`

HasSmCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


