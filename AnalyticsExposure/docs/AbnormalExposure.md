# AbnormalExposure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsis** | Pointer to **[]string** |  | [optional] 
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**Excep** | [**Exception**](Exception.md) |  | 
**Ratio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.   | [optional] 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**AddtMeasInfo** | Pointer to [**AdditionalMeasurement**](AdditionalMeasurement.md) |  | [optional] 

## Methods

### NewAbnormalExposure

`func NewAbnormalExposure(excep Exception, ) *AbnormalExposure`

NewAbnormalExposure instantiates a new AbnormalExposure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbnormalExposureWithDefaults

`func NewAbnormalExposureWithDefaults() *AbnormalExposure`

NewAbnormalExposureWithDefaults instantiates a new AbnormalExposure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsis

`func (o *AbnormalExposure) GetGpsis() []string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *AbnormalExposure) GetGpsisOk() (*[]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *AbnormalExposure) SetGpsis(v []string)`

SetGpsis sets Gpsis field to given value.

### HasGpsis

`func (o *AbnormalExposure) HasGpsis() bool`

HasGpsis returns a boolean if a field has been set.

### GetAppId

`func (o *AbnormalExposure) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AbnormalExposure) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AbnormalExposure) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AbnormalExposure) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetExcep

`func (o *AbnormalExposure) GetExcep() Exception`

GetExcep returns the Excep field if non-nil, zero value otherwise.

### GetExcepOk

`func (o *AbnormalExposure) GetExcepOk() (*Exception, bool)`

GetExcepOk returns a tuple with the Excep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcep

`func (o *AbnormalExposure) SetExcep(v Exception)`

SetExcep sets Excep field to given value.


### GetRatio

`func (o *AbnormalExposure) GetRatio() int32`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *AbnormalExposure) GetRatioOk() (*int32, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *AbnormalExposure) SetRatio(v int32)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *AbnormalExposure) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetConfidence

`func (o *AbnormalExposure) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *AbnormalExposure) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *AbnormalExposure) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *AbnormalExposure) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetAddtMeasInfo

`func (o *AbnormalExposure) GetAddtMeasInfo() AdditionalMeasurement`

GetAddtMeasInfo returns the AddtMeasInfo field if non-nil, zero value otherwise.

### GetAddtMeasInfoOk

`func (o *AbnormalExposure) GetAddtMeasInfoOk() (*AdditionalMeasurement, bool)`

GetAddtMeasInfoOk returns a tuple with the AddtMeasInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddtMeasInfo

`func (o *AbnormalExposure) SetAddtMeasInfo(v AdditionalMeasurement)`

SetAddtMeasInfo sets AddtMeasInfo field to given value.

### HasAddtMeasInfo

`func (o *AbnormalExposure) HasAddtMeasInfo() bool`

HasAddtMeasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


