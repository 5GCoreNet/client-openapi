# ConfigForPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**N6Ind** | Pointer to **bool** |  | [optional] 
**PtpEnable** | Pointer to **bool** |  | [optional] 
**LogSyncInter** | Pointer to **int32** |  | [optional] 
**LogSyncInterInd** | Pointer to **bool** |  | [optional] 
**LogAnnouInter** | Pointer to **int32** |  | [optional] 
**LogAnnouInterInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewConfigForPort

`func NewConfigForPort() *ConfigForPort`

NewConfigForPort instantiates a new ConfigForPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigForPortWithDefaults

`func NewConfigForPortWithDefaults() *ConfigForPort`

NewConfigForPortWithDefaults instantiates a new ConfigForPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *ConfigForPort) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *ConfigForPort) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *ConfigForPort) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *ConfigForPort) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *ConfigForPort) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *ConfigForPort) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *ConfigForPort) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *ConfigForPort) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetN6Ind

`func (o *ConfigForPort) GetN6Ind() bool`

GetN6Ind returns the N6Ind field if non-nil, zero value otherwise.

### GetN6IndOk

`func (o *ConfigForPort) GetN6IndOk() (*bool, bool)`

GetN6IndOk returns a tuple with the N6Ind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN6Ind

`func (o *ConfigForPort) SetN6Ind(v bool)`

SetN6Ind sets N6Ind field to given value.

### HasN6Ind

`func (o *ConfigForPort) HasN6Ind() bool`

HasN6Ind returns a boolean if a field has been set.

### GetPtpEnable

`func (o *ConfigForPort) GetPtpEnable() bool`

GetPtpEnable returns the PtpEnable field if non-nil, zero value otherwise.

### GetPtpEnableOk

`func (o *ConfigForPort) GetPtpEnableOk() (*bool, bool)`

GetPtpEnableOk returns a tuple with the PtpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpEnable

`func (o *ConfigForPort) SetPtpEnable(v bool)`

SetPtpEnable sets PtpEnable field to given value.

### HasPtpEnable

`func (o *ConfigForPort) HasPtpEnable() bool`

HasPtpEnable returns a boolean if a field has been set.

### GetLogSyncInter

`func (o *ConfigForPort) GetLogSyncInter() int32`

GetLogSyncInter returns the LogSyncInter field if non-nil, zero value otherwise.

### GetLogSyncInterOk

`func (o *ConfigForPort) GetLogSyncInterOk() (*int32, bool)`

GetLogSyncInterOk returns a tuple with the LogSyncInter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncInter

`func (o *ConfigForPort) SetLogSyncInter(v int32)`

SetLogSyncInter sets LogSyncInter field to given value.

### HasLogSyncInter

`func (o *ConfigForPort) HasLogSyncInter() bool`

HasLogSyncInter returns a boolean if a field has been set.

### GetLogSyncInterInd

`func (o *ConfigForPort) GetLogSyncInterInd() bool`

GetLogSyncInterInd returns the LogSyncInterInd field if non-nil, zero value otherwise.

### GetLogSyncInterIndOk

`func (o *ConfigForPort) GetLogSyncInterIndOk() (*bool, bool)`

GetLogSyncInterIndOk returns a tuple with the LogSyncInterInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncInterInd

`func (o *ConfigForPort) SetLogSyncInterInd(v bool)`

SetLogSyncInterInd sets LogSyncInterInd field to given value.

### HasLogSyncInterInd

`func (o *ConfigForPort) HasLogSyncInterInd() bool`

HasLogSyncInterInd returns a boolean if a field has been set.

### GetLogAnnouInter

`func (o *ConfigForPort) GetLogAnnouInter() int32`

GetLogAnnouInter returns the LogAnnouInter field if non-nil, zero value otherwise.

### GetLogAnnouInterOk

`func (o *ConfigForPort) GetLogAnnouInterOk() (*int32, bool)`

GetLogAnnouInterOk returns a tuple with the LogAnnouInter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAnnouInter

`func (o *ConfigForPort) SetLogAnnouInter(v int32)`

SetLogAnnouInter sets LogAnnouInter field to given value.

### HasLogAnnouInter

`func (o *ConfigForPort) HasLogAnnouInter() bool`

HasLogAnnouInter returns a boolean if a field has been set.

### GetLogAnnouInterInd

`func (o *ConfigForPort) GetLogAnnouInterInd() bool`

GetLogAnnouInterInd returns the LogAnnouInterInd field if non-nil, zero value otherwise.

### GetLogAnnouInterIndOk

`func (o *ConfigForPort) GetLogAnnouInterIndOk() (*bool, bool)`

GetLogAnnouInterIndOk returns a tuple with the LogAnnouInterInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAnnouInterInd

`func (o *ConfigForPort) SetLogAnnouInterInd(v bool)`

SetLogAnnouInterInd sets LogAnnouInterInd field to given value.

### HasLogAnnouInterInd

`func (o *ConfigForPort) HasLogAnnouInterInd() bool`

HasLogAnnouInterInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


