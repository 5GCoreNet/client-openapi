# NfIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfType** | [**NFType**](NFType.md) |  | 
**NfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 

## Methods

### NewNfIdentifier

`func NewNfIdentifier(nfType NFType, ) *NfIdentifier`

NewNfIdentifier instantiates a new NfIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfIdentifierWithDefaults

`func NewNfIdentifierWithDefaults() *NfIdentifier`

NewNfIdentifierWithDefaults instantiates a new NfIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfType

`func (o *NfIdentifier) GetNfType() NFType`

GetNfType returns the NfType field if non-nil, zero value otherwise.

### GetNfTypeOk

`func (o *NfIdentifier) GetNfTypeOk() (*NFType, bool)`

GetNfTypeOk returns a tuple with the NfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfType

`func (o *NfIdentifier) SetNfType(v NFType)`

SetNfType sets NfType field to given value.


### GetNfInstanceId

`func (o *NfIdentifier) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *NfIdentifier) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *NfIdentifier) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.

### HasNfInstanceId

`func (o *NfIdentifier) HasNfInstanceId() bool`

HasNfInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


