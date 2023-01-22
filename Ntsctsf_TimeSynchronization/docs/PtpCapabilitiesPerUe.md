# PtpCapabilitiesPerUe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**PtpCaps** | [**[]EventFilter**](EventFilter.md) |  | 

## Methods

### NewPtpCapabilitiesPerUe

`func NewPtpCapabilitiesPerUe(ptpCaps []EventFilter, ) *PtpCapabilitiesPerUe`

NewPtpCapabilitiesPerUe instantiates a new PtpCapabilitiesPerUe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPtpCapabilitiesPerUeWithDefaults

`func NewPtpCapabilitiesPerUeWithDefaults() *PtpCapabilitiesPerUe`

NewPtpCapabilitiesPerUeWithDefaults instantiates a new PtpCapabilitiesPerUe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *PtpCapabilitiesPerUe) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *PtpCapabilitiesPerUe) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *PtpCapabilitiesPerUe) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *PtpCapabilitiesPerUe) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *PtpCapabilitiesPerUe) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *PtpCapabilitiesPerUe) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *PtpCapabilitiesPerUe) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *PtpCapabilitiesPerUe) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetPtpCaps

`func (o *PtpCapabilitiesPerUe) GetPtpCaps() []EventFilter`

GetPtpCaps returns the PtpCaps field if non-nil, zero value otherwise.

### GetPtpCapsOk

`func (o *PtpCapabilitiesPerUe) GetPtpCapsOk() (*[]EventFilter, bool)`

GetPtpCapsOk returns a tuple with the PtpCaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpCaps

`func (o *PtpCapabilitiesPerUe) SetPtpCaps(v []EventFilter)`

SetPtpCaps sets PtpCaps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


