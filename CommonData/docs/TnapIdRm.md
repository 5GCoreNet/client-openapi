# TnapIdRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsId** | Pointer to **string** | This IE shall be present if the UE is accessing the 5GC via a trusted WLAN access network.When present, it shall contain the SSID of the access point to which the UE is attached, that is received over NGAP,  see IEEE Std 802.11-2012.   | [optional] 
**BssId** | Pointer to **string** | When present, it shall contain the BSSID of the access point to which the UE is attached, that is received over NGAP, see IEEE Std 802.11-2012.   | [optional] 
**CivicAddress** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 

## Methods

### NewTnapIdRm

`func NewTnapIdRm() *TnapIdRm`

NewTnapIdRm instantiates a new TnapIdRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTnapIdRmWithDefaults

`func NewTnapIdRmWithDefaults() *TnapIdRm`

NewTnapIdRmWithDefaults instantiates a new TnapIdRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsId

`func (o *TnapIdRm) GetSsId() string`

GetSsId returns the SsId field if non-nil, zero value otherwise.

### GetSsIdOk

`func (o *TnapIdRm) GetSsIdOk() (*string, bool)`

GetSsIdOk returns a tuple with the SsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsId

`func (o *TnapIdRm) SetSsId(v string)`

SetSsId sets SsId field to given value.

### HasSsId

`func (o *TnapIdRm) HasSsId() bool`

HasSsId returns a boolean if a field has been set.

### GetBssId

`func (o *TnapIdRm) GetBssId() string`

GetBssId returns the BssId field if non-nil, zero value otherwise.

### GetBssIdOk

`func (o *TnapIdRm) GetBssIdOk() (*string, bool)`

GetBssIdOk returns a tuple with the BssId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssId

`func (o *TnapIdRm) SetBssId(v string)`

SetBssId sets BssId field to given value.

### HasBssId

`func (o *TnapIdRm) HasBssId() bool`

HasBssId returns a boolean if a field has been set.

### GetCivicAddress

`func (o *TnapIdRm) GetCivicAddress() string`

GetCivicAddress returns the CivicAddress field if non-nil, zero value otherwise.

### GetCivicAddressOk

`func (o *TnapIdRm) GetCivicAddressOk() (*string, bool)`

GetCivicAddressOk returns a tuple with the CivicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddress

`func (o *TnapIdRm) SetCivicAddress(v string)`

SetCivicAddress sets CivicAddress field to given value.

### HasCivicAddress

`func (o *TnapIdRm) HasCivicAddress() bool`

HasCivicAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


