# SorInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**SteeringContainer** | Pointer to [**SteeringContainer**](SteeringContainer.md) |  | [optional] 
**SorAckIndication** | **bool** |  | 
**SorCmci** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**StoreSorCmciInMe** | Pointer to **bool** |  | [optional] 
**SorSendingTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 

## Methods

### NewSorInformation

`func NewSorInformation(sorAckIndication bool, sorSendingTime time.Time, ) *SorInformation`

NewSorInformation instantiates a new SorInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSorInformationWithDefaults

`func NewSorInformationWithDefaults() *SorInformation`

NewSorInformationWithDefaults instantiates a new SorInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *SorInformation) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SorInformation) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SorInformation) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SorInformation) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetSteeringContainer

`func (o *SorInformation) GetSteeringContainer() SteeringContainer`

GetSteeringContainer returns the SteeringContainer field if non-nil, zero value otherwise.

### GetSteeringContainerOk

`func (o *SorInformation) GetSteeringContainerOk() (*SteeringContainer, bool)`

GetSteeringContainerOk returns a tuple with the SteeringContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteeringContainer

`func (o *SorInformation) SetSteeringContainer(v SteeringContainer)`

SetSteeringContainer sets SteeringContainer field to given value.

### HasSteeringContainer

`func (o *SorInformation) HasSteeringContainer() bool`

HasSteeringContainer returns a boolean if a field has been set.

### GetSorAckIndication

`func (o *SorInformation) GetSorAckIndication() bool`

GetSorAckIndication returns the SorAckIndication field if non-nil, zero value otherwise.

### GetSorAckIndicationOk

`func (o *SorInformation) GetSorAckIndicationOk() (*bool, bool)`

GetSorAckIndicationOk returns a tuple with the SorAckIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorAckIndication

`func (o *SorInformation) SetSorAckIndication(v bool)`

SetSorAckIndication sets SorAckIndication field to given value.


### GetSorCmci

`func (o *SorInformation) GetSorCmci() string`

GetSorCmci returns the SorCmci field if non-nil, zero value otherwise.

### GetSorCmciOk

`func (o *SorInformation) GetSorCmciOk() (*string, bool)`

GetSorCmciOk returns a tuple with the SorCmci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorCmci

`func (o *SorInformation) SetSorCmci(v string)`

SetSorCmci sets SorCmci field to given value.

### HasSorCmci

`func (o *SorInformation) HasSorCmci() bool`

HasSorCmci returns a boolean if a field has been set.

### GetStoreSorCmciInMe

`func (o *SorInformation) GetStoreSorCmciInMe() bool`

GetStoreSorCmciInMe returns the StoreSorCmciInMe field if non-nil, zero value otherwise.

### GetStoreSorCmciInMeOk

`func (o *SorInformation) GetStoreSorCmciInMeOk() (*bool, bool)`

GetStoreSorCmciInMeOk returns a tuple with the StoreSorCmciInMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSorCmciInMe

`func (o *SorInformation) SetStoreSorCmciInMe(v bool)`

SetStoreSorCmciInMe sets StoreSorCmciInMe field to given value.

### HasStoreSorCmciInMe

`func (o *SorInformation) HasStoreSorCmciInMe() bool`

HasStoreSorCmciInMe returns a boolean if a field has been set.

### GetSorSendingTime

`func (o *SorInformation) GetSorSendingTime() time.Time`

GetSorSendingTime returns the SorSendingTime field if non-nil, zero value otherwise.

### GetSorSendingTimeOk

`func (o *SorInformation) GetSorSendingTimeOk() (*time.Time, bool)`

GetSorSendingTimeOk returns a tuple with the SorSendingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorSendingTime

`func (o *SorInformation) SetSorSendingTime(v time.Time)`

SetSorSendingTime sets SorSendingTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


