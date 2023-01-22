# LocationReportConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValServerId** | **string** |  | 
**ValTgtUe** | [**ValTargetUe**](ValTargetUe.md) |  | 
**ImmRep** | Pointer to **bool** |  | [optional] 
**MonDur** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RepPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**Accuracy** | Pointer to [**Accuracy**](Accuracy.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewLocationReportConfiguration

`func NewLocationReportConfiguration(valServerId string, valTgtUe ValTargetUe, ) *LocationReportConfiguration`

NewLocationReportConfiguration instantiates a new LocationReportConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationReportConfigurationWithDefaults

`func NewLocationReportConfigurationWithDefaults() *LocationReportConfiguration`

NewLocationReportConfigurationWithDefaults instantiates a new LocationReportConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValServerId

`func (o *LocationReportConfiguration) GetValServerId() string`

GetValServerId returns the ValServerId field if non-nil, zero value otherwise.

### GetValServerIdOk

`func (o *LocationReportConfiguration) GetValServerIdOk() (*string, bool)`

GetValServerIdOk returns a tuple with the ValServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValServerId

`func (o *LocationReportConfiguration) SetValServerId(v string)`

SetValServerId sets ValServerId field to given value.


### GetValTgtUe

`func (o *LocationReportConfiguration) GetValTgtUe() ValTargetUe`

GetValTgtUe returns the ValTgtUe field if non-nil, zero value otherwise.

### GetValTgtUeOk

`func (o *LocationReportConfiguration) GetValTgtUeOk() (*ValTargetUe, bool)`

GetValTgtUeOk returns a tuple with the ValTgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValTgtUe

`func (o *LocationReportConfiguration) SetValTgtUe(v ValTargetUe)`

SetValTgtUe sets ValTgtUe field to given value.


### GetImmRep

`func (o *LocationReportConfiguration) GetImmRep() bool`

GetImmRep returns the ImmRep field if non-nil, zero value otherwise.

### GetImmRepOk

`func (o *LocationReportConfiguration) GetImmRepOk() (*bool, bool)`

GetImmRepOk returns a tuple with the ImmRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmRep

`func (o *LocationReportConfiguration) SetImmRep(v bool)`

SetImmRep sets ImmRep field to given value.

### HasImmRep

`func (o *LocationReportConfiguration) HasImmRep() bool`

HasImmRep returns a boolean if a field has been set.

### GetMonDur

`func (o *LocationReportConfiguration) GetMonDur() time.Time`

GetMonDur returns the MonDur field if non-nil, zero value otherwise.

### GetMonDurOk

`func (o *LocationReportConfiguration) GetMonDurOk() (*time.Time, bool)`

GetMonDurOk returns a tuple with the MonDur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonDur

`func (o *LocationReportConfiguration) SetMonDur(v time.Time)`

SetMonDur sets MonDur field to given value.

### HasMonDur

`func (o *LocationReportConfiguration) HasMonDur() bool`

HasMonDur returns a boolean if a field has been set.

### GetRepPeriod

`func (o *LocationReportConfiguration) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *LocationReportConfiguration) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *LocationReportConfiguration) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *LocationReportConfiguration) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.

### GetAccuracy

`func (o *LocationReportConfiguration) GetAccuracy() Accuracy`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *LocationReportConfiguration) GetAccuracyOk() (*Accuracy, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *LocationReportConfiguration) SetAccuracy(v Accuracy)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *LocationReportConfiguration) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### GetSuppFeat

`func (o *LocationReportConfiguration) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *LocationReportConfiguration) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *LocationReportConfiguration) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *LocationReportConfiguration) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


