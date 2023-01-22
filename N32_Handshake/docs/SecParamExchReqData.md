# SecParamExchReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N32fContextId** | **string** |  | 
**JweCipherSuiteList** | Pointer to **[]string** |  | [optional] 
**JwsCipherSuiteList** | Pointer to **[]string** |  | [optional] 
**ProtectionPolicyInfo** | Pointer to [**ProtectionPolicy**](ProtectionPolicy.md) |  | [optional] 
**IpxProviderSecInfoList** | Pointer to [**[]IpxProviderSecInfo**](IpxProviderSecInfo.md) |  | [optional] 
**Sender** | Pointer to **string** | Fully Qualified Domain Name | [optional] 

## Methods

### NewSecParamExchReqData

`func NewSecParamExchReqData(n32fContextId string, ) *SecParamExchReqData`

NewSecParamExchReqData instantiates a new SecParamExchReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecParamExchReqDataWithDefaults

`func NewSecParamExchReqDataWithDefaults() *SecParamExchReqData`

NewSecParamExchReqDataWithDefaults instantiates a new SecParamExchReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN32fContextId

`func (o *SecParamExchReqData) GetN32fContextId() string`

GetN32fContextId returns the N32fContextId field if non-nil, zero value otherwise.

### GetN32fContextIdOk

`func (o *SecParamExchReqData) GetN32fContextIdOk() (*string, bool)`

GetN32fContextIdOk returns a tuple with the N32fContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN32fContextId

`func (o *SecParamExchReqData) SetN32fContextId(v string)`

SetN32fContextId sets N32fContextId field to given value.


### GetJweCipherSuiteList

`func (o *SecParamExchReqData) GetJweCipherSuiteList() []string`

GetJweCipherSuiteList returns the JweCipherSuiteList field if non-nil, zero value otherwise.

### GetJweCipherSuiteListOk

`func (o *SecParamExchReqData) GetJweCipherSuiteListOk() (*[]string, bool)`

GetJweCipherSuiteListOk returns a tuple with the JweCipherSuiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJweCipherSuiteList

`func (o *SecParamExchReqData) SetJweCipherSuiteList(v []string)`

SetJweCipherSuiteList sets JweCipherSuiteList field to given value.

### HasJweCipherSuiteList

`func (o *SecParamExchReqData) HasJweCipherSuiteList() bool`

HasJweCipherSuiteList returns a boolean if a field has been set.

### GetJwsCipherSuiteList

`func (o *SecParamExchReqData) GetJwsCipherSuiteList() []string`

GetJwsCipherSuiteList returns the JwsCipherSuiteList field if non-nil, zero value otherwise.

### GetJwsCipherSuiteListOk

`func (o *SecParamExchReqData) GetJwsCipherSuiteListOk() (*[]string, bool)`

GetJwsCipherSuiteListOk returns a tuple with the JwsCipherSuiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwsCipherSuiteList

`func (o *SecParamExchReqData) SetJwsCipherSuiteList(v []string)`

SetJwsCipherSuiteList sets JwsCipherSuiteList field to given value.

### HasJwsCipherSuiteList

`func (o *SecParamExchReqData) HasJwsCipherSuiteList() bool`

HasJwsCipherSuiteList returns a boolean if a field has been set.

### GetProtectionPolicyInfo

`func (o *SecParamExchReqData) GetProtectionPolicyInfo() ProtectionPolicy`

GetProtectionPolicyInfo returns the ProtectionPolicyInfo field if non-nil, zero value otherwise.

### GetProtectionPolicyInfoOk

`func (o *SecParamExchReqData) GetProtectionPolicyInfoOk() (*ProtectionPolicy, bool)`

GetProtectionPolicyInfoOk returns a tuple with the ProtectionPolicyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionPolicyInfo

`func (o *SecParamExchReqData) SetProtectionPolicyInfo(v ProtectionPolicy)`

SetProtectionPolicyInfo sets ProtectionPolicyInfo field to given value.

### HasProtectionPolicyInfo

`func (o *SecParamExchReqData) HasProtectionPolicyInfo() bool`

HasProtectionPolicyInfo returns a boolean if a field has been set.

### GetIpxProviderSecInfoList

`func (o *SecParamExchReqData) GetIpxProviderSecInfoList() []IpxProviderSecInfo`

GetIpxProviderSecInfoList returns the IpxProviderSecInfoList field if non-nil, zero value otherwise.

### GetIpxProviderSecInfoListOk

`func (o *SecParamExchReqData) GetIpxProviderSecInfoListOk() (*[]IpxProviderSecInfo, bool)`

GetIpxProviderSecInfoListOk returns a tuple with the IpxProviderSecInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpxProviderSecInfoList

`func (o *SecParamExchReqData) SetIpxProviderSecInfoList(v []IpxProviderSecInfo)`

SetIpxProviderSecInfoList sets IpxProviderSecInfoList field to given value.

### HasIpxProviderSecInfoList

`func (o *SecParamExchReqData) HasIpxProviderSecInfoList() bool`

HasIpxProviderSecInfoList returns a boolean if a field has been set.

### GetSender

`func (o *SecParamExchReqData) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *SecParamExchReqData) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *SecParamExchReqData) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *SecParamExchReqData) HasSender() bool`

HasSender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


