# NfInstanceIdCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 

## Methods

### NewNfInstanceIdCond

`func NewNfInstanceIdCond(nfInstanceId string, ) *NfInstanceIdCond`

NewNfInstanceIdCond instantiates a new NfInstanceIdCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfInstanceIdCondWithDefaults

`func NewNfInstanceIdCondWithDefaults() *NfInstanceIdCond`

NewNfInstanceIdCondWithDefaults instantiates a new NfInstanceIdCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *NfInstanceIdCond) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *NfInstanceIdCond) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *NfInstanceIdCond) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


