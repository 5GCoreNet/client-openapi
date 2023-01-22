# AccessTimeDistributionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supis** | Pointer to **[]string** |  | [optional] 
**Gpsis** | Pointer to **[]string** |  | [optional] 
**InterGrpId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**ExterGrpId** | Pointer to **string** | String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.   | [optional] 
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

### GetSupis

`func (o *AccessTimeDistributionData) GetSupis() []string`

GetSupis returns the Supis field if non-nil, zero value otherwise.

### GetSupisOk

`func (o *AccessTimeDistributionData) GetSupisOk() (*[]string, bool)`

GetSupisOk returns a tuple with the Supis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupis

`func (o *AccessTimeDistributionData) SetSupis(v []string)`

SetSupis sets Supis field to given value.

### HasSupis

`func (o *AccessTimeDistributionData) HasSupis() bool`

HasSupis returns a boolean if a field has been set.

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

### GetInterGrpId

`func (o *AccessTimeDistributionData) GetInterGrpId() string`

GetInterGrpId returns the InterGrpId field if non-nil, zero value otherwise.

### GetInterGrpIdOk

`func (o *AccessTimeDistributionData) GetInterGrpIdOk() (*string, bool)`

GetInterGrpIdOk returns a tuple with the InterGrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGrpId

`func (o *AccessTimeDistributionData) SetInterGrpId(v string)`

SetInterGrpId sets InterGrpId field to given value.

### HasInterGrpId

`func (o *AccessTimeDistributionData) HasInterGrpId() bool`

HasInterGrpId returns a boolean if a field has been set.

### GetExterGrpId

`func (o *AccessTimeDistributionData) GetExterGrpId() string`

GetExterGrpId returns the ExterGrpId field if non-nil, zero value otherwise.

### GetExterGrpIdOk

`func (o *AccessTimeDistributionData) GetExterGrpIdOk() (*string, bool)`

GetExterGrpIdOk returns a tuple with the ExterGrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterGrpId

`func (o *AccessTimeDistributionData) SetExterGrpId(v string)`

SetExterGrpId sets ExterGrpId field to given value.

### HasExterGrpId

`func (o *AccessTimeDistributionData) HasExterGrpId() bool`

HasExterGrpId returns a boolean if a field has been set.

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


