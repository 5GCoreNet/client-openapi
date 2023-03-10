# RedundantTransmissionExpPerTS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TsStart** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**TsDuration** | **int32** | indicating a time in seconds. | 
**ObsvRedTransExp** | [**ObservedRedundantTransExp**](ObservedRedundantTransExp.md) |  | 
**RedTransStatus** | Pointer to **bool** | Redundant Transmission Status. Set to \&quot;true\&quot; if redundant transmission was activated, otherwise set to \&quot;false\&quot;. Default value is \&quot;false\&quot; if omitted.  | [optional] 
**UeRatio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.   | [optional] 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewRedundantTransmissionExpPerTS

`func NewRedundantTransmissionExpPerTS(tsStart time.Time, tsDuration int32, obsvRedTransExp ObservedRedundantTransExp, ) *RedundantTransmissionExpPerTS`

NewRedundantTransmissionExpPerTS instantiates a new RedundantTransmissionExpPerTS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedundantTransmissionExpPerTSWithDefaults

`func NewRedundantTransmissionExpPerTSWithDefaults() *RedundantTransmissionExpPerTS`

NewRedundantTransmissionExpPerTSWithDefaults instantiates a new RedundantTransmissionExpPerTS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTsStart

`func (o *RedundantTransmissionExpPerTS) GetTsStart() time.Time`

GetTsStart returns the TsStart field if non-nil, zero value otherwise.

### GetTsStartOk

`func (o *RedundantTransmissionExpPerTS) GetTsStartOk() (*time.Time, bool)`

GetTsStartOk returns a tuple with the TsStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsStart

`func (o *RedundantTransmissionExpPerTS) SetTsStart(v time.Time)`

SetTsStart sets TsStart field to given value.


### GetTsDuration

`func (o *RedundantTransmissionExpPerTS) GetTsDuration() int32`

GetTsDuration returns the TsDuration field if non-nil, zero value otherwise.

### GetTsDurationOk

`func (o *RedundantTransmissionExpPerTS) GetTsDurationOk() (*int32, bool)`

GetTsDurationOk returns a tuple with the TsDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsDuration

`func (o *RedundantTransmissionExpPerTS) SetTsDuration(v int32)`

SetTsDuration sets TsDuration field to given value.


### GetObsvRedTransExp

`func (o *RedundantTransmissionExpPerTS) GetObsvRedTransExp() ObservedRedundantTransExp`

GetObsvRedTransExp returns the ObsvRedTransExp field if non-nil, zero value otherwise.

### GetObsvRedTransExpOk

`func (o *RedundantTransmissionExpPerTS) GetObsvRedTransExpOk() (*ObservedRedundantTransExp, bool)`

GetObsvRedTransExpOk returns a tuple with the ObsvRedTransExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsvRedTransExp

`func (o *RedundantTransmissionExpPerTS) SetObsvRedTransExp(v ObservedRedundantTransExp)`

SetObsvRedTransExp sets ObsvRedTransExp field to given value.


### GetRedTransStatus

`func (o *RedundantTransmissionExpPerTS) GetRedTransStatus() bool`

GetRedTransStatus returns the RedTransStatus field if non-nil, zero value otherwise.

### GetRedTransStatusOk

`func (o *RedundantTransmissionExpPerTS) GetRedTransStatusOk() (*bool, bool)`

GetRedTransStatusOk returns a tuple with the RedTransStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedTransStatus

`func (o *RedundantTransmissionExpPerTS) SetRedTransStatus(v bool)`

SetRedTransStatus sets RedTransStatus field to given value.

### HasRedTransStatus

`func (o *RedundantTransmissionExpPerTS) HasRedTransStatus() bool`

HasRedTransStatus returns a boolean if a field has been set.

### GetUeRatio

`func (o *RedundantTransmissionExpPerTS) GetUeRatio() int32`

GetUeRatio returns the UeRatio field if non-nil, zero value otherwise.

### GetUeRatioOk

`func (o *RedundantTransmissionExpPerTS) GetUeRatioOk() (*int32, bool)`

GetUeRatioOk returns a tuple with the UeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRatio

`func (o *RedundantTransmissionExpPerTS) SetUeRatio(v int32)`

SetUeRatio sets UeRatio field to given value.

### HasUeRatio

`func (o *RedundantTransmissionExpPerTS) HasUeRatio() bool`

HasUeRatio returns a boolean if a field has been set.

### GetConfidence

`func (o *RedundantTransmissionExpPerTS) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *RedundantTransmissionExpPerTS) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *RedundantTransmissionExpPerTS) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *RedundantTransmissionExpPerTS) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


