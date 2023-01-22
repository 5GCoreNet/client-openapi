# SlicePolicyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbrUl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**MbrDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**RemainMbrUl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**RemainMbrDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSlicePolicyData

`func NewSlicePolicyData() *SlicePolicyData`

NewSlicePolicyData instantiates a new SlicePolicyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlicePolicyDataWithDefaults

`func NewSlicePolicyDataWithDefaults() *SlicePolicyData`

NewSlicePolicyDataWithDefaults instantiates a new SlicePolicyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbrUl

`func (o *SlicePolicyData) GetMbrUl() string`

GetMbrUl returns the MbrUl field if non-nil, zero value otherwise.

### GetMbrUlOk

`func (o *SlicePolicyData) GetMbrUlOk() (*string, bool)`

GetMbrUlOk returns a tuple with the MbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbrUl

`func (o *SlicePolicyData) SetMbrUl(v string)`

SetMbrUl sets MbrUl field to given value.

### HasMbrUl

`func (o *SlicePolicyData) HasMbrUl() bool`

HasMbrUl returns a boolean if a field has been set.

### GetMbrDl

`func (o *SlicePolicyData) GetMbrDl() string`

GetMbrDl returns the MbrDl field if non-nil, zero value otherwise.

### GetMbrDlOk

`func (o *SlicePolicyData) GetMbrDlOk() (*string, bool)`

GetMbrDlOk returns a tuple with the MbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbrDl

`func (o *SlicePolicyData) SetMbrDl(v string)`

SetMbrDl sets MbrDl field to given value.

### HasMbrDl

`func (o *SlicePolicyData) HasMbrDl() bool`

HasMbrDl returns a boolean if a field has been set.

### GetRemainMbrUl

`func (o *SlicePolicyData) GetRemainMbrUl() string`

GetRemainMbrUl returns the RemainMbrUl field if non-nil, zero value otherwise.

### GetRemainMbrUlOk

`func (o *SlicePolicyData) GetRemainMbrUlOk() (*string, bool)`

GetRemainMbrUlOk returns a tuple with the RemainMbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainMbrUl

`func (o *SlicePolicyData) SetRemainMbrUl(v string)`

SetRemainMbrUl sets RemainMbrUl field to given value.

### HasRemainMbrUl

`func (o *SlicePolicyData) HasRemainMbrUl() bool`

HasRemainMbrUl returns a boolean if a field has been set.

### GetRemainMbrDl

`func (o *SlicePolicyData) GetRemainMbrDl() string`

GetRemainMbrDl returns the RemainMbrDl field if non-nil, zero value otherwise.

### GetRemainMbrDlOk

`func (o *SlicePolicyData) GetRemainMbrDlOk() (*string, bool)`

GetRemainMbrDlOk returns a tuple with the RemainMbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainMbrDl

`func (o *SlicePolicyData) SetRemainMbrDl(v string)`

SetRemainMbrDl sets RemainMbrDl field to given value.

### HasRemainMbrDl

`func (o *SlicePolicyData) HasRemainMbrDl() bool`

HasRemainMbrDl returns a boolean if a field has been set.

### GetSuppFeat

`func (o *SlicePolicyData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *SlicePolicyData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *SlicePolicyData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *SlicePolicyData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetResetIds

`func (o *SlicePolicyData) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *SlicePolicyData) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *SlicePolicyData) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *SlicePolicyData) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


