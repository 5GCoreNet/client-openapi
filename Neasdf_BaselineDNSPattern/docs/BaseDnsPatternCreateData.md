# BaseDnsPatternCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**BaseDnsMdtList** | Pointer to [**map[string]BaselineDnsMdt**](BaselineDnsMdt.md) | map of baseline DNS message detection templates where a valid JSON string serves as key | [optional] 
**BaseDnsAitList** | Pointer to [**map[string]BaselineDnsAit**](BaselineDnsAit.md) | map of Baseline DNS action information Template where a valid JSON string serves as key | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewBaseDnsPatternCreateData

`func NewBaseDnsPatternCreateData() *BaseDnsPatternCreateData`

NewBaseDnsPatternCreateData instantiates a new BaseDnsPatternCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseDnsPatternCreateDataWithDefaults

`func NewBaseDnsPatternCreateDataWithDefaults() *BaseDnsPatternCreateData`

NewBaseDnsPatternCreateDataWithDefaults instantiates a new BaseDnsPatternCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *BaseDnsPatternCreateData) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BaseDnsPatternCreateData) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BaseDnsPatternCreateData) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BaseDnsPatternCreateData) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetBaseDnsMdtList

`func (o *BaseDnsPatternCreateData) GetBaseDnsMdtList() map[string]BaselineDnsMdt`

GetBaseDnsMdtList returns the BaseDnsMdtList field if non-nil, zero value otherwise.

### GetBaseDnsMdtListOk

`func (o *BaseDnsPatternCreateData) GetBaseDnsMdtListOk() (*map[string]BaselineDnsMdt, bool)`

GetBaseDnsMdtListOk returns a tuple with the BaseDnsMdtList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDnsMdtList

`func (o *BaseDnsPatternCreateData) SetBaseDnsMdtList(v map[string]BaselineDnsMdt)`

SetBaseDnsMdtList sets BaseDnsMdtList field to given value.

### HasBaseDnsMdtList

`func (o *BaseDnsPatternCreateData) HasBaseDnsMdtList() bool`

HasBaseDnsMdtList returns a boolean if a field has been set.

### GetBaseDnsAitList

`func (o *BaseDnsPatternCreateData) GetBaseDnsAitList() map[string]BaselineDnsAit`

GetBaseDnsAitList returns the BaseDnsAitList field if non-nil, zero value otherwise.

### GetBaseDnsAitListOk

`func (o *BaseDnsPatternCreateData) GetBaseDnsAitListOk() (*map[string]BaselineDnsAit, bool)`

GetBaseDnsAitListOk returns a tuple with the BaseDnsAitList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDnsAitList

`func (o *BaseDnsPatternCreateData) SetBaseDnsAitList(v map[string]BaselineDnsAit)`

SetBaseDnsAitList sets BaseDnsAitList field to given value.

### HasBaseDnsAitList

`func (o *BaseDnsPatternCreateData) HasBaseDnsAitList() bool`

HasBaseDnsAitList returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *BaseDnsPatternCreateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *BaseDnsPatternCreateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *BaseDnsPatternCreateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *BaseDnsPatternCreateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


