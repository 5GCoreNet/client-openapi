# AkmaKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**AKId** | **string** | Represents an AKMA Key Identifier. | 
**KAkma** | **string** |  | 

## Methods

### NewAkmaKeyInfo

`func NewAkmaKeyInfo(supi string, aKId string, kAkma string, ) *AkmaKeyInfo`

NewAkmaKeyInfo instantiates a new AkmaKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAkmaKeyInfoWithDefaults

`func NewAkmaKeyInfoWithDefaults() *AkmaKeyInfo`

NewAkmaKeyInfoWithDefaults instantiates a new AkmaKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuppFeat

`func (o *AkmaKeyInfo) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AkmaKeyInfo) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AkmaKeyInfo) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AkmaKeyInfo) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetSupi

`func (o *AkmaKeyInfo) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *AkmaKeyInfo) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *AkmaKeyInfo) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetAKId

`func (o *AkmaKeyInfo) GetAKId() string`

GetAKId returns the AKId field if non-nil, zero value otherwise.

### GetAKIdOk

`func (o *AkmaKeyInfo) GetAKIdOk() (*string, bool)`

GetAKIdOk returns a tuple with the AKId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAKId

`func (o *AkmaKeyInfo) SetAKId(v string)`

SetAKId sets AKId field to given value.


### GetKAkma

`func (o *AkmaKeyInfo) GetKAkma() string`

GetKAkma returns the KAkma field if non-nil, zero value otherwise.

### GetKAkmaOk

`func (o *AkmaKeyInfo) GetKAkmaOk() (*string, bool)`

GetKAkmaOk returns a tuple with the KAkma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKAkma

`func (o *AkmaKeyInfo) SetKAkma(v string)`

SetKAkma sets KAkma field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


