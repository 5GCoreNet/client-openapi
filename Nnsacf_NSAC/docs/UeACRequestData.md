# UeACRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeACRequestInfo** | [**[]UeACRequestInfo**](UeACRequestInfo.md) |  | 
**NfId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**NfType** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**EacNotificationUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

## Methods

### NewUeACRequestData

`func NewUeACRequestData(ueACRequestInfo []UeACRequestInfo, nfId string, ) *UeACRequestData`

NewUeACRequestData instantiates a new UeACRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeACRequestDataWithDefaults

`func NewUeACRequestDataWithDefaults() *UeACRequestData`

NewUeACRequestDataWithDefaults instantiates a new UeACRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeACRequestInfo

`func (o *UeACRequestData) GetUeACRequestInfo() []UeACRequestInfo`

GetUeACRequestInfo returns the UeACRequestInfo field if non-nil, zero value otherwise.

### GetUeACRequestInfoOk

`func (o *UeACRequestData) GetUeACRequestInfoOk() (*[]UeACRequestInfo, bool)`

GetUeACRequestInfoOk returns a tuple with the UeACRequestInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeACRequestInfo

`func (o *UeACRequestData) SetUeACRequestInfo(v []UeACRequestInfo)`

SetUeACRequestInfo sets UeACRequestInfo field to given value.


### GetNfId

`func (o *UeACRequestData) GetNfId() string`

GetNfId returns the NfId field if non-nil, zero value otherwise.

### GetNfIdOk

`func (o *UeACRequestData) GetNfIdOk() (*string, bool)`

GetNfIdOk returns a tuple with the NfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfId

`func (o *UeACRequestData) SetNfId(v string)`

SetNfId sets NfId field to given value.


### GetNfType

`func (o *UeACRequestData) GetNfType() NFType`

GetNfType returns the NfType field if non-nil, zero value otherwise.

### GetNfTypeOk

`func (o *UeACRequestData) GetNfTypeOk() (*NFType, bool)`

GetNfTypeOk returns a tuple with the NfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfType

`func (o *UeACRequestData) SetNfType(v NFType)`

SetNfType sets NfType field to given value.

### HasNfType

`func (o *UeACRequestData) HasNfType() bool`

HasNfType returns a boolean if a field has been set.

### GetEacNotificationUri

`func (o *UeACRequestData) GetEacNotificationUri() string`

GetEacNotificationUri returns the EacNotificationUri field if non-nil, zero value otherwise.

### GetEacNotificationUriOk

`func (o *UeACRequestData) GetEacNotificationUriOk() (*string, bool)`

GetEacNotificationUriOk returns a tuple with the EacNotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEacNotificationUri

`func (o *UeACRequestData) SetEacNotificationUri(v string)`

SetEacNotificationUri sets EacNotificationUri field to given value.

### HasEacNotificationUri

`func (o *UeACRequestData) HasEacNotificationUri() bool`

HasEacNotificationUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


