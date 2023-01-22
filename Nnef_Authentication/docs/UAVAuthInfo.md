# UAVAuthInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**ServiceLevelId** | **string** |  | 
**AuthNotificationURI** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**IpAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**AuthServerAddress** | Pointer to **string** |  | [optional] 
**AuthMsg** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**UeLocInfo** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**SNssai** | Pointer to [**ExtSnssai**](ExtSnssai.md) |  | [optional] 
**NfType** | [**NFType**](NFType.md) |  | 

## Methods

### NewUAVAuthInfo

`func NewUAVAuthInfo(gpsi string, serviceLevelId string, nfType NFType, ) *UAVAuthInfo`

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


### GetAuthNotificationURI

`func (o *UAVAuthInfo) GetAuthNotificationURI() string`

GetAuthNotificationURI returns the AuthNotificationURI field if non-nil, zero value otherwise.

### GetAuthNotificationURIOk

`func (o *UAVAuthInfo) GetAuthNotificationURIOk() (*string, bool)`

GetAuthNotificationURIOk returns a tuple with the AuthNotificationURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthNotificationURI

`func (o *UAVAuthInfo) SetAuthNotificationURI(v string)`

SetAuthNotificationURI sets AuthNotificationURI field to given value.

### HasAuthNotificationURI

`func (o *UAVAuthInfo) HasAuthNotificationURI() bool`

HasAuthNotificationURI returns a boolean if a field has been set.

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

### GetAuthServerAddress

`func (o *UAVAuthInfo) GetAuthServerAddress() string`

GetAuthServerAddress returns the AuthServerAddress field if non-nil, zero value otherwise.

### GetAuthServerAddressOk

`func (o *UAVAuthInfo) GetAuthServerAddressOk() (*string, bool)`

GetAuthServerAddressOk returns a tuple with the AuthServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServerAddress

`func (o *UAVAuthInfo) SetAuthServerAddress(v string)`

SetAuthServerAddress sets AuthServerAddress field to given value.

### HasAuthServerAddress

`func (o *UAVAuthInfo) HasAuthServerAddress() bool`

HasAuthServerAddress returns a boolean if a field has been set.

### GetAuthMsg

`func (o *UAVAuthInfo) GetAuthMsg() RefToBinaryData`

GetAuthMsg returns the AuthMsg field if non-nil, zero value otherwise.

### GetAuthMsgOk

`func (o *UAVAuthInfo) GetAuthMsgOk() (*RefToBinaryData, bool)`

GetAuthMsgOk returns a tuple with the AuthMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMsg

`func (o *UAVAuthInfo) SetAuthMsg(v RefToBinaryData)`

SetAuthMsg sets AuthMsg field to given value.

### HasAuthMsg

`func (o *UAVAuthInfo) HasAuthMsg() bool`

HasAuthMsg returns a boolean if a field has been set.

### GetUeLocInfo

`func (o *UAVAuthInfo) GetUeLocInfo() UserLocation`

GetUeLocInfo returns the UeLocInfo field if non-nil, zero value otherwise.

### GetUeLocInfoOk

`func (o *UAVAuthInfo) GetUeLocInfoOk() (*UserLocation, bool)`

GetUeLocInfoOk returns a tuple with the UeLocInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocInfo

`func (o *UAVAuthInfo) SetUeLocInfo(v UserLocation)`

SetUeLocInfo sets UeLocInfo field to given value.

### HasUeLocInfo

`func (o *UAVAuthInfo) HasUeLocInfo() bool`

HasUeLocInfo returns a boolean if a field has been set.

### GetDnn

`func (o *UAVAuthInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *UAVAuthInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *UAVAuthInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *UAVAuthInfo) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSNssai

`func (o *UAVAuthInfo) GetSNssai() ExtSnssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *UAVAuthInfo) GetSNssaiOk() (*ExtSnssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *UAVAuthInfo) SetSNssai(v ExtSnssai)`

SetSNssai sets SNssai field to given value.

### HasSNssai

`func (o *UAVAuthInfo) HasSNssai() bool`

HasSNssai returns a boolean if a field has been set.

### GetNfType

`func (o *UAVAuthInfo) GetNfType() NFType`

GetNfType returns the NfType field if non-nil, zero value otherwise.

### GetNfTypeOk

`func (o *UAVAuthInfo) GetNfTypeOk() (*NFType, bool)`

GetNfTypeOk returns a tuple with the NfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfType

`func (o *UAVAuthInfo) SetNfType(v NFType)`

SetNfType sets NfType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


