# RatFreqInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFreq** | Pointer to **bool** | Set to \&quot;true\&quot; to indicate to handle all the frequencies the NWDAF received, otherwise  set to \&quot;false\&quot; or omit. The \&quot;allFreq\&quot; attribute and the \&quot;freq\&quot; attribute are mutually  exclusive.  | [optional] 
**AllRat** | Pointer to **bool** | Set to \&quot;true\&quot; to indicate to handle all the RAT Types the NWDAF received, otherwise  set to \&quot;false\&quot; or omit. The \&quot;allRat\&quot; attribute and the \&quot;ratType\&quot; attribute are mutually  exclusive.  | [optional] 
**Freq** | Pointer to **int32** | Integer value indicating the ARFCN applicable for a downlink, uplink or bi-directional (TDD) NR global frequency raster, as definition of \&quot;ARFCN-ValueNR\&quot; IE in clause 6.3.2 of 3GPP TS 38.331.  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**SvcExpThreshold** | Pointer to [**ThresholdLevel**](ThresholdLevel.md) |  | [optional] 
**MatchingDir** | Pointer to [**MatchingDirection**](MatchingDirection.md) |  | [optional] 

## Methods

### NewRatFreqInformation

`func NewRatFreqInformation() *RatFreqInformation`

NewRatFreqInformation instantiates a new RatFreqInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatFreqInformationWithDefaults

`func NewRatFreqInformationWithDefaults() *RatFreqInformation`

NewRatFreqInformationWithDefaults instantiates a new RatFreqInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFreq

`func (o *RatFreqInformation) GetAllFreq() bool`

GetAllFreq returns the AllFreq field if non-nil, zero value otherwise.

### GetAllFreqOk

`func (o *RatFreqInformation) GetAllFreqOk() (*bool, bool)`

GetAllFreqOk returns a tuple with the AllFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFreq

`func (o *RatFreqInformation) SetAllFreq(v bool)`

SetAllFreq sets AllFreq field to given value.

### HasAllFreq

`func (o *RatFreqInformation) HasAllFreq() bool`

HasAllFreq returns a boolean if a field has been set.

### GetAllRat

`func (o *RatFreqInformation) GetAllRat() bool`

GetAllRat returns the AllRat field if non-nil, zero value otherwise.

### GetAllRatOk

`func (o *RatFreqInformation) GetAllRatOk() (*bool, bool)`

GetAllRatOk returns a tuple with the AllRat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRat

`func (o *RatFreqInformation) SetAllRat(v bool)`

SetAllRat sets AllRat field to given value.

### HasAllRat

`func (o *RatFreqInformation) HasAllRat() bool`

HasAllRat returns a boolean if a field has been set.

### GetFreq

`func (o *RatFreqInformation) GetFreq() int32`

GetFreq returns the Freq field if non-nil, zero value otherwise.

### GetFreqOk

`func (o *RatFreqInformation) GetFreqOk() (*int32, bool)`

GetFreqOk returns a tuple with the Freq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreq

`func (o *RatFreqInformation) SetFreq(v int32)`

SetFreq sets Freq field to given value.

### HasFreq

`func (o *RatFreqInformation) HasFreq() bool`

HasFreq returns a boolean if a field has been set.

### GetRatType

`func (o *RatFreqInformation) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *RatFreqInformation) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *RatFreqInformation) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *RatFreqInformation) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetSvcExpThreshold

`func (o *RatFreqInformation) GetSvcExpThreshold() ThresholdLevel`

GetSvcExpThreshold returns the SvcExpThreshold field if non-nil, zero value otherwise.

### GetSvcExpThresholdOk

`func (o *RatFreqInformation) GetSvcExpThresholdOk() (*ThresholdLevel, bool)`

GetSvcExpThresholdOk returns a tuple with the SvcExpThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExpThreshold

`func (o *RatFreqInformation) SetSvcExpThreshold(v ThresholdLevel)`

SetSvcExpThreshold sets SvcExpThreshold field to given value.

### HasSvcExpThreshold

`func (o *RatFreqInformation) HasSvcExpThreshold() bool`

HasSvcExpThreshold returns a boolean if a field has been set.

### GetMatchingDir

`func (o *RatFreqInformation) GetMatchingDir() MatchingDirection`

GetMatchingDir returns the MatchingDir field if non-nil, zero value otherwise.

### GetMatchingDirOk

`func (o *RatFreqInformation) GetMatchingDirOk() (*MatchingDirection, bool)`

GetMatchingDirOk returns a tuple with the MatchingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingDir

`func (o *RatFreqInformation) SetMatchingDir(v MatchingDirection)`

SetMatchingDir sets MatchingDir field to given value.

### HasMatchingDir

`func (o *RatFreqInformation) HasMatchingDir() bool`

HasMatchingDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


