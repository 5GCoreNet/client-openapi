# SmPolicyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmPolicySnssaiData** | [**map[string]SmPolicySnssaiData**](SmPolicySnssaiData.md) | Contains Session Management Policy data per S-NSSAI for all the SNSSAIs of the subscriber. The key of the map is the S-NSSAI.  | 
**UmDataLimits** | Pointer to [**map[string]UsageMonDataLimit**](UsageMonDataLimit.md) | Contains a list of usage monitoring profiles associated with the subscriber. The limit identifier is used as the key of the map.  | [optional] 
**UmData** | Pointer to [**map[string]UsageMonData**](UsageMonData.md) | Contains the remaining allowed usage data associated with the subscriber. The limit identifier is used as the key of the map.  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSmPolicyData

`func NewSmPolicyData(smPolicySnssaiData map[string]SmPolicySnssaiData, ) *SmPolicyData`

NewSmPolicyData instantiates a new SmPolicyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmPolicyDataWithDefaults

`func NewSmPolicyDataWithDefaults() *SmPolicyData`

NewSmPolicyDataWithDefaults instantiates a new SmPolicyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmPolicySnssaiData

`func (o *SmPolicyData) GetSmPolicySnssaiData() map[string]SmPolicySnssaiData`

GetSmPolicySnssaiData returns the SmPolicySnssaiData field if non-nil, zero value otherwise.

### GetSmPolicySnssaiDataOk

`func (o *SmPolicyData) GetSmPolicySnssaiDataOk() (*map[string]SmPolicySnssaiData, bool)`

GetSmPolicySnssaiDataOk returns a tuple with the SmPolicySnssaiData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmPolicySnssaiData

`func (o *SmPolicyData) SetSmPolicySnssaiData(v map[string]SmPolicySnssaiData)`

SetSmPolicySnssaiData sets SmPolicySnssaiData field to given value.


### GetUmDataLimits

`func (o *SmPolicyData) GetUmDataLimits() map[string]UsageMonDataLimit`

GetUmDataLimits returns the UmDataLimits field if non-nil, zero value otherwise.

### GetUmDataLimitsOk

`func (o *SmPolicyData) GetUmDataLimitsOk() (*map[string]UsageMonDataLimit, bool)`

GetUmDataLimitsOk returns a tuple with the UmDataLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmDataLimits

`func (o *SmPolicyData) SetUmDataLimits(v map[string]UsageMonDataLimit)`

SetUmDataLimits sets UmDataLimits field to given value.

### HasUmDataLimits

`func (o *SmPolicyData) HasUmDataLimits() bool`

HasUmDataLimits returns a boolean if a field has been set.

### GetUmData

`func (o *SmPolicyData) GetUmData() map[string]UsageMonData`

GetUmData returns the UmData field if non-nil, zero value otherwise.

### GetUmDataOk

`func (o *SmPolicyData) GetUmDataOk() (*map[string]UsageMonData, bool)`

GetUmDataOk returns a tuple with the UmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmData

`func (o *SmPolicyData) SetUmData(v map[string]UsageMonData)`

SetUmData sets UmData field to given value.

### HasUmData

`func (o *SmPolicyData) HasUmData() bool`

HasUmData returns a boolean if a field has been set.

### GetSuppFeat

`func (o *SmPolicyData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *SmPolicyData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *SmPolicyData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *SmPolicyData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetResetIds

`func (o *SmPolicyData) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *SmPolicyData) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *SmPolicyData) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *SmPolicyData) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


