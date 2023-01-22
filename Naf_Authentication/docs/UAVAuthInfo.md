# UAVAuthInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**ServiceLevelId** | **string** |  | 
**NotifyUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NotifyCorrId** | Pointer to **string** |  | [optional] 
**IpAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**AuthMsg** | Pointer to **string** |  | [optional] 
**UavLocInfo** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewUAVAuthInfo

`func NewUAVAuthInfo(gpsi string, serviceLevelId string, ) *UAVAuthInfo`

NewUAVAuthInfo instantiates a new UAVAuthInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUAVAuthInfoWithDefaults

`func NewUAVAuthInfoWithDefaults() *UAVAuthInfo`

NewUAVAuthInfoWithDefaults instantiates a new UAVAuthInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *UAVAuthInfo) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *UAVAuthInfo) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *UAVAuthInfo) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.


### GetServiceLevelId

`func (o *UAVAuthInfo) GetServiceLevelId() string`

GetServiceLevelId returns the ServiceLevelId field if non-nil, zero value otherwise.

### GetServiceLevelIdOk

`func (o *UAVAuthInfo) GetServiceLevelIdOk() (*string, bool)`

GetServiceLevelIdOk returns a tuple with the ServiceLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelId

`func (o *UAVAuthInfo) SetServiceLevelId(v string)`

SetServiceLevelId sets ServiceLevelId field to given value.


### GetNotifyUri

`func (o *UAVAuthInfo) GetNotifyUri() string`

GetNotifyUri returns the NotifyUri field if non-nil, zero value otherwise.

### GetNotifyUriOk

`func (o *UAVAuthInfo) GetNotifyUriOk() (*string, bool)`

GetNotifyUriOk returns a tuple with the NotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUri

`func (o *UAVAuthInfo) SetNotifyUri(v string)`

SetNotifyUri sets NotifyUri field to given value.

### HasNotifyUri

`func (o *UAVAuthInfo) HasNotifyUri() bool`

HasNotifyUri returns a boolean if a field has been set.

### GetNotifyCorrId

`func (o *UAVAuthInfo) GetNotifyCorrId() string`

GetNotifyCorrId returns the NotifyCorrId field if non-nil, zero value otherwise.

### GetNotifyCorrIdOk

`func (o *UAVAuthInfo) GetNotifyCorrIdOk() (*string, bool)`

GetNotifyCorrIdOk returns a tuple with the NotifyCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrId

`func (o *UAVAuthInfo) SetNotifyCorrId(v string)`

SetNotifyCorrId sets NotifyCorrId field to given value.

### HasNotifyCorrId

`func (o *UAVAuthInfo) HasNotifyCorrId() bool`

HasNotifyCorrId returns a boolean if a field has been set.

### GetIpAddr

`func (o *UAVAuthInfo) GetIpAddr() IpAddr`

GetIpAddr returns the IpAddr field if non-nil, zero value otherwise.

### GetIpAddrOk

`func (o *UAVAuthInfo) GetIpAddrOk() (*IpAddr, bool)`

GetIpAddrOk returns a tuple with the IpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddr

`func (o *UAVAuthInfo) SetIpAddr(v IpAddr)`

SetIpAddr sets IpAddr field to given value.

### HasIpAddr

`func (o *UAVAuthInfo) HasIpAddr() bool`

HasIpAddr returns a boolean if a field has been set.

### GetPei

`func (o *UAVAuthInfo) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *UAVAuthInfo) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *UAVAuthInfo) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *UAVAuthInfo) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetAuthMsg

`func (o *UAVAuthInfo) GetAuthMsg() string`

GetAuthMsg returns the AuthMsg field if non-nil, zero value otherwise.

### GetAuthMsgOk

`func (o *UAVAuthInfo) GetAuthMsgOk() (*string, bool)`

GetAuthMsgOk returns a tuple with the AuthMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMsg

`func (o *UAVAuthInfo) SetAuthMsg(v string)`

SetAuthMsg sets AuthMsg field to given value.

### HasAuthMsg

`func (o *UAVAuthInfo) HasAuthMsg() bool`

HasAuthMsg returns a boolean if a field has been set.

### GetUavLocInfo

`func (o *UAVAuthInfo) GetUavLocInfo() LocationArea5G`

GetUavLocInfo returns the UavLocInfo field if non-nil, zero value otherwise.

### GetUavLocInfoOk

`func (o *UAVAuthInfo) GetUavLocInfoOk() (*LocationArea5G, bool)`

GetUavLocInfoOk returns a tuple with the UavLocInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUavLocInfo

`func (o *UAVAuthInfo) SetUavLocInfo(v LocationArea5G)`

SetUavLocInfo sets UavLocInfo field to given value.

### HasUavLocInfo

`func (o *UAVAuthInfo) HasUavLocInfo() bool`

HasUavLocInfo returns a boolean if a field has been set.

### GetSuppFeat

`func (o *UAVAuthInfo) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *UAVAuthInfo) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *UAVAuthInfo) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *UAVAuthInfo) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


