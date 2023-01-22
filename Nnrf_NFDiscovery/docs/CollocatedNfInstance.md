# CollocatedNfInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**NfType** | [**CollocatedNfType**](CollocatedNfType.md) |  | 

## Methods

### NewCollocatedNfInstance

`func NewCollocatedNfInstance(nfInstanceId string, nfType CollocatedNfType, ) *CollocatedNfInstance`

NewCollocatedNfInstance instantiates a new CollocatedNfInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollocatedNfInstanceWithDefaults

`func NewCollocatedNfInstanceWithDefaults() *CollocatedNfInstance`

NewCollocatedNfInstanceWithDefaults instantiates a new CollocatedNfInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *CollocatedNfInstance) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *CollocatedNfInstance) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *CollocatedNfInstance) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetNfType

`func (o *CollocatedNfInstance) GetNfType() CollocatedNfType`

GetNfType returns the NfType field if non-nil, zero value otherwise.

### GetNfTypeOk

`func (o *CollocatedNfInstance) GetNfTypeOk() (*CollocatedNfType, bool)`

GetNfTypeOk returns a tuple with the NfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfType

`func (o *CollocatedNfInstance) SetNfType(v CollocatedNfType)`

SetNfType sets NfType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


