# AuthorizedNetworkSliceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedNssaiList** | Pointer to [**[]AllowedNssai**](AllowedNssai.md) |  | [optional] 
**ConfiguredNssai** | Pointer to [**[]ConfiguredSnssai**](ConfiguredSnssai.md) |  | [optional] 
**TargetAmfSet** | Pointer to **string** |  | [optional] 
**CandidateAmfList** | Pointer to **[]string** |  | [optional] 
**RejectedNssaiInPlmn** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**RejectedNssaiInTa** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**NsiInformation** | Pointer to [**NsiInformation**](NsiInformation.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**NrfAmfSet** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NrfAmfSetNfMgtUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NrfAmfSetAccessTokenUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NrfOauth2Required** | Pointer to **map[string]bool** | Map indicating whether the NRF requires Oauth2-based authorization for accessing its services. The key of the map shall be the name of an NRF service, e.g. \&quot;nnrf-nfm\&quot; or \&quot;nnrf-disc\&quot;  | [optional] 
**TargetAmfServiceSet** | Pointer to **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \&quot;set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)   &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NID&gt; encoded as defined in clause 5.4.2 (\&quot;Nid\&quot; data type definition)  &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2  &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510  &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.  | [optional] 
**TargetNssai** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**NsagInfos** | Pointer to [**[]NsagInfo**](NsagInfo.md) |  | [optional] 

## Methods

### NewAuthorizedNetworkSliceInfo

`func NewAuthorizedNetworkSliceInfo() *AuthorizedNetworkSliceInfo`

NewAuthorizedNetworkSliceInfo instantiates a new AuthorizedNetworkSliceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizedNetworkSliceInfoWithDefaults

`func NewAuthorizedNetworkSliceInfoWithDefaults() *AuthorizedNetworkSliceInfo`

NewAuthorizedNetworkSliceInfoWithDefaults instantiates a new AuthorizedNetworkSliceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedNssaiList

`func (o *AuthorizedNetworkSliceInfo) GetAllowedNssaiList() []AllowedNssai`

GetAllowedNssaiList returns the AllowedNssaiList field if non-nil, zero value otherwise.

### GetAllowedNssaiListOk

`func (o *AuthorizedNetworkSliceInfo) GetAllowedNssaiListOk() (*[]AllowedNssai, bool)`

GetAllowedNssaiListOk returns a tuple with the AllowedNssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNssaiList

`func (o *AuthorizedNetworkSliceInfo) SetAllowedNssaiList(v []AllowedNssai)`

SetAllowedNssaiList sets AllowedNssaiList field to given value.

### HasAllowedNssaiList

`func (o *AuthorizedNetworkSliceInfo) HasAllowedNssaiList() bool`

HasAllowedNssaiList returns a boolean if a field has been set.

### GetConfiguredNssai

`func (o *AuthorizedNetworkSliceInfo) GetConfiguredNssai() []ConfiguredSnssai`

GetConfiguredNssai returns the ConfiguredNssai field if non-nil, zero value otherwise.

### GetConfiguredNssaiOk

`func (o *AuthorizedNetworkSliceInfo) GetConfiguredNssaiOk() (*[]ConfiguredSnssai, bool)`

GetConfiguredNssaiOk returns a tuple with the ConfiguredNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredNssai

`func (o *AuthorizedNetworkSliceInfo) SetConfiguredNssai(v []ConfiguredSnssai)`

SetConfiguredNssai sets ConfiguredNssai field to given value.

### HasConfiguredNssai

`func (o *AuthorizedNetworkSliceInfo) HasConfiguredNssai() bool`

HasConfiguredNssai returns a boolean if a field has been set.

### GetTargetAmfSet

`func (o *AuthorizedNetworkSliceInfo) GetTargetAmfSet() string`

GetTargetAmfSet returns the TargetAmfSet field if non-nil, zero value otherwise.

### GetTargetAmfSetOk

`func (o *AuthorizedNetworkSliceInfo) GetTargetAmfSetOk() (*string, bool)`

GetTargetAmfSetOk returns a tuple with the TargetAmfSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmfSet

`func (o *AuthorizedNetworkSliceInfo) SetTargetAmfSet(v string)`

SetTargetAmfSet sets TargetAmfSet field to given value.

### HasTargetAmfSet

`func (o *AuthorizedNetworkSliceInfo) HasTargetAmfSet() bool`

HasTargetAmfSet returns a boolean if a field has been set.

### GetCandidateAmfList

`func (o *AuthorizedNetworkSliceInfo) GetCandidateAmfList() []string`

GetCandidateAmfList returns the CandidateAmfList field if non-nil, zero value otherwise.

### GetCandidateAmfListOk

`func (o *AuthorizedNetworkSliceInfo) GetCandidateAmfListOk() (*[]string, bool)`

GetCandidateAmfListOk returns a tuple with the CandidateAmfList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateAmfList

`func (o *AuthorizedNetworkSliceInfo) SetCandidateAmfList(v []string)`

SetCandidateAmfList sets CandidateAmfList field to given value.

### HasCandidateAmfList

`func (o *AuthorizedNetworkSliceInfo) HasCandidateAmfList() bool`

HasCandidateAmfList returns a boolean if a field has been set.

### GetRejectedNssaiInPlmn

`func (o *AuthorizedNetworkSliceInfo) GetRejectedNssaiInPlmn() []Snssai`

GetRejectedNssaiInPlmn returns the RejectedNssaiInPlmn field if non-nil, zero value otherwise.

### GetRejectedNssaiInPlmnOk

`func (o *AuthorizedNetworkSliceInfo) GetRejectedNssaiInPlmnOk() (*[]Snssai, bool)`

GetRejectedNssaiInPlmnOk returns a tuple with the RejectedNssaiInPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedNssaiInPlmn

`func (o *AuthorizedNetworkSliceInfo) SetRejectedNssaiInPlmn(v []Snssai)`

SetRejectedNssaiInPlmn sets RejectedNssaiInPlmn field to given value.

### HasRejectedNssaiInPlmn

`func (o *AuthorizedNetworkSliceInfo) HasRejectedNssaiInPlmn() bool`

HasRejectedNssaiInPlmn returns a boolean if a field has been set.

### GetRejectedNssaiInTa

`func (o *AuthorizedNetworkSliceInfo) GetRejectedNssaiInTa() []Snssai`

GetRejectedNssaiInTa returns the RejectedNssaiInTa field if non-nil, zero value otherwise.

### GetRejectedNssaiInTaOk

`func (o *AuthorizedNetworkSliceInfo) GetRejectedNssaiInTaOk() (*[]Snssai, bool)`

GetRejectedNssaiInTaOk returns a tuple with the RejectedNssaiInTa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedNssaiInTa

`func (o *AuthorizedNetworkSliceInfo) SetRejectedNssaiInTa(v []Snssai)`

SetRejectedNssaiInTa sets RejectedNssaiInTa field to given value.

### HasRejectedNssaiInTa

`func (o *AuthorizedNetworkSliceInfo) HasRejectedNssaiInTa() bool`

HasRejectedNssaiInTa returns a boolean if a field has been set.

### GetNsiInformation

`func (o *AuthorizedNetworkSliceInfo) GetNsiInformation() NsiInformation`

GetNsiInformation returns the NsiInformation field if non-nil, zero value otherwise.

### GetNsiInformationOk

`func (o *AuthorizedNetworkSliceInfo) GetNsiInformationOk() (*NsiInformation, bool)`

GetNsiInformationOk returns a tuple with the NsiInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiInformation

`func (o *AuthorizedNetworkSliceInfo) SetNsiInformation(v NsiInformation)`

SetNsiInformation sets NsiInformation field to given value.

### HasNsiInformation

`func (o *AuthorizedNetworkSliceInfo) HasNsiInformation() bool`

HasNsiInformation returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *AuthorizedNetworkSliceInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AuthorizedNetworkSliceInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AuthorizedNetworkSliceInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AuthorizedNetworkSliceInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetNrfAmfSet

`func (o *AuthorizedNetworkSliceInfo) GetNrfAmfSet() string`

GetNrfAmfSet returns the NrfAmfSet field if non-nil, zero value otherwise.

### GetNrfAmfSetOk

`func (o *AuthorizedNetworkSliceInfo) GetNrfAmfSetOk() (*string, bool)`

GetNrfAmfSetOk returns a tuple with the NrfAmfSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfAmfSet

`func (o *AuthorizedNetworkSliceInfo) SetNrfAmfSet(v string)`

SetNrfAmfSet sets NrfAmfSet field to given value.

### HasNrfAmfSet

`func (o *AuthorizedNetworkSliceInfo) HasNrfAmfSet() bool`

HasNrfAmfSet returns a boolean if a field has been set.

### GetNrfAmfSetNfMgtUri

`func (o *AuthorizedNetworkSliceInfo) GetNrfAmfSetNfMgtUri() string`

GetNrfAmfSetNfMgtUri returns the NrfAmfSetNfMgtUri field if non-nil, zero value otherwise.

### GetNrfAmfSetNfMgtUriOk

`func (o *AuthorizedNetworkSliceInfo) GetNrfAmfSetNfMgtUriOk() (*string, bool)`

GetNrfAmfSetNfMgtUriOk returns a tuple with the NrfAmfSetNfMgtUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfAmfSetNfMgtUri

`func (o *AuthorizedNetworkSliceInfo) SetNrfAmfSetNfMgtUri(v string)`

SetNrfAmfSetNfMgtUri sets NrfAmfSetNfMgtUri field to given value.

### HasNrfAmfSetNfMgtUri

`func (o *AuthorizedNetworkSliceInfo) HasNrfAmfSetNfMgtUri() bool`

HasNrfAmfSetNfMgtUri returns a boolean if a field has been set.

### GetNrfAmfSetAccessTokenUri

`func (o *AuthorizedNetworkSliceInfo) GetNrfAmfSetAccessTokenUri() string`

GetNrfAmfSetAccessTokenUri returns the NrfAmfSetAccessTokenUri field if non-nil, zero value otherwise.

### GetNrfAmfSetAccessTokenUriOk

`func (o *AuthorizedNetworkSliceInfo) GetNrfAmfSetAccessTokenUriOk() (*string, bool)`

GetNrfAmfSetAccessTokenUriOk returns a tuple with the NrfAmfSetAccessTokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfAmfSetAccessTokenUri

`func (o *AuthorizedNetworkSliceInfo) SetNrfAmfSetAccessTokenUri(v string)`

SetNrfAmfSetAccessTokenUri sets NrfAmfSetAccessTokenUri field to given value.

### HasNrfAmfSetAccessTokenUri

`func (o *AuthorizedNetworkSliceInfo) HasNrfAmfSetAccessTokenUri() bool`

HasNrfAmfSetAccessTokenUri returns a boolean if a field has been set.

### GetNrfOauth2Required

`func (o *AuthorizedNetworkSliceInfo) GetNrfOauth2Required() map[string]bool`

GetNrfOauth2Required returns the NrfOauth2Required field if non-nil, zero value otherwise.

### GetNrfOauth2RequiredOk

`func (o *AuthorizedNetworkSliceInfo) GetNrfOauth2RequiredOk() (*map[string]bool, bool)`

GetNrfOauth2RequiredOk returns a tuple with the NrfOauth2Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfOauth2Required

`func (o *AuthorizedNetworkSliceInfo) SetNrfOauth2Required(v map[string]bool)`

SetNrfOauth2Required sets NrfOauth2Required field to given value.

### HasNrfOauth2Required

`func (o *AuthorizedNetworkSliceInfo) HasNrfOauth2Required() bool`

HasNrfOauth2Required returns a boolean if a field has been set.

### GetTargetAmfServiceSet

`func (o *AuthorizedNetworkSliceInfo) GetTargetAmfServiceSet() string`

GetTargetAmfServiceSet returns the TargetAmfServiceSet field if non-nil, zero value otherwise.

### GetTargetAmfServiceSetOk

`func (o *AuthorizedNetworkSliceInfo) GetTargetAmfServiceSetOk() (*string, bool)`

GetTargetAmfServiceSetOk returns a tuple with the TargetAmfServiceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmfServiceSet

`func (o *AuthorizedNetworkSliceInfo) SetTargetAmfServiceSet(v string)`

SetTargetAmfServiceSet sets TargetAmfServiceSet field to given value.

### HasTargetAmfServiceSet

`func (o *AuthorizedNetworkSliceInfo) HasTargetAmfServiceSet() bool`

HasTargetAmfServiceSet returns a boolean if a field has been set.

### GetTargetNssai

`func (o *AuthorizedNetworkSliceInfo) GetTargetNssai() []Snssai`

GetTargetNssai returns the TargetNssai field if non-nil, zero value otherwise.

### GetTargetNssaiOk

`func (o *AuthorizedNetworkSliceInfo) GetTargetNssaiOk() (*[]Snssai, bool)`

GetTargetNssaiOk returns a tuple with the TargetNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNssai

`func (o *AuthorizedNetworkSliceInfo) SetTargetNssai(v []Snssai)`

SetTargetNssai sets TargetNssai field to given value.

### HasTargetNssai

`func (o *AuthorizedNetworkSliceInfo) HasTargetNssai() bool`

HasTargetNssai returns a boolean if a field has been set.

### GetNsagInfos

`func (o *AuthorizedNetworkSliceInfo) GetNsagInfos() []NsagInfo`

GetNsagInfos returns the NsagInfos field if non-nil, zero value otherwise.

### GetNsagInfosOk

`func (o *AuthorizedNetworkSliceInfo) GetNsagInfosOk() (*[]NsagInfo, bool)`

GetNsagInfosOk returns a tuple with the NsagInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsagInfos

`func (o *AuthorizedNetworkSliceInfo) SetNsagInfos(v []NsagInfo)`

SetNsagInfos sets NsagInfos field to given value.

### HasNsagInfos

`func (o *AuthorizedNetworkSliceInfo) HasNsagInfos() bool`

HasNsagInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


