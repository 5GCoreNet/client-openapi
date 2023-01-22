# TwapIdRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsId** | **string** | This IE shall contain the SSID of the access point to which the UE is attached, that is received over NGAP, see IEEE Std 802.11-2012.   | 
**BssId** | Pointer to **string** | When present, it shall contain the BSSID of the access point to which the UE is attached, for trusted WLAN access, see IEEE Std 802.11-2012.   | [optional] 
**CivicAddress** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 

## Methods

### NewTwapIdRm

`func NewTwapIdRm(ssId string, ) *TwapIdRm`

NewTwapIdRm instantiates a new TwapIdRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwapIdRmWithDefaults

`func NewTwapIdRmWithDefaults() *TwapIdRm`

NewTwapIdRmWithDefaults instantiates a new TwapIdRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsId

`func (o *TwapIdRm) GetSsId() string`

GetSsId returns the SsId field if non-nil, zero value otherwise.

### GetSsIdOk

`func (o *TwapIdRm) GetSsIdOk() (*string, bool)`

GetSsIdOk returns a tuple with the SsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsId

`func (o *TwapIdRm) SetSsId(v string)`

SetSsId sets SsId field to given value.


### GetBssId

`func (o *TwapIdRm) GetBssId() string`

GetBssId returns the BssId field if non-nil, zero value otherwise.

### GetBssIdOk

`func (o *TwapIdRm) GetBssIdOk() (*string, bool)`

GetBssIdOk returns a tuple with the BssId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssId

`func (o *TwapIdRm) SetBssId(v string)`

SetBssId sets BssId field to given value.

### HasBssId

`func (o *TwapIdRm) HasBssId() bool`

HasBssId returns a boolean if a field has been set.

### GetCivicAddress

`func (o *TwapIdRm) GetCivicAddress() string`

GetCivicAddress returns the CivicAddress field if non-nil, zero value otherwise.

### GetCivicAddressOk

`func (o *TwapIdRm) GetCivicAddressOk() (*string, bool)`

GetCivicAddressOk returns a tuple with the CivicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddress

`func (o *TwapIdRm) SetCivicAddress(v string)`

SetCivicAddress sets CivicAddress field to given value.

### HasCivicAddress

`func (o *TwapIdRm) HasCivicAddress() bool`

HasCivicAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


