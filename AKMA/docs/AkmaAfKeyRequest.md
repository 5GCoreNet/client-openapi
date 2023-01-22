# AkmaAfKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**AfId** | **string** | Represents an AF identifier. | 
**AKId** | **string** | Represents an AKMA Key Identifier. | 
**AnonInd** | Pointer to **bool** | Indicates whether an anonymous user access. Set to \&quot;true\&quot; if an anonymous user access is  requested; otherwise set to \&quot;false\&quot;. Default value is \&quot;false\&quot; if omitted.  | [optional] [default to false]

## Methods

### NewAkmaAfKeyRequest

`func NewAkmaAfKeyRequest(afId string, aKId string, ) *AkmaAfKeyRequest`

NewAkmaAfKeyRequest instantiates a new AkmaAfKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAkmaAfKeyRequestWithDefaults

`func NewAkmaAfKeyRequestWithDefaults() *AkmaAfKeyRequest`

NewAkmaAfKeyRequestWithDefaults instantiates a new AkmaAfKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuppFeat

`func (o *AkmaAfKeyRequest) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AkmaAfKeyRequest) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AkmaAfKeyRequest) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AkmaAfKeyRequest) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetAfId

`func (o *AkmaAfKeyRequest) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *AkmaAfKeyRequest) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *AkmaAfKeyRequest) SetAfId(v string)`

SetAfId sets AfId field to given value.


### GetAKId

`func (o *AkmaAfKeyRequest) GetAKId() string`

GetAKId returns the AKId field if non-nil, zero value otherwise.

### GetAKIdOk

`func (o *AkmaAfKeyRequest) GetAKIdOk() (*string, bool)`

GetAKIdOk returns a tuple with the AKId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAKId

`func (o *AkmaAfKeyRequest) SetAKId(v string)`

SetAKId sets AKId field to given value.


### GetAnonInd

`func (o *AkmaAfKeyRequest) GetAnonInd() bool`

GetAnonInd returns the AnonInd field if non-nil, zero value otherwise.

### GetAnonIndOk

`func (o *AkmaAfKeyRequest) GetAnonIndOk() (*bool, bool)`

GetAnonIndOk returns a tuple with the AnonInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonInd

`func (o *AkmaAfKeyRequest) SetAnonInd(v bool)`

SetAnonInd sets AnonInd field to given value.

### HasAnonInd

`func (o *AkmaAfKeyRequest) HasAnonInd() bool`

HasAnonInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


