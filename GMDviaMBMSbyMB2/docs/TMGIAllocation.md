# TMGIAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ExternalGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**MbmsLocArea** | Pointer to [**MbmsLocArea**](MbmsLocArea.md) |  | [optional] 
**TmgiExpiration** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI with \&quot;readOnly&#x3D;true\&quot; property. | [optional] [readonly] 

## Methods

### NewTMGIAllocation

`func NewTMGIAllocation() *TMGIAllocation`

NewTMGIAllocation instantiates a new TMGIAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTMGIAllocationWithDefaults

`func NewTMGIAllocationWithDefaults() *TMGIAllocation`

NewTMGIAllocationWithDefaults instantiates a new TMGIAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *TMGIAllocation) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *TMGIAllocation) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *TMGIAllocation) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *TMGIAllocation) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *TMGIAllocation) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *TMGIAllocation) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *TMGIAllocation) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *TMGIAllocation) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetExternalGroupId

`func (o *TMGIAllocation) GetExternalGroupId() string`

GetExternalGroupId returns the ExternalGroupId field if non-nil, zero value otherwise.

### GetExternalGroupIdOk

`func (o *TMGIAllocation) GetExternalGroupIdOk() (*string, bool)`

GetExternalGroupIdOk returns a tuple with the ExternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupId

`func (o *TMGIAllocation) SetExternalGroupId(v string)`

SetExternalGroupId sets ExternalGroupId field to given value.

### HasExternalGroupId

`func (o *TMGIAllocation) HasExternalGroupId() bool`

HasExternalGroupId returns a boolean if a field has been set.

### GetMbmsLocArea

`func (o *TMGIAllocation) GetMbmsLocArea() MbmsLocArea`

GetMbmsLocArea returns the MbmsLocArea field if non-nil, zero value otherwise.

### GetMbmsLocAreaOk

`func (o *TMGIAllocation) GetMbmsLocAreaOk() (*MbmsLocArea, bool)`

GetMbmsLocAreaOk returns a tuple with the MbmsLocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbmsLocArea

`func (o *TMGIAllocation) SetMbmsLocArea(v MbmsLocArea)`

SetMbmsLocArea sets MbmsLocArea field to given value.

### HasMbmsLocArea

`func (o *TMGIAllocation) HasMbmsLocArea() bool`

HasMbmsLocArea returns a boolean if a field has been set.

### GetTmgiExpiration

`func (o *TMGIAllocation) GetTmgiExpiration() time.Time`

GetTmgiExpiration returns the TmgiExpiration field if non-nil, zero value otherwise.

### GetTmgiExpirationOk

`func (o *TMGIAllocation) GetTmgiExpirationOk() (*time.Time, bool)`

GetTmgiExpirationOk returns a tuple with the TmgiExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmgiExpiration

`func (o *TMGIAllocation) SetTmgiExpiration(v time.Time)`

SetTmgiExpiration sets TmgiExpiration field to given value.

### HasTmgiExpiration

`func (o *TMGIAllocation) HasTmgiExpiration() bool`

HasTmgiExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


