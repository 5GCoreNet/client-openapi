# AkmaAfKeyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Expiry** | **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | 
**Kaf** | **string** |  | 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 

## Methods

### NewAkmaAfKeyData

`func NewAkmaAfKeyData(expiry time.Time, kaf string, ) *AkmaAfKeyData`

NewAkmaAfKeyData instantiates a new AkmaAfKeyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAkmaAfKeyDataWithDefaults

`func NewAkmaAfKeyDataWithDefaults() *AkmaAfKeyData`

NewAkmaAfKeyDataWithDefaults instantiates a new AkmaAfKeyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuppFeat

`func (o *AkmaAfKeyData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AkmaAfKeyData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AkmaAfKeyData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AkmaAfKeyData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetGpsi

`func (o *AkmaAfKeyData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *AkmaAfKeyData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *AkmaAfKeyData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *AkmaAfKeyData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetExpiry

`func (o *AkmaAfKeyData) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *AkmaAfKeyData) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *AkmaAfKeyData) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.


### GetKaf

`func (o *AkmaAfKeyData) GetKaf() string`

GetKaf returns the Kaf field if non-nil, zero value otherwise.

### GetKafOk

`func (o *AkmaAfKeyData) GetKafOk() (*string, bool)`

GetKafOk returns a tuple with the Kaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKaf

`func (o *AkmaAfKeyData) SetKaf(v string)`

SetKaf sets Kaf field to given value.


### GetSupi

`func (o *AkmaAfKeyData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *AkmaAfKeyData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *AkmaAfKeyData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *AkmaAfKeyData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


