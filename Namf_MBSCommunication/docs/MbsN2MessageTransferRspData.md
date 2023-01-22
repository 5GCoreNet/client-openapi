# MbsN2MessageTransferRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**N2InformationTransferResult**](N2InformationTransferResult.md) |  | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewMbsN2MessageTransferRspData

`func NewMbsN2MessageTransferRspData(result N2InformationTransferResult, ) *MbsN2MessageTransferRspData`

NewMbsN2MessageTransferRspData instantiates a new MbsN2MessageTransferRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsN2MessageTransferRspDataWithDefaults

`func NewMbsN2MessageTransferRspDataWithDefaults() *MbsN2MessageTransferRspData`

NewMbsN2MessageTransferRspDataWithDefaults instantiates a new MbsN2MessageTransferRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *MbsN2MessageTransferRspData) GetResult() N2InformationTransferResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MbsN2MessageTransferRspData) GetResultOk() (*N2InformationTransferResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MbsN2MessageTransferRspData) SetResult(v N2InformationTransferResult)`

SetResult sets Result field to given value.


### GetSupportedFeatures

`func (o *MbsN2MessageTransferRspData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *MbsN2MessageTransferRspData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *MbsN2MessageTransferRspData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *MbsN2MessageTransferRspData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


