# UpuInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpuDataList** | [**[]UpuData**](UpuData.md) |  | 
**UpuHeader** | Pointer to **string** | Contains the \&quot;UPU Header\&quot; IE as specified in clause 9.11.3.53A of 3GPP TS 24.501  (octet 4), encoded as 2 hexadecimal characters. | [optional] 
**UpuAckInd** | **bool** | Contains the indication of whether the acknowledgement from UE is needed. | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**UpuTransparentInfo** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 

## Methods

### NewUpuInfo

`func NewUpuInfo(upuDataList []UpuData, upuAckInd bool, ) *UpuInfo`

NewUpuInfo instantiates a new UpuInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpuInfoWithDefaults

`func NewUpuInfoWithDefaults() *UpuInfo`

NewUpuInfoWithDefaults instantiates a new UpuInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpuDataList

`func (o *UpuInfo) GetUpuDataList() []UpuData`

GetUpuDataList returns the UpuDataList field if non-nil, zero value otherwise.

### GetUpuDataListOk

`func (o *UpuInfo) GetUpuDataListOk() (*[]UpuData, bool)`

GetUpuDataListOk returns a tuple with the UpuDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuDataList

`func (o *UpuInfo) SetUpuDataList(v []UpuData)`

SetUpuDataList sets UpuDataList field to given value.


### GetUpuHeader

`func (o *UpuInfo) GetUpuHeader() string`

GetUpuHeader returns the UpuHeader field if non-nil, zero value otherwise.

### GetUpuHeaderOk

`func (o *UpuInfo) GetUpuHeaderOk() (*string, bool)`

GetUpuHeaderOk returns a tuple with the UpuHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuHeader

`func (o *UpuInfo) SetUpuHeader(v string)`

SetUpuHeader sets UpuHeader field to given value.

### HasUpuHeader

`func (o *UpuInfo) HasUpuHeader() bool`

HasUpuHeader returns a boolean if a field has been set.

### GetUpuAckInd

`func (o *UpuInfo) GetUpuAckInd() bool`

GetUpuAckInd returns the UpuAckInd field if non-nil, zero value otherwise.

### GetUpuAckIndOk

`func (o *UpuInfo) GetUpuAckIndOk() (*bool, bool)`

GetUpuAckIndOk returns a tuple with the UpuAckInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuAckInd

`func (o *UpuInfo) SetUpuAckInd(v bool)`

SetUpuAckInd sets UpuAckInd field to given value.


### GetSupportedFeatures

`func (o *UpuInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *UpuInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *UpuInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *UpuInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetUpuTransparentInfo

`func (o *UpuInfo) GetUpuTransparentInfo() string`

GetUpuTransparentInfo returns the UpuTransparentInfo field if non-nil, zero value otherwise.

### GetUpuTransparentInfoOk

`func (o *UpuInfo) GetUpuTransparentInfoOk() (*string, bool)`

GetUpuTransparentInfoOk returns a tuple with the UpuTransparentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuTransparentInfo

`func (o *UpuInfo) SetUpuTransparentInfo(v string)`

SetUpuTransparentInfo sets UpuTransparentInfo field to given value.

### HasUpuTransparentInfo

`func (o *UpuInfo) HasUpuTransparentInfo() bool`

HasUpuTransparentInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


