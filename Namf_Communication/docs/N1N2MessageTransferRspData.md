# N1N2MessageTransferRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cause** | [**N1N2MessageTransferCause**](N1N2MessageTransferCause.md) |  | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewN1N2MessageTransferRspData

`func NewN1N2MessageTransferRspData(cause N1N2MessageTransferCause, ) *N1N2MessageTransferRspData`

NewN1N2MessageTransferRspData instantiates a new N1N2MessageTransferRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN1N2MessageTransferRspDataWithDefaults

`func NewN1N2MessageTransferRspDataWithDefaults() *N1N2MessageTransferRspData`

NewN1N2MessageTransferRspDataWithDefaults instantiates a new N1N2MessageTransferRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCause

`func (o *N1N2MessageTransferRspData) GetCause() N1N2MessageTransferCause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *N1N2MessageTransferRspData) GetCauseOk() (*N1N2MessageTransferCause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *N1N2MessageTransferRspData) SetCause(v N1N2MessageTransferCause)`

SetCause sets Cause field to given value.


### GetSupportedFeatures

`func (o *N1N2MessageTransferRspData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *N1N2MessageTransferRspData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *N1N2MessageTransferRspData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *N1N2MessageTransferRspData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


