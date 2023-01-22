# N2InformationTransferReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**RatSelector** | Pointer to [**RatSelector**](RatSelector.md) |  | [optional] 
**GlobalRanNodeList** | Pointer to [**[]GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 
**N2Information** | [**N2InfoContainer**](N2InfoContainer.md) |  | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewN2InformationTransferReqData

`func NewN2InformationTransferReqData(n2Information N2InfoContainer, ) *N2InformationTransferReqData`

NewN2InformationTransferReqData instantiates a new N2InformationTransferReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN2InformationTransferReqDataWithDefaults

`func NewN2InformationTransferReqDataWithDefaults() *N2InformationTransferReqData`

NewN2InformationTransferReqDataWithDefaults instantiates a new N2InformationTransferReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaiList

`func (o *N2InformationTransferReqData) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *N2InformationTransferReqData) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *N2InformationTransferReqData) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *N2InformationTransferReqData) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetRatSelector

`func (o *N2InformationTransferReqData) GetRatSelector() RatSelector`

GetRatSelector returns the RatSelector field if non-nil, zero value otherwise.

### GetRatSelectorOk

`func (o *N2InformationTransferReqData) GetRatSelectorOk() (*RatSelector, bool)`

GetRatSelectorOk returns a tuple with the RatSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatSelector

`func (o *N2InformationTransferReqData) SetRatSelector(v RatSelector)`

SetRatSelector sets RatSelector field to given value.

### HasRatSelector

`func (o *N2InformationTransferReqData) HasRatSelector() bool`

HasRatSelector returns a boolean if a field has been set.

### GetGlobalRanNodeList

`func (o *N2InformationTransferReqData) GetGlobalRanNodeList() []GlobalRanNodeId`

GetGlobalRanNodeList returns the GlobalRanNodeList field if non-nil, zero value otherwise.

### GetGlobalRanNodeListOk

`func (o *N2InformationTransferReqData) GetGlobalRanNodeListOk() (*[]GlobalRanNodeId, bool)`

GetGlobalRanNodeListOk returns a tuple with the GlobalRanNodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRanNodeList

`func (o *N2InformationTransferReqData) SetGlobalRanNodeList(v []GlobalRanNodeId)`

SetGlobalRanNodeList sets GlobalRanNodeList field to given value.

### HasGlobalRanNodeList

`func (o *N2InformationTransferReqData) HasGlobalRanNodeList() bool`

HasGlobalRanNodeList returns a boolean if a field has been set.

### GetN2Information

`func (o *N2InformationTransferReqData) GetN2Information() N2InfoContainer`

GetN2Information returns the N2Information field if non-nil, zero value otherwise.

### GetN2InformationOk

`func (o *N2InformationTransferReqData) GetN2InformationOk() (*N2InfoContainer, bool)`

GetN2InformationOk returns a tuple with the N2Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2Information

`func (o *N2InformationTransferReqData) SetN2Information(v N2InfoContainer)`

SetN2Information sets N2Information field to given value.


### GetSupportedFeatures

`func (o *N2InformationTransferReqData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *N2InformationTransferReqData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *N2InformationTransferReqData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *N2InformationTransferReqData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


