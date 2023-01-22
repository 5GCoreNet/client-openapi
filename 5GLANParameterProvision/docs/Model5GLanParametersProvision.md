# Model5GLanParametersProvision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**Var5gLanParams** | [**Model5GLanParameters**](Model5GLanParameters.md) |  | 
**SuppFeat** | **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | 

## Methods

### NewModel5GLanParametersProvision

`func NewModel5GLanParametersProvision(var5gLanParams Model5GLanParameters, suppFeat string, ) *Model5GLanParametersProvision`

NewModel5GLanParametersProvision instantiates a new Model5GLanParametersProvision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel5GLanParametersProvisionWithDefaults

`func NewModel5GLanParametersProvisionWithDefaults() *Model5GLanParametersProvision`

NewModel5GLanParametersProvisionWithDefaults instantiates a new Model5GLanParametersProvision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *Model5GLanParametersProvision) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Model5GLanParametersProvision) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Model5GLanParametersProvision) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Model5GLanParametersProvision) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetVar5gLanParams

`func (o *Model5GLanParametersProvision) GetVar5gLanParams() Model5GLanParameters`

GetVar5gLanParams returns the Var5gLanParams field if non-nil, zero value otherwise.

### GetVar5gLanParamsOk

`func (o *Model5GLanParametersProvision) GetVar5gLanParamsOk() (*Model5GLanParameters, bool)`

GetVar5gLanParamsOk returns a tuple with the Var5gLanParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gLanParams

`func (o *Model5GLanParametersProvision) SetVar5gLanParams(v Model5GLanParameters)`

SetVar5gLanParams sets Var5gLanParams field to given value.


### GetSuppFeat

`func (o *Model5GLanParametersProvision) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *Model5GLanParametersProvision) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *Model5GLanParametersProvision) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


