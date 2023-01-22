# MulticastTransportAddressChangeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LlSsm** | [**Ssm**](Ssm.md) |  | 
**CTeid** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | 
**AreaSessionId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 16-bit integer. | [optional] 

## Methods

### NewMulticastTransportAddressChangeInfo

`func NewMulticastTransportAddressChangeInfo(llSsm Ssm, cTeid int32, ) *MulticastTransportAddressChangeInfo`

NewMulticastTransportAddressChangeInfo instantiates a new MulticastTransportAddressChangeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulticastTransportAddressChangeInfoWithDefaults

`func NewMulticastTransportAddressChangeInfoWithDefaults() *MulticastTransportAddressChangeInfo`

NewMulticastTransportAddressChangeInfoWithDefaults instantiates a new MulticastTransportAddressChangeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLlSsm

`func (o *MulticastTransportAddressChangeInfo) GetLlSsm() Ssm`

GetLlSsm returns the LlSsm field if non-nil, zero value otherwise.

### GetLlSsmOk

`func (o *MulticastTransportAddressChangeInfo) GetLlSsmOk() (*Ssm, bool)`

GetLlSsmOk returns a tuple with the LlSsm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlSsm

`func (o *MulticastTransportAddressChangeInfo) SetLlSsm(v Ssm)`

SetLlSsm sets LlSsm field to given value.


### GetCTeid

`func (o *MulticastTransportAddressChangeInfo) GetCTeid() int32`

GetCTeid returns the CTeid field if non-nil, zero value otherwise.

### GetCTeidOk

`func (o *MulticastTransportAddressChangeInfo) GetCTeidOk() (*int32, bool)`

GetCTeidOk returns a tuple with the CTeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTeid

`func (o *MulticastTransportAddressChangeInfo) SetCTeid(v int32)`

SetCTeid sets CTeid field to given value.


### GetAreaSessionId

`func (o *MulticastTransportAddressChangeInfo) GetAreaSessionId() int32`

GetAreaSessionId returns the AreaSessionId field if non-nil, zero value otherwise.

### GetAreaSessionIdOk

`func (o *MulticastTransportAddressChangeInfo) GetAreaSessionIdOk() (*int32, bool)`

GetAreaSessionIdOk returns a tuple with the AreaSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaSessionId

`func (o *MulticastTransportAddressChangeInfo) SetAreaSessionId(v int32)`

SetAreaSessionId sets AreaSessionId field to given value.

### HasAreaSessionId

`func (o *MulticastTransportAddressChangeInfo) HasAreaSessionId() bool`

HasAreaSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


