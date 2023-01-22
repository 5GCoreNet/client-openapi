# AppliedBdtPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**BdtRefId** | **string** | string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154. | 
**SuppFeat** | **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | 
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 

## Methods

### NewAppliedBdtPolicy

`func NewAppliedBdtPolicy(bdtRefId string, suppFeat string, ) *AppliedBdtPolicy`

NewAppliedBdtPolicy instantiates a new AppliedBdtPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedBdtPolicyWithDefaults

`func NewAppliedBdtPolicyWithDefaults() *AppliedBdtPolicy`

NewAppliedBdtPolicyWithDefaults instantiates a new AppliedBdtPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalGroupId

`func (o *AppliedBdtPolicy) GetExternalGroupId() string`

GetExternalGroupId returns the ExternalGroupId field if non-nil, zero value otherwise.

### GetExternalGroupIdOk

`func (o *AppliedBdtPolicy) GetExternalGroupIdOk() (*string, bool)`

GetExternalGroupIdOk returns a tuple with the ExternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupId

`func (o *AppliedBdtPolicy) SetExternalGroupId(v string)`

SetExternalGroupId sets ExternalGroupId field to given value.

### HasExternalGroupId

`func (o *AppliedBdtPolicy) HasExternalGroupId() bool`

HasExternalGroupId returns a boolean if a field has been set.

### GetGpsi

`func (o *AppliedBdtPolicy) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *AppliedBdtPolicy) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *AppliedBdtPolicy) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *AppliedBdtPolicy) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetBdtRefId

`func (o *AppliedBdtPolicy) GetBdtRefId() string`

GetBdtRefId returns the BdtRefId field if non-nil, zero value otherwise.

### GetBdtRefIdOk

`func (o *AppliedBdtPolicy) GetBdtRefIdOk() (*string, bool)`

GetBdtRefIdOk returns a tuple with the BdtRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtRefId

`func (o *AppliedBdtPolicy) SetBdtRefId(v string)`

SetBdtRefId sets BdtRefId field to given value.


### GetSuppFeat

`func (o *AppliedBdtPolicy) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AppliedBdtPolicy) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AppliedBdtPolicy) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.


### GetSelf

`func (o *AppliedBdtPolicy) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AppliedBdtPolicy) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AppliedBdtPolicy) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AppliedBdtPolicy) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


