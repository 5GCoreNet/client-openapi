# BaselineDnsAit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AitId** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**EcsOption** | Pointer to [**EcsOption**](EcsOption.md) |  | [optional] 
**DnsServerAddressList** | Pointer to [**[]IpAddr**](IpAddr.md) |  | [optional] 

## Methods

### NewBaselineDnsAit

`func NewBaselineDnsAit(aitId string, ) *BaselineDnsAit`

NewBaselineDnsAit instantiates a new BaselineDnsAit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaselineDnsAitWithDefaults

`func NewBaselineDnsAitWithDefaults() *BaselineDnsAit`

NewBaselineDnsAitWithDefaults instantiates a new BaselineDnsAit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAitId

`func (o *BaselineDnsAit) GetAitId() string`

GetAitId returns the AitId field if non-nil, zero value otherwise.

### GetAitIdOk

`func (o *BaselineDnsAit) GetAitIdOk() (*string, bool)`

GetAitIdOk returns a tuple with the AitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAitId

`func (o *BaselineDnsAit) SetAitId(v string)`

SetAitId sets AitId field to given value.


### GetLabel

`func (o *BaselineDnsAit) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BaselineDnsAit) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BaselineDnsAit) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BaselineDnsAit) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetEcsOption

`func (o *BaselineDnsAit) GetEcsOption() EcsOption`

GetEcsOption returns the EcsOption field if non-nil, zero value otherwise.

### GetEcsOptionOk

`func (o *BaselineDnsAit) GetEcsOptionOk() (*EcsOption, bool)`

GetEcsOptionOk returns a tuple with the EcsOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsOption

`func (o *BaselineDnsAit) SetEcsOption(v EcsOption)`

SetEcsOption sets EcsOption field to given value.

### HasEcsOption

`func (o *BaselineDnsAit) HasEcsOption() bool`

HasEcsOption returns a boolean if a field has been set.

### GetDnsServerAddressList

`func (o *BaselineDnsAit) GetDnsServerAddressList() []IpAddr`

GetDnsServerAddressList returns the DnsServerAddressList field if non-nil, zero value otherwise.

### GetDnsServerAddressListOk

`func (o *BaselineDnsAit) GetDnsServerAddressListOk() (*[]IpAddr, bool)`

GetDnsServerAddressListOk returns a tuple with the DnsServerAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerAddressList

`func (o *BaselineDnsAit) SetDnsServerAddressList(v []IpAddr)`

SetDnsServerAddressList sets DnsServerAddressList field to given value.

### HasDnsServerAddressList

`func (o *BaselineDnsAit) HasDnsServerAddressList() bool`

HasDnsServerAddressList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


