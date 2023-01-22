# ProseKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var5gPrukId** | **string** | The 5GPRUK ID is string in NAI format as specified in clause 28.7.19 of 3GPP TS 23.003.  | 
**RelayServiceCode** | **int32** | Relay Service Code to identify a connectivity service provided by the UE-to-Network relay.  | 

## Methods

### NewProseKeyRequest

`func NewProseKeyRequest(var5gPrukId string, relayServiceCode int32, ) *ProseKeyRequest`

NewProseKeyRequest instantiates a new ProseKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProseKeyRequestWithDefaults

`func NewProseKeyRequestWithDefaults() *ProseKeyRequest`

NewProseKeyRequestWithDefaults instantiates a new ProseKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar5gPrukId

`func (o *ProseKeyRequest) GetVar5gPrukId() string`

GetVar5gPrukId returns the Var5gPrukId field if non-nil, zero value otherwise.

### GetVar5gPrukIdOk

`func (o *ProseKeyRequest) GetVar5gPrukIdOk() (*string, bool)`

GetVar5gPrukIdOk returns a tuple with the Var5gPrukId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gPrukId

`func (o *ProseKeyRequest) SetVar5gPrukId(v string)`

SetVar5gPrukId sets Var5gPrukId field to given value.


### GetRelayServiceCode

`func (o *ProseKeyRequest) GetRelayServiceCode() int32`

GetRelayServiceCode returns the RelayServiceCode field if non-nil, zero value otherwise.

### GetRelayServiceCodeOk

`func (o *ProseKeyRequest) GetRelayServiceCodeOk() (*int32, bool)`

GetRelayServiceCodeOk returns a tuple with the RelayServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayServiceCode

`func (o *ProseKeyRequest) SetRelayServiceCode(v int32)`

SetRelayServiceCode sets RelayServiceCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


