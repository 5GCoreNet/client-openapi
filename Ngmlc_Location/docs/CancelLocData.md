# CancelLocData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**ExtGroupId** | Pointer to **string** | String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.   | [optional] 
**IntGroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**HgmlcCallBackUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**LdrReference** | **string** | LDR Reference. | 
**LmfIdentification** | Pointer to **string** | LMF identification. | [optional] 
**AmfId** | Pointer to **string** | String identifying the AMF ID composed of AMF Region ID (8 bits), AMF Set ID (10 bits) and AMF  Pointer (6 bits) as specified in clause 2.10.1 of 3GPP TS 23.003. It is encoded as a string of  6 hexadecimal characters (i.e., 24 bits).   | [optional] 

## Methods

### NewCancelLocData

`func NewCancelLocData(hgmlcCallBackUri string, ldrReference string, ) *CancelLocData`

NewCancelLocData instantiates a new CancelLocData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelLocDataWithDefaults

`func NewCancelLocDataWithDefaults() *CancelLocData`

NewCancelLocDataWithDefaults instantiates a new CancelLocData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *CancelLocData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *CancelLocData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *CancelLocData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *CancelLocData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSupi

`func (o *CancelLocData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *CancelLocData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *CancelLocData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *CancelLocData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetExtGroupId

`func (o *CancelLocData) GetExtGroupId() string`

GetExtGroupId returns the ExtGroupId field if non-nil, zero value otherwise.

### GetExtGroupIdOk

`func (o *CancelLocData) GetExtGroupIdOk() (*string, bool)`

GetExtGroupIdOk returns a tuple with the ExtGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtGroupId

`func (o *CancelLocData) SetExtGroupId(v string)`

SetExtGroupId sets ExtGroupId field to given value.

### HasExtGroupId

`func (o *CancelLocData) HasExtGroupId() bool`

HasExtGroupId returns a boolean if a field has been set.

### GetIntGroupId

`func (o *CancelLocData) GetIntGroupId() string`

GetIntGroupId returns the IntGroupId field if non-nil, zero value otherwise.

### GetIntGroupIdOk

`func (o *CancelLocData) GetIntGroupIdOk() (*string, bool)`

GetIntGroupIdOk returns a tuple with the IntGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntGroupId

`func (o *CancelLocData) SetIntGroupId(v string)`

SetIntGroupId sets IntGroupId field to given value.

### HasIntGroupId

`func (o *CancelLocData) HasIntGroupId() bool`

HasIntGroupId returns a boolean if a field has been set.

### GetHgmlcCallBackUri

`func (o *CancelLocData) GetHgmlcCallBackUri() string`

GetHgmlcCallBackUri returns the HgmlcCallBackUri field if non-nil, zero value otherwise.

### GetHgmlcCallBackUriOk

`func (o *CancelLocData) GetHgmlcCallBackUriOk() (*string, bool)`

GetHgmlcCallBackUriOk returns a tuple with the HgmlcCallBackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgmlcCallBackUri

`func (o *CancelLocData) SetHgmlcCallBackUri(v string)`

SetHgmlcCallBackUri sets HgmlcCallBackUri field to given value.


### GetLdrReference

`func (o *CancelLocData) GetLdrReference() string`

GetLdrReference returns the LdrReference field if non-nil, zero value otherwise.

### GetLdrReferenceOk

`func (o *CancelLocData) GetLdrReferenceOk() (*string, bool)`

GetLdrReferenceOk returns a tuple with the LdrReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdrReference

`func (o *CancelLocData) SetLdrReference(v string)`

SetLdrReference sets LdrReference field to given value.


### GetLmfIdentification

`func (o *CancelLocData) GetLmfIdentification() string`

GetLmfIdentification returns the LmfIdentification field if non-nil, zero value otherwise.

### GetLmfIdentificationOk

`func (o *CancelLocData) GetLmfIdentificationOk() (*string, bool)`

GetLmfIdentificationOk returns a tuple with the LmfIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmfIdentification

`func (o *CancelLocData) SetLmfIdentification(v string)`

SetLmfIdentification sets LmfIdentification field to given value.

### HasLmfIdentification

`func (o *CancelLocData) HasLmfIdentification() bool`

HasLmfIdentification returns a boolean if a field has been set.

### GetAmfId

`func (o *CancelLocData) GetAmfId() string`

GetAmfId returns the AmfId field if non-nil, zero value otherwise.

### GetAmfIdOk

`func (o *CancelLocData) GetAmfIdOk() (*string, bool)`

GetAmfIdOk returns a tuple with the AmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfId

`func (o *CancelLocData) SetAmfId(v string)`

SetAmfId sets AmfId field to given value.

### HasAmfId

`func (o *CancelLocData) HasAmfId() bool`

HasAmfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


