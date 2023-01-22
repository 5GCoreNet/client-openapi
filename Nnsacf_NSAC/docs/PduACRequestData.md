# PduACRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduACRequestInfo** | [**[]PduACRequestInfo**](PduACRequestInfo.md) |  | 
**NfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**PgwFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 

## Methods

### NewPduACRequestData

`func NewPduACRequestData(pduACRequestInfo []PduACRequestInfo, ) *PduACRequestData`

NewPduACRequestData instantiates a new PduACRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduACRequestDataWithDefaults

`func NewPduACRequestDataWithDefaults() *PduACRequestData`

NewPduACRequestDataWithDefaults instantiates a new PduACRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduACRequestInfo

`func (o *PduACRequestData) GetPduACRequestInfo() []PduACRequestInfo`

GetPduACRequestInfo returns the PduACRequestInfo field if non-nil, zero value otherwise.

### GetPduACRequestInfoOk

`func (o *PduACRequestData) GetPduACRequestInfoOk() (*[]PduACRequestInfo, bool)`

GetPduACRequestInfoOk returns a tuple with the PduACRequestInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduACRequestInfo

`func (o *PduACRequestData) SetPduACRequestInfo(v []PduACRequestInfo)`

SetPduACRequestInfo sets PduACRequestInfo field to given value.


### GetNfId

`func (o *PduACRequestData) GetNfId() string`

GetNfId returns the NfId field if non-nil, zero value otherwise.

### GetNfIdOk

`func (o *PduACRequestData) GetNfIdOk() (*string, bool)`

GetNfIdOk returns a tuple with the NfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfId

`func (o *PduACRequestData) SetNfId(v string)`

SetNfId sets NfId field to given value.

### HasNfId

`func (o *PduACRequestData) HasNfId() bool`

HasNfId returns a boolean if a field has been set.

### GetPgwFqdn

`func (o *PduACRequestData) GetPgwFqdn() string`

GetPgwFqdn returns the PgwFqdn field if non-nil, zero value otherwise.

### GetPgwFqdnOk

`func (o *PduACRequestData) GetPgwFqdnOk() (*string, bool)`

GetPgwFqdnOk returns a tuple with the PgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdn

`func (o *PduACRequestData) SetPgwFqdn(v string)`

SetPgwFqdn sets PgwFqdn field to given value.

### HasPgwFqdn

`func (o *PduACRequestData) HasPgwFqdn() bool`

HasPgwFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


