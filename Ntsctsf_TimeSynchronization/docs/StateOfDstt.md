# StateOfDstt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**State** | **bool** | When the PTP port state is Leader, Follower or Passive, it is included and set to true to indicate the state of configuration for DS-TT port is active; when PTP port state is in any other case, it is included and set to false to indicate the state of configuration for DS-TT port is inactive. Default value is false.  | 

## Methods

### NewStateOfDstt

`func NewStateOfDstt(state bool, ) *StateOfDstt`

NewStateOfDstt instantiates a new StateOfDstt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateOfDsttWithDefaults

`func NewStateOfDsttWithDefaults() *StateOfDstt`

NewStateOfDsttWithDefaults instantiates a new StateOfDstt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *StateOfDstt) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *StateOfDstt) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *StateOfDstt) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *StateOfDstt) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *StateOfDstt) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *StateOfDstt) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *StateOfDstt) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *StateOfDstt) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetState

`func (o *StateOfDstt) GetState() bool`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StateOfDstt) GetStateOk() (*bool, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StateOfDstt) SetState(v bool)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


