# IpSmGwAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpSmGwNumber** | **string** | String containing an additional or basic MSISDN | 
**IpSmGwDiaUri** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**IpSmGwDiaRealm** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**IpSmGwSbiSupInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewIpSmGwAddress

`func NewIpSmGwAddress(ipSmGwNumber string, ) *IpSmGwAddress`

NewIpSmGwAddress instantiates a new IpSmGwAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpSmGwAddressWithDefaults

`func NewIpSmGwAddressWithDefaults() *IpSmGwAddress`

NewIpSmGwAddressWithDefaults instantiates a new IpSmGwAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpSmGwNumber

`func (o *IpSmGwAddress) GetIpSmGwNumber() string`

GetIpSmGwNumber returns the IpSmGwNumber field if non-nil, zero value otherwise.

### GetIpSmGwNumberOk

`func (o *IpSmGwAddress) GetIpSmGwNumberOk() (*string, bool)`

GetIpSmGwNumberOk returns a tuple with the IpSmGwNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSmGwNumber

`func (o *IpSmGwAddress) SetIpSmGwNumber(v string)`

SetIpSmGwNumber sets IpSmGwNumber field to given value.


### GetIpSmGwDiaUri

`func (o *IpSmGwAddress) GetIpSmGwDiaUri() string`

GetIpSmGwDiaUri returns the IpSmGwDiaUri field if non-nil, zero value otherwise.

### GetIpSmGwDiaUriOk

`func (o *IpSmGwAddress) GetIpSmGwDiaUriOk() (*string, bool)`

GetIpSmGwDiaUriOk returns a tuple with the IpSmGwDiaUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSmGwDiaUri

`func (o *IpSmGwAddress) SetIpSmGwDiaUri(v string)`

SetIpSmGwDiaUri sets IpSmGwDiaUri field to given value.

### HasIpSmGwDiaUri

`func (o *IpSmGwAddress) HasIpSmGwDiaUri() bool`

HasIpSmGwDiaUri returns a boolean if a field has been set.

### GetIpSmGwDiaRealm

`func (o *IpSmGwAddress) GetIpSmGwDiaRealm() string`

GetIpSmGwDiaRealm returns the IpSmGwDiaRealm field if non-nil, zero value otherwise.

### GetIpSmGwDiaRealmOk

`func (o *IpSmGwAddress) GetIpSmGwDiaRealmOk() (*string, bool)`

GetIpSmGwDiaRealmOk returns a tuple with the IpSmGwDiaRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSmGwDiaRealm

`func (o *IpSmGwAddress) SetIpSmGwDiaRealm(v string)`

SetIpSmGwDiaRealm sets IpSmGwDiaRealm field to given value.

### HasIpSmGwDiaRealm

`func (o *IpSmGwAddress) HasIpSmGwDiaRealm() bool`

HasIpSmGwDiaRealm returns a boolean if a field has been set.

### GetIpSmGwSbiSupInd

`func (o *IpSmGwAddress) GetIpSmGwSbiSupInd() bool`

GetIpSmGwSbiSupInd returns the IpSmGwSbiSupInd field if non-nil, zero value otherwise.

### GetIpSmGwSbiSupIndOk

`func (o *IpSmGwAddress) GetIpSmGwSbiSupIndOk() (*bool, bool)`

GetIpSmGwSbiSupIndOk returns a tuple with the IpSmGwSbiSupInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSmGwSbiSupInd

`func (o *IpSmGwAddress) SetIpSmGwSbiSupInd(v bool)`

SetIpSmGwSbiSupInd sets IpSmGwSbiSupInd field to given value.

### HasIpSmGwSbiSupInd

`func (o *IpSmGwAddress) HasIpSmGwSbiSupInd() bool`

HasIpSmGwSbiSupInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


