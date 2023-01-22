# AccessControlPolicyListExt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiInvokerPolicies** | Pointer to [**[]ApiInvokerPolicy**](ApiInvokerPolicy.md) | Policy of each API invoker. | [optional] 
**ApiId** | **string** |  | 

## Methods

### NewAccessControlPolicyListExt

`func NewAccessControlPolicyListExt(apiId string, ) *AccessControlPolicyListExt`

NewAccessControlPolicyListExt instantiates a new AccessControlPolicyListExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlPolicyListExtWithDefaults

`func NewAccessControlPolicyListExtWithDefaults() *AccessControlPolicyListExt`

NewAccessControlPolicyListExtWithDefaults instantiates a new AccessControlPolicyListExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiInvokerPolicies

`func (o *AccessControlPolicyListExt) GetApiInvokerPolicies() []ApiInvokerPolicy`

GetApiInvokerPolicies returns the ApiInvokerPolicies field if non-nil, zero value otherwise.

### GetApiInvokerPoliciesOk

`func (o *AccessControlPolicyListExt) GetApiInvokerPoliciesOk() (*[]ApiInvokerPolicy, bool)`

GetApiInvokerPoliciesOk returns a tuple with the ApiInvokerPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInvokerPolicies

`func (o *AccessControlPolicyListExt) SetApiInvokerPolicies(v []ApiInvokerPolicy)`

SetApiInvokerPolicies sets ApiInvokerPolicies field to given value.

### HasApiInvokerPolicies

`func (o *AccessControlPolicyListExt) HasApiInvokerPolicies() bool`

HasApiInvokerPolicies returns a boolean if a field has been set.

### GetApiId

`func (o *AccessControlPolicyListExt) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *AccessControlPolicyListExt) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *AccessControlPolicyListExt) SetApiId(v string)`

SetApiId sets ApiId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


