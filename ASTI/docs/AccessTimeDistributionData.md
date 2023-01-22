# AccessTimeDistributionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsis** | Pointer to **[]string** |  | [optional] 
**ExterGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**AsTimeDisParam** | [**AsTimeDistributionParam**](AsTimeDistributionParam.md) |  | 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewAccessTimeDistributionData

`func NewAccessTimeDistributionData(asTimeDisParam AsTimeDistributionParam, ) *AccessTimeDistributionData`

NewAccessTimeDistributionData instantiates a new AccessTimeDistributionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTimeDistributionDataWithDefaults

`func NewAccessTimeDistributionDataWithDefaults() *AccessTimeDistributionData`

NewAccessTimeDistributionDataWithDefaults instantiates a new AccessTimeDistributionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsis

`func (o *AccessTimeDistributionData) GetGpsis() []string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *AccessTimeDistributionData) GetGpsisOk() (*[]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *AccessTimeDistributionData) SetGpsis(v []string)`

SetGpsis sets Gpsis field to given value.

### HasGpsis

`func (o *AccessTimeDistributionData) HasGpsis() bool`

HasGpsis returns a boolean if a field has been set.

### GetExterGroupId

`func (o *AccessTimeDistributionData) GetExterGroupId() string`

GetExterGroupId returns the ExterGroupId field if non-nil, zero value otherwise.

### GetExterGroupIdOk

`func (o *AccessTimeDistributionData) GetExterGroupIdOk() (*string, bool)`

GetExterGroupIdOk returns a tuple with the ExterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterGroupId

`func (o *AccessTimeDistributionData) SetExterGroupId(v string)`

SetExterGroupId sets ExterGroupId field to given value.

### HasExterGroupId

`func (o *AccessTimeDistributionData) HasExterGroupId() bool`

HasExterGroupId returns a boolean if a field has been set.

### GetAsTimeDisParam

`func (o *AccessTimeDistributionData) GetAsTimeDisParam() AsTimeDistributionParam`

GetAsTimeDisParam returns the AsTimeDisParam field if non-nil, zero value otherwise.

### GetAsTimeDisParamOk

`func (o *AccessTimeDistributionData) GetAsTimeDisParamOk() (*AsTimeDistributionParam, bool)`

GetAsTimeDisParamOk returns a tuple with the AsTimeDisParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTimeDisParam

`func (o *AccessTimeDistributionData) SetAsTimeDisParam(v AsTimeDistributionParam)`

SetAsTimeDisParam sets AsTimeDisParam field to given value.


### GetSuppFeat

`func (o *AccessTimeDistributionData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AccessTimeDistributionData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AccessTimeDistributionData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AccessTimeDistributionData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


