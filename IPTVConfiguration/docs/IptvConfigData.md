# IptvConfigData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**ExterGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**AfAppId** | **string** |  | 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**MultiAccCtrls** | [**map[string]MulticastAccessControl**](MulticastAccessControl.md) | Identifies a list of multicast address access control information. Any string value can be used as a key of the map.  | 
**MtcProviderId** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 
**SuppFeat** | **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | 

## Methods

### NewIptvConfigData

`func NewIptvConfigData(afAppId string, multiAccCtrls map[string]MulticastAccessControl, suppFeat string, ) *IptvConfigData`

NewIptvConfigData instantiates a new IptvConfigData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIptvConfigDataWithDefaults

`func NewIptvConfigDataWithDefaults() *IptvConfigData`

NewIptvConfigDataWithDefaults instantiates a new IptvConfigData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *IptvConfigData) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *IptvConfigData) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *IptvConfigData) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *IptvConfigData) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetGpsi

`func (o *IptvConfigData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *IptvConfigData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *IptvConfigData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *IptvConfigData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetExterGroupId

`func (o *IptvConfigData) GetExterGroupId() string`

GetExterGroupId returns the ExterGroupId field if non-nil, zero value otherwise.

### GetExterGroupIdOk

`func (o *IptvConfigData) GetExterGroupIdOk() (*string, bool)`

GetExterGroupIdOk returns a tuple with the ExterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterGroupId

`func (o *IptvConfigData) SetExterGroupId(v string)`

SetExterGroupId sets ExterGroupId field to given value.

### HasExterGroupId

`func (o *IptvConfigData) HasExterGroupId() bool`

HasExterGroupId returns a boolean if a field has been set.

### GetAfAppId

`func (o *IptvConfigData) GetAfAppId() string`

GetAfAppId returns the AfAppId field if non-nil, zero value otherwise.

### GetAfAppIdOk

`func (o *IptvConfigData) GetAfAppIdOk() (*string, bool)`

GetAfAppIdOk returns a tuple with the AfAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAppId

`func (o *IptvConfigData) SetAfAppId(v string)`

SetAfAppId sets AfAppId field to given value.


### GetDnn

`func (o *IptvConfigData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *IptvConfigData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *IptvConfigData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *IptvConfigData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *IptvConfigData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *IptvConfigData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *IptvConfigData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *IptvConfigData) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetMultiAccCtrls

`func (o *IptvConfigData) GetMultiAccCtrls() map[string]MulticastAccessControl`

GetMultiAccCtrls returns the MultiAccCtrls field if non-nil, zero value otherwise.

### GetMultiAccCtrlsOk

`func (o *IptvConfigData) GetMultiAccCtrlsOk() (*map[string]MulticastAccessControl, bool)`

GetMultiAccCtrlsOk returns a tuple with the MultiAccCtrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAccCtrls

`func (o *IptvConfigData) SetMultiAccCtrls(v map[string]MulticastAccessControl)`

SetMultiAccCtrls sets MultiAccCtrls field to given value.


### GetMtcProviderId

`func (o *IptvConfigData) GetMtcProviderId() string`

GetMtcProviderId returns the MtcProviderId field if non-nil, zero value otherwise.

### GetMtcProviderIdOk

`func (o *IptvConfigData) GetMtcProviderIdOk() (*string, bool)`

GetMtcProviderIdOk returns a tuple with the MtcProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderId

`func (o *IptvConfigData) SetMtcProviderId(v string)`

SetMtcProviderId sets MtcProviderId field to given value.

### HasMtcProviderId

`func (o *IptvConfigData) HasMtcProviderId() bool`

HasMtcProviderId returns a boolean if a field has been set.

### GetSuppFeat

`func (o *IptvConfigData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *IptvConfigData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *IptvConfigData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


