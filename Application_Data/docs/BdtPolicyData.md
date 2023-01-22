# BdtPolicyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterGroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**BdtRefId** | **string** | string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154. | 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**ResUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBdtPolicyData

`func NewBdtPolicyData(bdtRefId string, ) *BdtPolicyData`

NewBdtPolicyData instantiates a new BdtPolicyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBdtPolicyDataWithDefaults

`func NewBdtPolicyDataWithDefaults() *BdtPolicyData`

NewBdtPolicyDataWithDefaults instantiates a new BdtPolicyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterGroupId

`func (o *BdtPolicyData) GetInterGroupId() string`

GetInterGroupId returns the InterGroupId field if non-nil, zero value otherwise.

### GetInterGroupIdOk

`func (o *BdtPolicyData) GetInterGroupIdOk() (*string, bool)`

GetInterGroupIdOk returns a tuple with the InterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGroupId

`func (o *BdtPolicyData) SetInterGroupId(v string)`

SetInterGroupId sets InterGroupId field to given value.

### HasInterGroupId

`func (o *BdtPolicyData) HasInterGroupId() bool`

HasInterGroupId returns a boolean if a field has been set.

### GetSupi

`func (o *BdtPolicyData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *BdtPolicyData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *BdtPolicyData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *BdtPolicyData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetBdtRefId

`func (o *BdtPolicyData) GetBdtRefId() string`

GetBdtRefId returns the BdtRefId field if non-nil, zero value otherwise.

### GetBdtRefIdOk

`func (o *BdtPolicyData) GetBdtRefIdOk() (*string, bool)`

GetBdtRefIdOk returns a tuple with the BdtRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtRefId

`func (o *BdtPolicyData) SetBdtRefId(v string)`

SetBdtRefId sets BdtRefId field to given value.


### GetDnn

`func (o *BdtPolicyData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *BdtPolicyData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *BdtPolicyData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *BdtPolicyData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *BdtPolicyData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *BdtPolicyData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *BdtPolicyData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *BdtPolicyData) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetResUri

`func (o *BdtPolicyData) GetResUri() string`

GetResUri returns the ResUri field if non-nil, zero value otherwise.

### GetResUriOk

`func (o *BdtPolicyData) GetResUriOk() (*string, bool)`

GetResUriOk returns a tuple with the ResUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResUri

`func (o *BdtPolicyData) SetResUri(v string)`

SetResUri sets ResUri field to given value.

### HasResUri

`func (o *BdtPolicyData) HasResUri() bool`

HasResUri returns a boolean if a field has been set.

### GetResetIds

`func (o *BdtPolicyData) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *BdtPolicyData) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *BdtPolicyData) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *BdtPolicyData) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


