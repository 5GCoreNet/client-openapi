# SorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SteeringContainer** | Pointer to [**SteeringContainer**](SteeringContainer.md) |  | [optional] 
**AckInd** | **bool** | Contains indication whether the acknowledgement from UE is needed. | 
**SorHeader** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**SorTransparentInfo** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewSorInfo

`func NewSorInfo(ackInd bool, ) *SorInfo`

NewSorInfo instantiates a new SorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSorInfoWithDefaults

`func NewSorInfoWithDefaults() *SorInfo`

NewSorInfoWithDefaults instantiates a new SorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSteeringContainer

`func (o *SorInfo) GetSteeringContainer() SteeringContainer`

GetSteeringContainer returns the SteeringContainer field if non-nil, zero value otherwise.

### GetSteeringContainerOk

`func (o *SorInfo) GetSteeringContainerOk() (*SteeringContainer, bool)`

GetSteeringContainerOk returns a tuple with the SteeringContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteeringContainer

`func (o *SorInfo) SetSteeringContainer(v SteeringContainer)`

SetSteeringContainer sets SteeringContainer field to given value.

### HasSteeringContainer

`func (o *SorInfo) HasSteeringContainer() bool`

HasSteeringContainer returns a boolean if a field has been set.

### GetAckInd

`func (o *SorInfo) GetAckInd() bool`

GetAckInd returns the AckInd field if non-nil, zero value otherwise.

### GetAckIndOk

`func (o *SorInfo) GetAckIndOk() (*bool, bool)`

GetAckIndOk returns a tuple with the AckInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckInd

`func (o *SorInfo) SetAckInd(v bool)`

SetAckInd sets AckInd field to given value.


### GetSorHeader

`func (o *SorInfo) GetSorHeader() string`

GetSorHeader returns the SorHeader field if non-nil, zero value otherwise.

### GetSorHeaderOk

`func (o *SorInfo) GetSorHeaderOk() (*string, bool)`

GetSorHeaderOk returns a tuple with the SorHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorHeader

`func (o *SorInfo) SetSorHeader(v string)`

SetSorHeader sets SorHeader field to given value.

### HasSorHeader

`func (o *SorInfo) HasSorHeader() bool`

HasSorHeader returns a boolean if a field has been set.

### GetSorTransparentInfo

`func (o *SorInfo) GetSorTransparentInfo() string`

GetSorTransparentInfo returns the SorTransparentInfo field if non-nil, zero value otherwise.

### GetSorTransparentInfoOk

`func (o *SorInfo) GetSorTransparentInfoOk() (*string, bool)`

GetSorTransparentInfoOk returns a tuple with the SorTransparentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorTransparentInfo

`func (o *SorInfo) SetSorTransparentInfo(v string)`

SetSorTransparentInfo sets SorTransparentInfo field to given value.

### HasSorTransparentInfo

`func (o *SorInfo) HasSorTransparentInfo() bool`

HasSorTransparentInfo returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SorInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SorInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SorInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SorInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


