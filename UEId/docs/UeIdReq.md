# UeIdReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfId** | **string** |  | 
**AppPortId** | Pointer to **int32** | Unsigned integer with valid values between 0 and 65535. | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**IpDomain** | Pointer to **string** |  | [optional] 
**MtcProviderId** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**UeIpAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 
**UeMacAddr** | Pointer to **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | [optional] 

## Methods

### NewUeIdReq

`func NewUeIdReq(afId string, ) *UeIdReq`

NewUeIdReq instantiates a new UeIdReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeIdReqWithDefaults

`func NewUeIdReqWithDefaults() *UeIdReq`

NewUeIdReqWithDefaults instantiates a new UeIdReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfId

`func (o *UeIdReq) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *UeIdReq) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *UeIdReq) SetAfId(v string)`

SetAfId sets AfId field to given value.


### GetAppPortId

`func (o *UeIdReq) GetAppPortId() int32`

GetAppPortId returns the AppPortId field if non-nil, zero value otherwise.

### GetAppPortIdOk

`func (o *UeIdReq) GetAppPortIdOk() (*int32, bool)`

GetAppPortIdOk returns a tuple with the AppPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPortId

`func (o *UeIdReq) SetAppPortId(v int32)`

SetAppPortId sets AppPortId field to given value.

### HasAppPortId

`func (o *UeIdReq) HasAppPortId() bool`

HasAppPortId returns a boolean if a field has been set.

### GetDnn

`func (o *UeIdReq) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *UeIdReq) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *UeIdReq) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *UeIdReq) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetIpDomain

`func (o *UeIdReq) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *UeIdReq) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *UeIdReq) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *UeIdReq) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetMtcProviderId

`func (o *UeIdReq) GetMtcProviderId() string`

GetMtcProviderId returns the MtcProviderId field if non-nil, zero value otherwise.

### GetMtcProviderIdOk

`func (o *UeIdReq) GetMtcProviderIdOk() (*string, bool)`

GetMtcProviderIdOk returns a tuple with the MtcProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderId

`func (o *UeIdReq) SetMtcProviderId(v string)`

SetMtcProviderId sets MtcProviderId field to given value.

### HasMtcProviderId

`func (o *UeIdReq) HasMtcProviderId() bool`

HasMtcProviderId returns a boolean if a field has been set.

### GetSnssai

`func (o *UeIdReq) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *UeIdReq) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *UeIdReq) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *UeIdReq) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetUeIpAddr

`func (o *UeIdReq) GetUeIpAddr() IpAddr`

GetUeIpAddr returns the UeIpAddr field if non-nil, zero value otherwise.

### GetUeIpAddrOk

`func (o *UeIdReq) GetUeIpAddrOk() (*IpAddr, bool)`

GetUeIpAddrOk returns a tuple with the UeIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddr

`func (o *UeIdReq) SetUeIpAddr(v IpAddr)`

SetUeIpAddr sets UeIpAddr field to given value.

### HasUeIpAddr

`func (o *UeIdReq) HasUeIpAddr() bool`

HasUeIpAddr returns a boolean if a field has been set.

### GetUeMacAddr

`func (o *UeIdReq) GetUeMacAddr() string`

GetUeMacAddr returns the UeMacAddr field if non-nil, zero value otherwise.

### GetUeMacAddrOk

`func (o *UeIdReq) GetUeMacAddrOk() (*string, bool)`

GetUeMacAddrOk returns a tuple with the UeMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMacAddr

`func (o *UeIdReq) SetUeMacAddr(v string)`

SetUeMacAddr sets UeMacAddr field to given value.

### HasUeMacAddr

`func (o *UeIdReq) HasUeMacAddr() bool`

HasUeMacAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


