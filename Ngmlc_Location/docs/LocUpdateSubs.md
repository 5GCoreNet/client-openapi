# LocUpdateSubs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**NotifURI** | **string** | String providing an URI formatted according to RFC 3986. | 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 

## Methods

### NewLocUpdateSubs

`func NewLocUpdateSubs(nfInstanceId string, notifURI string, ) *LocUpdateSubs`

NewLocUpdateSubs instantiates a new LocUpdateSubs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocUpdateSubsWithDefaults

`func NewLocUpdateSubsWithDefaults() *LocUpdateSubs`

NewLocUpdateSubsWithDefaults instantiates a new LocUpdateSubs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *LocUpdateSubs) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *LocUpdateSubs) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *LocUpdateSubs) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetNotifURI

`func (o *LocUpdateSubs) GetNotifURI() string`

GetNotifURI returns the NotifURI field if non-nil, zero value otherwise.

### GetNotifURIOk

`func (o *LocUpdateSubs) GetNotifURIOk() (*string, bool)`

GetNotifURIOk returns a tuple with the NotifURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifURI

`func (o *LocUpdateSubs) SetNotifURI(v string)`

SetNotifURI sets NotifURI field to given value.


### GetGpsi

`func (o *LocUpdateSubs) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *LocUpdateSubs) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *LocUpdateSubs) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *LocUpdateSubs) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSupi

`func (o *LocUpdateSubs) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *LocUpdateSubs) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *LocUpdateSubs) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *LocUpdateSubs) HasSupi() bool`

HasSupi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


