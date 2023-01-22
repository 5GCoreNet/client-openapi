# SmsRegistrationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpSmGwNumber** | **string** | String containing an additional or basic MSISDN | 
**ScAddress** | Pointer to **string** | String containing an additional or basic MSISDN | [optional] 

## Methods

### NewSmsRegistrationInfo

`func NewSmsRegistrationInfo(ipSmGwNumber string, ) *SmsRegistrationInfo`

NewSmsRegistrationInfo instantiates a new SmsRegistrationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsRegistrationInfoWithDefaults

`func NewSmsRegistrationInfoWithDefaults() *SmsRegistrationInfo`

NewSmsRegistrationInfoWithDefaults instantiates a new SmsRegistrationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpSmGwNumber

`func (o *SmsRegistrationInfo) GetIpSmGwNumber() string`

GetIpSmGwNumber returns the IpSmGwNumber field if non-nil, zero value otherwise.

### GetIpSmGwNumberOk

`func (o *SmsRegistrationInfo) GetIpSmGwNumberOk() (*string, bool)`

GetIpSmGwNumberOk returns a tuple with the IpSmGwNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSmGwNumber

`func (o *SmsRegistrationInfo) SetIpSmGwNumber(v string)`

SetIpSmGwNumber sets IpSmGwNumber field to given value.


### GetScAddress

`func (o *SmsRegistrationInfo) GetScAddress() string`

GetScAddress returns the ScAddress field if non-nil, zero value otherwise.

### GetScAddressOk

`func (o *SmsRegistrationInfo) GetScAddressOk() (*string, bool)`

GetScAddressOk returns a tuple with the ScAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScAddress

`func (o *SmsRegistrationInfo) SetScAddress(v string)`

SetScAddress sets ScAddress field to given value.

### HasScAddress

`func (o *SmsRegistrationInfo) HasScAddress() bool`

HasScAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


