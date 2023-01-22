# DispersionCollection1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**UeAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 
**DataUsage** | [**UsageThreshold**](UsageThreshold.md) |  | 
**FlowDesp** | Pointer to **string** | Defines a packet filter of an IP flow. | [optional] 
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**Dnais** | Pointer to **[]string** |  | [optional] 
**AppDur** | Pointer to **int32** | indicating a time in seconds. | [optional] 

## Methods

### NewDispersionCollection1

`func NewDispersionCollection1(dataUsage UsageThreshold, ) *DispersionCollection1`

NewDispersionCollection1 instantiates a new DispersionCollection1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDispersionCollection1WithDefaults

`func NewDispersionCollection1WithDefaults() *DispersionCollection1`

NewDispersionCollection1WithDefaults instantiates a new DispersionCollection1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *DispersionCollection1) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *DispersionCollection1) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *DispersionCollection1) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *DispersionCollection1) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSupi

`func (o *DispersionCollection1) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *DispersionCollection1) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *DispersionCollection1) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *DispersionCollection1) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetUeAddr

`func (o *DispersionCollection1) GetUeAddr() IpAddr`

GetUeAddr returns the UeAddr field if non-nil, zero value otherwise.

### GetUeAddrOk

`func (o *DispersionCollection1) GetUeAddrOk() (*IpAddr, bool)`

GetUeAddrOk returns a tuple with the UeAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeAddr

`func (o *DispersionCollection1) SetUeAddr(v IpAddr)`

SetUeAddr sets UeAddr field to given value.

### HasUeAddr

`func (o *DispersionCollection1) HasUeAddr() bool`

HasUeAddr returns a boolean if a field has been set.

### GetDataUsage

`func (o *DispersionCollection1) GetDataUsage() UsageThreshold`

GetDataUsage returns the DataUsage field if non-nil, zero value otherwise.

### GetDataUsageOk

`func (o *DispersionCollection1) GetDataUsageOk() (*UsageThreshold, bool)`

GetDataUsageOk returns a tuple with the DataUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUsage

`func (o *DispersionCollection1) SetDataUsage(v UsageThreshold)`

SetDataUsage sets DataUsage field to given value.


### GetFlowDesp

`func (o *DispersionCollection1) GetFlowDesp() string`

GetFlowDesp returns the FlowDesp field if non-nil, zero value otherwise.

### GetFlowDespOk

`func (o *DispersionCollection1) GetFlowDespOk() (*string, bool)`

GetFlowDespOk returns a tuple with the FlowDesp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDesp

`func (o *DispersionCollection1) SetFlowDesp(v string)`

SetFlowDesp sets FlowDesp field to given value.

### HasFlowDesp

`func (o *DispersionCollection1) HasFlowDesp() bool`

HasFlowDesp returns a boolean if a field has been set.

### GetAppId

`func (o *DispersionCollection1) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *DispersionCollection1) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *DispersionCollection1) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *DispersionCollection1) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDnais

`func (o *DispersionCollection1) GetDnais() []string`

GetDnais returns the Dnais field if non-nil, zero value otherwise.

### GetDnaisOk

`func (o *DispersionCollection1) GetDnaisOk() (*[]string, bool)`

GetDnaisOk returns a tuple with the Dnais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnais

`func (o *DispersionCollection1) SetDnais(v []string)`

SetDnais sets Dnais field to given value.

### HasDnais

`func (o *DispersionCollection1) HasDnais() bool`

HasDnais returns a boolean if a field has been set.

### GetAppDur

`func (o *DispersionCollection1) GetAppDur() int32`

GetAppDur returns the AppDur field if non-nil, zero value otherwise.

### GetAppDurOk

`func (o *DispersionCollection1) GetAppDurOk() (*int32, bool)`

GetAppDurOk returns a tuple with the AppDur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDur

`func (o *DispersionCollection1) SetAppDur(v int32)`

SetAppDur sets AppDur field to given value.

### HasAppDur

`func (o *DispersionCollection1) HasAppDur() bool`

HasAppDur returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


