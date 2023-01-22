# CpInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**MtcProviderId** | Pointer to **string** | Identifies the MTC Service Provider and/or MTC Application. | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**ExternalId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information. | [optional] 
**Msisdn** | Pointer to **string** | string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN. | [optional] 
**ExternalGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**CpParameterSets** | [**map[string]CpParameterSet**](CpParameterSet.md) | Identifies a set of CP parameter information that may be part of this CpInfo structure. Any string value can be used as a key of the map. | 
**CpReports** | Pointer to [**map[string]CpReport**](CpReport.md) | Supplied by the SCEF and contains the CP set identifiers for which CP parameter(s) are not added or modified successfully. The failure reason is also included. Each element provides the related information for one or more CP set identifier(s) and is identified in the map via the failure identifier as key. | [optional] [readonly] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**UeIpAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 
**UeMacAddr** | Pointer to **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | [optional] 

## Methods

### NewCpInfo

`func NewCpInfo(cpParameterSets map[string]CpParameterSet, ) *CpInfo`

NewCpInfo instantiates a new CpInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpInfoWithDefaults

`func NewCpInfoWithDefaults() *CpInfo`

NewCpInfoWithDefaults instantiates a new CpInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *CpInfo) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *CpInfo) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *CpInfo) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *CpInfo) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *CpInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *CpInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *CpInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *CpInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetMtcProviderId

`func (o *CpInfo) GetMtcProviderId() string`

GetMtcProviderId returns the MtcProviderId field if non-nil, zero value otherwise.

### GetMtcProviderIdOk

`func (o *CpInfo) GetMtcProviderIdOk() (*string, bool)`

GetMtcProviderIdOk returns a tuple with the MtcProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderId

`func (o *CpInfo) SetMtcProviderId(v string)`

SetMtcProviderId sets MtcProviderId field to given value.

### HasMtcProviderId

`func (o *CpInfo) HasMtcProviderId() bool`

HasMtcProviderId returns a boolean if a field has been set.

### GetDnn

`func (o *CpInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *CpInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *CpInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *CpInfo) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetExternalId

`func (o *CpInfo) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CpInfo) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CpInfo) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CpInfo) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMsisdn

`func (o *CpInfo) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *CpInfo) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *CpInfo) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *CpInfo) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetExternalGroupId

`func (o *CpInfo) GetExternalGroupId() string`

GetExternalGroupId returns the ExternalGroupId field if non-nil, zero value otherwise.

### GetExternalGroupIdOk

`func (o *CpInfo) GetExternalGroupIdOk() (*string, bool)`

GetExternalGroupIdOk returns a tuple with the ExternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupId

`func (o *CpInfo) SetExternalGroupId(v string)`

SetExternalGroupId sets ExternalGroupId field to given value.

### HasExternalGroupId

`func (o *CpInfo) HasExternalGroupId() bool`

HasExternalGroupId returns a boolean if a field has been set.

### GetCpParameterSets

`func (o *CpInfo) GetCpParameterSets() map[string]CpParameterSet`

GetCpParameterSets returns the CpParameterSets field if non-nil, zero value otherwise.

### GetCpParameterSetsOk

`func (o *CpInfo) GetCpParameterSetsOk() (*map[string]CpParameterSet, bool)`

GetCpParameterSetsOk returns a tuple with the CpParameterSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpParameterSets

`func (o *CpInfo) SetCpParameterSets(v map[string]CpParameterSet)`

SetCpParameterSets sets CpParameterSets field to given value.


### GetCpReports

`func (o *CpInfo) GetCpReports() map[string]CpReport`

GetCpReports returns the CpReports field if non-nil, zero value otherwise.

### GetCpReportsOk

`func (o *CpInfo) GetCpReportsOk() (*map[string]CpReport, bool)`

GetCpReportsOk returns a tuple with the CpReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpReports

`func (o *CpInfo) SetCpReports(v map[string]CpReport)`

SetCpReports sets CpReports field to given value.

### HasCpReports

`func (o *CpInfo) HasCpReports() bool`

HasCpReports returns a boolean if a field has been set.

### GetSnssai

`func (o *CpInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *CpInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *CpInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *CpInfo) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetUeIpAddr

`func (o *CpInfo) GetUeIpAddr() IpAddr`

GetUeIpAddr returns the UeIpAddr field if non-nil, zero value otherwise.

### GetUeIpAddrOk

`func (o *CpInfo) GetUeIpAddrOk() (*IpAddr, bool)`

GetUeIpAddrOk returns a tuple with the UeIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddr

`func (o *CpInfo) SetUeIpAddr(v IpAddr)`

SetUeIpAddr sets UeIpAddr field to given value.

### HasUeIpAddr

`func (o *CpInfo) HasUeIpAddr() bool`

HasUeIpAddr returns a boolean if a field has been set.

### GetUeMacAddr

`func (o *CpInfo) GetUeMacAddr() string`

GetUeMacAddr returns the UeMacAddr field if non-nil, zero value otherwise.

### GetUeMacAddrOk

`func (o *CpInfo) GetUeMacAddrOk() (*string, bool)`

GetUeMacAddrOk returns a tuple with the UeMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMacAddr

`func (o *CpInfo) SetUeMacAddr(v string)`

SetUeMacAddr sets UeMacAddr field to given value.

### HasUeMacAddr

`func (o *CpInfo) HasUeMacAddr() bool`

HasUeMacAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


