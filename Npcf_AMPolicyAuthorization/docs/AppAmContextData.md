# AppAmContextData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**TermNotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**EvSubsc** | Pointer to [**AmEventsSubscData**](AmEventsSubscData.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**Expiry** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**HighThruInd** | Pointer to **bool** | Indicates whether high throughput is desired for the indicated UE traffic. | [optional] 
**CovReq** | Pointer to [**[]ServiceAreaCoverageInfo**](ServiceAreaCoverageInfo.md) | Identifies a list of Tracking Areas per serving network where service is allowed. | [optional] 
**AsTimeDisParam** | Pointer to [**NullableAsTimeDistributionParam**](AsTimeDistributionParam.md) |  | [optional] 

## Methods

### NewAppAmContextData

`func NewAppAmContextData(supi string, termNotifUri string, ) *AppAmContextData`

NewAppAmContextData instantiates a new AppAmContextData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAmContextDataWithDefaults

`func NewAppAmContextDataWithDefaults() *AppAmContextData`

NewAppAmContextDataWithDefaults instantiates a new AppAmContextData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *AppAmContextData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *AppAmContextData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *AppAmContextData) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetGpsi

`func (o *AppAmContextData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *AppAmContextData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *AppAmContextData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *AppAmContextData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetTermNotifUri

`func (o *AppAmContextData) GetTermNotifUri() string`

GetTermNotifUri returns the TermNotifUri field if non-nil, zero value otherwise.

### GetTermNotifUriOk

`func (o *AppAmContextData) GetTermNotifUriOk() (*string, bool)`

GetTermNotifUriOk returns a tuple with the TermNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermNotifUri

`func (o *AppAmContextData) SetTermNotifUri(v string)`

SetTermNotifUri sets TermNotifUri field to given value.


### GetEvSubsc

`func (o *AppAmContextData) GetEvSubsc() AmEventsSubscData`

GetEvSubsc returns the EvSubsc field if non-nil, zero value otherwise.

### GetEvSubscOk

`func (o *AppAmContextData) GetEvSubscOk() (*AmEventsSubscData, bool)`

GetEvSubscOk returns a tuple with the EvSubsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvSubsc

`func (o *AppAmContextData) SetEvSubsc(v AmEventsSubscData)`

SetEvSubsc sets EvSubsc field to given value.

### HasEvSubsc

`func (o *AppAmContextData) HasEvSubsc() bool`

HasEvSubsc returns a boolean if a field has been set.

### GetSuppFeat

`func (o *AppAmContextData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AppAmContextData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AppAmContextData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AppAmContextData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetExpiry

`func (o *AppAmContextData) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *AppAmContextData) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *AppAmContextData) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *AppAmContextData) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetHighThruInd

`func (o *AppAmContextData) GetHighThruInd() bool`

GetHighThruInd returns the HighThruInd field if non-nil, zero value otherwise.

### GetHighThruIndOk

`func (o *AppAmContextData) GetHighThruIndOk() (*bool, bool)`

GetHighThruIndOk returns a tuple with the HighThruInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighThruInd

`func (o *AppAmContextData) SetHighThruInd(v bool)`

SetHighThruInd sets HighThruInd field to given value.

### HasHighThruInd

`func (o *AppAmContextData) HasHighThruInd() bool`

HasHighThruInd returns a boolean if a field has been set.

### GetCovReq

`func (o *AppAmContextData) GetCovReq() []ServiceAreaCoverageInfo`

GetCovReq returns the CovReq field if non-nil, zero value otherwise.

### GetCovReqOk

`func (o *AppAmContextData) GetCovReqOk() (*[]ServiceAreaCoverageInfo, bool)`

GetCovReqOk returns a tuple with the CovReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCovReq

`func (o *AppAmContextData) SetCovReq(v []ServiceAreaCoverageInfo)`

SetCovReq sets CovReq field to given value.

### HasCovReq

`func (o *AppAmContextData) HasCovReq() bool`

HasCovReq returns a boolean if a field has been set.

### GetAsTimeDisParam

`func (o *AppAmContextData) GetAsTimeDisParam() AsTimeDistributionParam`

GetAsTimeDisParam returns the AsTimeDisParam field if non-nil, zero value otherwise.

### GetAsTimeDisParamOk

`func (o *AppAmContextData) GetAsTimeDisParamOk() (*AsTimeDistributionParam, bool)`

GetAsTimeDisParamOk returns a tuple with the AsTimeDisParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTimeDisParam

`func (o *AppAmContextData) SetAsTimeDisParam(v AsTimeDistributionParam)`

SetAsTimeDisParam sets AsTimeDisParam field to given value.

### HasAsTimeDisParam

`func (o *AppAmContextData) HasAsTimeDisParam() bool`

HasAsTimeDisParam returns a boolean if a field has been set.

### SetAsTimeDisParamNil

`func (o *AppAmContextData) SetAsTimeDisParamNil(b bool)`

 SetAsTimeDisParamNil sets the value for AsTimeDisParam to be an explicit nil

### UnsetAsTimeDisParam
`func (o *AppAmContextData) UnsetAsTimeDisParam()`

UnsetAsTimeDisParam ensures that no value is present for AsTimeDisParam, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


