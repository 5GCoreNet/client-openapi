# AmInfluData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppIds** | Pointer to **[]string** | Identifies one or more applications. | [optional] 
**DnnSnssaiInfos** | Pointer to [**[]DnnSnssaiInformation**](DnnSnssaiInformation.md) | Identifies one or more DNN, S-NSSAI combinations. | [optional] 
**InterGroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**AnyUeInd** | Pointer to **bool** | Indicates whether the data is applicable for any UE. | [optional] 
**PolicyDuration** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**EvSubs** | Pointer to [**[]AmInfluEvent**](AmInfluEvent.md) | List of AM related events for which a subscription is required. | [optional] 
**NotifUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NotifCorrId** | Pointer to **string** | Notification correlation identifier. | [optional] 
**Headers** | Pointer to **[]string** | Contains the headers provisioned by the NEF. | [optional] 
**ThruReq** | Pointer to **bool** | Indicates whether high throughput is desired for the indicated UE traffic. | [optional] 
**CovReq** | Pointer to [**[]ServiceAreaCoverageInfo**](ServiceAreaCoverageInfo.md) | Indicates the service area coverage requirement. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ResUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAmInfluData

`func NewAmInfluData() *AmInfluData`

NewAmInfluData instantiates a new AmInfluData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmInfluDataWithDefaults

`func NewAmInfluDataWithDefaults() *AmInfluData`

NewAmInfluDataWithDefaults instantiates a new AmInfluData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppIds

`func (o *AmInfluData) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *AmInfluData) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *AmInfluData) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *AmInfluData) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetDnnSnssaiInfos

`func (o *AmInfluData) GetDnnSnssaiInfos() []DnnSnssaiInformation`

GetDnnSnssaiInfos returns the DnnSnssaiInfos field if non-nil, zero value otherwise.

### GetDnnSnssaiInfosOk

`func (o *AmInfluData) GetDnnSnssaiInfosOk() (*[]DnnSnssaiInformation, bool)`

GetDnnSnssaiInfosOk returns a tuple with the DnnSnssaiInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnSnssaiInfos

`func (o *AmInfluData) SetDnnSnssaiInfos(v []DnnSnssaiInformation)`

SetDnnSnssaiInfos sets DnnSnssaiInfos field to given value.

### HasDnnSnssaiInfos

`func (o *AmInfluData) HasDnnSnssaiInfos() bool`

HasDnnSnssaiInfos returns a boolean if a field has been set.

### GetInterGroupId

`func (o *AmInfluData) GetInterGroupId() string`

GetInterGroupId returns the InterGroupId field if non-nil, zero value otherwise.

### GetInterGroupIdOk

`func (o *AmInfluData) GetInterGroupIdOk() (*string, bool)`

GetInterGroupIdOk returns a tuple with the InterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGroupId

`func (o *AmInfluData) SetInterGroupId(v string)`

SetInterGroupId sets InterGroupId field to given value.

### HasInterGroupId

`func (o *AmInfluData) HasInterGroupId() bool`

HasInterGroupId returns a boolean if a field has been set.

### GetSupi

`func (o *AmInfluData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *AmInfluData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *AmInfluData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *AmInfluData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetAnyUeInd

`func (o *AmInfluData) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *AmInfluData) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *AmInfluData) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *AmInfluData) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetPolicyDuration

`func (o *AmInfluData) GetPolicyDuration() int32`

GetPolicyDuration returns the PolicyDuration field if non-nil, zero value otherwise.

### GetPolicyDurationOk

`func (o *AmInfluData) GetPolicyDurationOk() (*int32, bool)`

GetPolicyDurationOk returns a tuple with the PolicyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDuration

`func (o *AmInfluData) SetPolicyDuration(v int32)`

SetPolicyDuration sets PolicyDuration field to given value.

### HasPolicyDuration

`func (o *AmInfluData) HasPolicyDuration() bool`

HasPolicyDuration returns a boolean if a field has been set.

### GetEvSubs

`func (o *AmInfluData) GetEvSubs() []AmInfluEvent`

GetEvSubs returns the EvSubs field if non-nil, zero value otherwise.

### GetEvSubsOk

`func (o *AmInfluData) GetEvSubsOk() (*[]AmInfluEvent, bool)`

GetEvSubsOk returns a tuple with the EvSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvSubs

`func (o *AmInfluData) SetEvSubs(v []AmInfluEvent)`

SetEvSubs sets EvSubs field to given value.

### HasEvSubs

`func (o *AmInfluData) HasEvSubs() bool`

HasEvSubs returns a boolean if a field has been set.

### GetNotifUri

`func (o *AmInfluData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *AmInfluData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *AmInfluData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.

### HasNotifUri

`func (o *AmInfluData) HasNotifUri() bool`

HasNotifUri returns a boolean if a field has been set.

### GetNotifCorrId

`func (o *AmInfluData) GetNotifCorrId() string`

GetNotifCorrId returns the NotifCorrId field if non-nil, zero value otherwise.

### GetNotifCorrIdOk

`func (o *AmInfluData) GetNotifCorrIdOk() (*string, bool)`

GetNotifCorrIdOk returns a tuple with the NotifCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorrId

`func (o *AmInfluData) SetNotifCorrId(v string)`

SetNotifCorrId sets NotifCorrId field to given value.

### HasNotifCorrId

`func (o *AmInfluData) HasNotifCorrId() bool`

HasNotifCorrId returns a boolean if a field has been set.

### GetHeaders

`func (o *AmInfluData) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AmInfluData) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AmInfluData) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AmInfluData) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetThruReq

`func (o *AmInfluData) GetThruReq() bool`

GetThruReq returns the ThruReq field if non-nil, zero value otherwise.

### GetThruReqOk

`func (o *AmInfluData) GetThruReqOk() (*bool, bool)`

GetThruReqOk returns a tuple with the ThruReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThruReq

`func (o *AmInfluData) SetThruReq(v bool)`

SetThruReq sets ThruReq field to given value.

### HasThruReq

`func (o *AmInfluData) HasThruReq() bool`

HasThruReq returns a boolean if a field has been set.

### GetCovReq

`func (o *AmInfluData) GetCovReq() []ServiceAreaCoverageInfo`

GetCovReq returns the CovReq field if non-nil, zero value otherwise.

### GetCovReqOk

`func (o *AmInfluData) GetCovReqOk() (*[]ServiceAreaCoverageInfo, bool)`

GetCovReqOk returns a tuple with the CovReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCovReq

`func (o *AmInfluData) SetCovReq(v []ServiceAreaCoverageInfo)`

SetCovReq sets CovReq field to given value.

### HasCovReq

`func (o *AmInfluData) HasCovReq() bool`

HasCovReq returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *AmInfluData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AmInfluData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AmInfluData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AmInfluData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetResUri

`func (o *AmInfluData) GetResUri() string`

GetResUri returns the ResUri field if non-nil, zero value otherwise.

### GetResUriOk

`func (o *AmInfluData) GetResUriOk() (*string, bool)`

GetResUriOk returns a tuple with the ResUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResUri

`func (o *AmInfluData) SetResUri(v string)`

SetResUri sets ResUri field to given value.

### HasResUri

`func (o *AmInfluData) HasResUri() bool`

HasResUri returns a boolean if a field has been set.

### GetResetIds

`func (o *AmInfluData) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *AmInfluData) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *AmInfluData) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *AmInfluData) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


