# UeContextCreatedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeContext** | [**UeContext**](UeContext.md) |  | 
**TargetToSourceData** | [**N2InfoContent**](N2InfoContent.md) |  | 
**PduSessionList** | [**[]N2SmInformation**](N2SmInformation.md) |  | 
**FailedSessionList** | Pointer to [**[]N2SmInformation**](N2SmInformation.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**PcfReselectedInd** | Pointer to **bool** |  | [optional] 
**AnalyticsNotUsedList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUeContextCreatedData

`func NewUeContextCreatedData(ueContext UeContext, targetToSourceData N2InfoContent, pduSessionList []N2SmInformation, ) *UeContextCreatedData`

NewUeContextCreatedData instantiates a new UeContextCreatedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeContextCreatedDataWithDefaults

`func NewUeContextCreatedDataWithDefaults() *UeContextCreatedData`

NewUeContextCreatedDataWithDefaults instantiates a new UeContextCreatedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeContext

`func (o *UeContextCreatedData) GetUeContext() UeContext`

GetUeContext returns the UeContext field if non-nil, zero value otherwise.

### GetUeContextOk

`func (o *UeContextCreatedData) GetUeContextOk() (*UeContext, bool)`

GetUeContextOk returns a tuple with the UeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeContext

`func (o *UeContextCreatedData) SetUeContext(v UeContext)`

SetUeContext sets UeContext field to given value.


### GetTargetToSourceData

`func (o *UeContextCreatedData) GetTargetToSourceData() N2InfoContent`

GetTargetToSourceData returns the TargetToSourceData field if non-nil, zero value otherwise.

### GetTargetToSourceDataOk

`func (o *UeContextCreatedData) GetTargetToSourceDataOk() (*N2InfoContent, bool)`

GetTargetToSourceDataOk returns a tuple with the TargetToSourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetToSourceData

`func (o *UeContextCreatedData) SetTargetToSourceData(v N2InfoContent)`

SetTargetToSourceData sets TargetToSourceData field to given value.


### GetPduSessionList

`func (o *UeContextCreatedData) GetPduSessionList() []N2SmInformation`

GetPduSessionList returns the PduSessionList field if non-nil, zero value otherwise.

### GetPduSessionListOk

`func (o *UeContextCreatedData) GetPduSessionListOk() (*[]N2SmInformation, bool)`

GetPduSessionListOk returns a tuple with the PduSessionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionList

`func (o *UeContextCreatedData) SetPduSessionList(v []N2SmInformation)`

SetPduSessionList sets PduSessionList field to given value.


### GetFailedSessionList

`func (o *UeContextCreatedData) GetFailedSessionList() []N2SmInformation`

GetFailedSessionList returns the FailedSessionList field if non-nil, zero value otherwise.

### GetFailedSessionListOk

`func (o *UeContextCreatedData) GetFailedSessionListOk() (*[]N2SmInformation, bool)`

GetFailedSessionListOk returns a tuple with the FailedSessionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSessionList

`func (o *UeContextCreatedData) SetFailedSessionList(v []N2SmInformation)`

SetFailedSessionList sets FailedSessionList field to given value.

### HasFailedSessionList

`func (o *UeContextCreatedData) HasFailedSessionList() bool`

HasFailedSessionList returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *UeContextCreatedData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *UeContextCreatedData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *UeContextCreatedData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *UeContextCreatedData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetPcfReselectedInd

`func (o *UeContextCreatedData) GetPcfReselectedInd() bool`

GetPcfReselectedInd returns the PcfReselectedInd field if non-nil, zero value otherwise.

### GetPcfReselectedIndOk

`func (o *UeContextCreatedData) GetPcfReselectedIndOk() (*bool, bool)`

GetPcfReselectedIndOk returns a tuple with the PcfReselectedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfReselectedInd

`func (o *UeContextCreatedData) SetPcfReselectedInd(v bool)`

SetPcfReselectedInd sets PcfReselectedInd field to given value.

### HasPcfReselectedInd

`func (o *UeContextCreatedData) HasPcfReselectedInd() bool`

HasPcfReselectedInd returns a boolean if a field has been set.

### GetAnalyticsNotUsedList

`func (o *UeContextCreatedData) GetAnalyticsNotUsedList() []string`

GetAnalyticsNotUsedList returns the AnalyticsNotUsedList field if non-nil, zero value otherwise.

### GetAnalyticsNotUsedListOk

`func (o *UeContextCreatedData) GetAnalyticsNotUsedListOk() (*[]string, bool)`

GetAnalyticsNotUsedListOk returns a tuple with the AnalyticsNotUsedList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsNotUsedList

`func (o *UeContextCreatedData) SetAnalyticsNotUsedList(v []string)`

SetAnalyticsNotUsedList sets AnalyticsNotUsedList field to given value.

### HasAnalyticsNotUsedList

`func (o *UeContextCreatedData) HasAnalyticsNotUsedList() bool`

HasAnalyticsNotUsedList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


