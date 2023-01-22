# UEContextRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**UnauthenticatedSupi** | Pointer to **bool** |  | [optional] [default to false]
**NgapCause** | [**NgApCause**](NgApCause.md) |  | 

## Methods

### NewUEContextRelease

`func NewUEContextRelease(ngapCause NgApCause, ) *UEContextRelease`

NewUEContextRelease instantiates a new UEContextRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUEContextReleaseWithDefaults

`func NewUEContextReleaseWithDefaults() *UEContextRelease`

NewUEContextReleaseWithDefaults instantiates a new UEContextRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *UEContextRelease) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UEContextRelease) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UEContextRelease) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *UEContextRelease) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetUnauthenticatedSupi

`func (o *UEContextRelease) GetUnauthenticatedSupi() bool`

GetUnauthenticatedSupi returns the UnauthenticatedSupi field if non-nil, zero value otherwise.

### GetUnauthenticatedSupiOk

`func (o *UEContextRelease) GetUnauthenticatedSupiOk() (*bool, bool)`

GetUnauthenticatedSupiOk returns a tuple with the UnauthenticatedSupi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedSupi

`func (o *UEContextRelease) SetUnauthenticatedSupi(v bool)`

SetUnauthenticatedSupi sets UnauthenticatedSupi field to given value.

### HasUnauthenticatedSupi

`func (o *UEContextRelease) HasUnauthenticatedSupi() bool`

HasUnauthenticatedSupi returns a boolean if a field has been set.

### GetNgapCause

`func (o *UEContextRelease) GetNgapCause() NgApCause`

GetNgapCause returns the NgapCause field if non-nil, zero value otherwise.

### GetNgapCauseOk

`func (o *UEContextRelease) GetNgapCauseOk() (*NgApCause, bool)`

GetNgapCauseOk returns a tuple with the NgapCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgapCause

`func (o *UEContextRelease) SetNgapCause(v NgApCause)`

SetNgapCause sets NgapCause field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


