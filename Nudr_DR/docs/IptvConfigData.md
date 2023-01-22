# IptvConfigData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**InterGroupId** | Pointer to **interface{}** | Identifies a group of users. | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**AfAppId** | **string** |  | 
**MultiAccCtrls** | [**map[string]MulticastAccessControl**](MulticastAccessControl.md) | Identifies a list of multicast address access control information. Any string value can be used as a key of the map.  | 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ResUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIptvConfigData

`func NewIptvConfigData(afAppId string, multiAccCtrls map[string]MulticastAccessControl, ) *IptvConfigData`

NewIptvConfigData instantiates a new IptvConfigData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIptvConfigDataWithDefaults

`func NewIptvConfigDataWithDefaults() *IptvConfigData`

NewIptvConfigDataWithDefaults instantiates a new IptvConfigData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *IptvConfigData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *IptvConfigData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *IptvConfigData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *IptvConfigData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetInterGroupId

`func (o *IptvConfigData) GetInterGroupId() interface{}`

GetInterGroupId returns the InterGroupId field if non-nil, zero value otherwise.

### GetInterGroupIdOk

`func (o *IptvConfigData) GetInterGroupIdOk() (*interface{}, bool)`

GetInterGroupIdOk returns a tuple with the InterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGroupId

`func (o *IptvConfigData) SetInterGroupId(v interface{})`

SetInterGroupId sets InterGroupId field to given value.

### HasInterGroupId

`func (o *IptvConfigData) HasInterGroupId() bool`

HasInterGroupId returns a boolean if a field has been set.

### SetInterGroupIdNil

`func (o *IptvConfigData) SetInterGroupIdNil(b bool)`

 SetInterGroupIdNil sets the value for InterGroupId to be an explicit nil

### UnsetInterGroupId
`func (o *IptvConfigData) UnsetInterGroupId()`

UnsetInterGroupId ensures that no value is present for InterGroupId, not even an explicit nil
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

### HasSuppFeat

`func (o *IptvConfigData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetResUri

`func (o *IptvConfigData) GetResUri() string`

GetResUri returns the ResUri field if non-nil, zero value otherwise.

### GetResUriOk

`func (o *IptvConfigData) GetResUriOk() (*string, bool)`

GetResUriOk returns a tuple with the ResUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResUri

`func (o *IptvConfigData) SetResUri(v string)`

SetResUri sets ResUri field to given value.

### HasResUri

`func (o *IptvConfigData) HasResUri() bool`

HasResUri returns a boolean if a field has been set.

### GetResetIds

`func (o *IptvConfigData) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *IptvConfigData) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *IptvConfigData) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *IptvConfigData) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


