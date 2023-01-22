# PcfUeCallbackInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**BindingInfo** | Pointer to **string** |  | [optional] 

## Methods

### NewPcfUeCallbackInfo

`func NewPcfUeCallbackInfo(callbackUri string, ) *PcfUeCallbackInfo`

NewPcfUeCallbackInfo instantiates a new PcfUeCallbackInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcfUeCallbackInfoWithDefaults

`func NewPcfUeCallbackInfoWithDefaults() *PcfUeCallbackInfo`

NewPcfUeCallbackInfoWithDefaults instantiates a new PcfUeCallbackInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUri

`func (o *PcfUeCallbackInfo) GetCallbackUri() string`

GetCallbackUri returns the CallbackUri field if non-nil, zero value otherwise.

### GetCallbackUriOk

`func (o *PcfUeCallbackInfo) GetCallbackUriOk() (*string, bool)`

GetCallbackUriOk returns a tuple with the CallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUri

`func (o *PcfUeCallbackInfo) SetCallbackUri(v string)`

SetCallbackUri sets CallbackUri field to given value.


### GetBindingInfo

`func (o *PcfUeCallbackInfo) GetBindingInfo() string`

GetBindingInfo returns the BindingInfo field if non-nil, zero value otherwise.

### GetBindingInfoOk

`func (o *PcfUeCallbackInfo) GetBindingInfoOk() (*string, bool)`

GetBindingInfoOk returns a tuple with the BindingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingInfo

`func (o *PcfUeCallbackInfo) SetBindingInfo(v string)`

SetBindingInfo sets BindingInfo field to given value.

### HasBindingInfo

`func (o *PcfUeCallbackInfo) HasBindingInfo() bool`

HasBindingInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


