# SecParamExchRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N32fContextId** | **string** |  | 
**SelectedJweCipherSuite** | Pointer to **string** |  | [optional] 
**SelectedJwsCipherSuite** | Pointer to **string** |  | [optional] 
**SelProtectionPolicyInfo** | Pointer to [**ProtectionPolicy**](ProtectionPolicy.md) |  | [optional] 
**IpxProviderSecInfoList** | Pointer to [**[]IpxProviderSecInfo**](IpxProviderSecInfo.md) |  | [optional] 
**Sender** | Pointer to **string** | Fully Qualified Domain Name | [optional] 

## Methods

### NewSecParamExchRspData

`func NewSecParamExchRspData(n32fContextId string, ) *SecParamExchRspData`

NewSecParamExchRspData instantiates a new SecParamExchRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecParamExchRspDataWithDefaults

`func NewSecParamExchRspDataWithDefaults() *SecParamExchRspData`

NewSecParamExchRspDataWithDefaults instantiates a new SecParamExchRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN32fContextId

`func (o *SecParamExchRspData) GetN32fContextId() string`

GetN32fContextId returns the N32fContextId field if non-nil, zero value otherwise.

### GetN32fContextIdOk

`func (o *SecParamExchRspData) GetN32fContextIdOk() (*string, bool)`

GetN32fContextIdOk returns a tuple with the N32fContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN32fContextId

`func (o *SecParamExchRspData) SetN32fContextId(v string)`

SetN32fContextId sets N32fContextId field to given value.


### GetSelectedJweCipherSuite

`func (o *SecParamExchRspData) GetSelectedJweCipherSuite() string`

GetSelectedJweCipherSuite returns the SelectedJweCipherSuite field if non-nil, zero value otherwise.

### GetSelectedJweCipherSuiteOk

`func (o *SecParamExchRspData) GetSelectedJweCipherSuiteOk() (*string, bool)`

GetSelectedJweCipherSuiteOk returns a tuple with the SelectedJweCipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedJweCipherSuite

`func (o *SecParamExchRspData) SetSelectedJweCipherSuite(v string)`

SetSelectedJweCipherSuite sets SelectedJweCipherSuite field to given value.

### HasSelectedJweCipherSuite

`func (o *SecParamExchRspData) HasSelectedJweCipherSuite() bool`

HasSelectedJweCipherSuite returns a boolean if a field has been set.

### GetSelectedJwsCipherSuite

`func (o *SecParamExchRspData) GetSelectedJwsCipherSuite() string`

GetSelectedJwsCipherSuite returns the SelectedJwsCipherSuite field if non-nil, zero value otherwise.

### GetSelectedJwsCipherSuiteOk

`func (o *SecParamExchRspData) GetSelectedJwsCipherSuiteOk() (*string, bool)`

GetSelectedJwsCipherSuiteOk returns a tuple with the SelectedJwsCipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedJwsCipherSuite

`func (o *SecParamExchRspData) SetSelectedJwsCipherSuite(v string)`

SetSelectedJwsCipherSuite sets SelectedJwsCipherSuite field to given value.

### HasSelectedJwsCipherSuite

`func (o *SecParamExchRspData) HasSelectedJwsCipherSuite() bool`

HasSelectedJwsCipherSuite returns a boolean if a field has been set.

### GetSelProtectionPolicyInfo

`func (o *SecParamExchRspData) GetSelProtectionPolicyInfo() ProtectionPolicy`

GetSelProtectionPolicyInfo returns the SelProtectionPolicyInfo field if non-nil, zero value otherwise.

### GetSelProtectionPolicyInfoOk

`func (o *SecParamExchRspData) GetSelProtectionPolicyInfoOk() (*ProtectionPolicy, bool)`

GetSelProtectionPolicyInfoOk returns a tuple with the SelProtectionPolicyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelProtectionPolicyInfo

`func (o *SecParamExchRspData) SetSelProtectionPolicyInfo(v ProtectionPolicy)`

SetSelProtectionPolicyInfo sets SelProtectionPolicyInfo field to given value.

### HasSelProtectionPolicyInfo

`func (o *SecParamExchRspData) HasSelProtectionPolicyInfo() bool`

HasSelProtectionPolicyInfo returns a boolean if a field has been set.

### GetIpxProviderSecInfoList

`func (o *SecParamExchRspData) GetIpxProviderSecInfoList() []IpxProviderSecInfo`

GetIpxProviderSecInfoList returns the IpxProviderSecInfoList field if non-nil, zero value otherwise.

### GetIpxProviderSecInfoListOk

`func (o *SecParamExchRspData) GetIpxProviderSecInfoListOk() (*[]IpxProviderSecInfo, bool)`

GetIpxProviderSecInfoListOk returns a tuple with the IpxProviderSecInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpxProviderSecInfoList

`func (o *SecParamExchRspData) SetIpxProviderSecInfoList(v []IpxProviderSecInfo)`

SetIpxProviderSecInfoList sets IpxProviderSecInfoList field to given value.

### HasIpxProviderSecInfoList

`func (o *SecParamExchRspData) HasIpxProviderSecInfoList() bool`

HasIpxProviderSecInfoList returns a boolean if a field has been set.

### GetSender

`func (o *SecParamExchRspData) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *SecParamExchRspData) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *SecParamExchRspData) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *SecParamExchRspData) HasSender() bool`

HasSender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


