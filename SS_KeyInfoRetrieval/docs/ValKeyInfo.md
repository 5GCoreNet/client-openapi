# ValKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserUri** | **string** | string providing an URI formatted according to IETF RFC 3986. | 
**SkmsId** | Pointer to **string** | String identifying the key management server. | [optional] 
**ValService** | **string** | Unique identifier of a VAL Service. | 
**ValTgtUe** | Pointer to [**ValTargetUe**](ValTargetUe.md) |  | [optional] 
**KeyInfo** | **string** | Key management information specific to VAL service, VAL User or VAL UE. | 

## Methods

### NewValKeyInfo

`func NewValKeyInfo(userUri string, valService string, keyInfo string, ) *ValKeyInfo`

NewValKeyInfo instantiates a new ValKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValKeyInfoWithDefaults

`func NewValKeyInfoWithDefaults() *ValKeyInfo`

NewValKeyInfoWithDefaults instantiates a new ValKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserUri

`func (o *ValKeyInfo) GetUserUri() string`

GetUserUri returns the UserUri field if non-nil, zero value otherwise.

### GetUserUriOk

`func (o *ValKeyInfo) GetUserUriOk() (*string, bool)`

GetUserUriOk returns a tuple with the UserUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUri

`func (o *ValKeyInfo) SetUserUri(v string)`

SetUserUri sets UserUri field to given value.


### GetSkmsId

`func (o *ValKeyInfo) GetSkmsId() string`

GetSkmsId returns the SkmsId field if non-nil, zero value otherwise.

### GetSkmsIdOk

`func (o *ValKeyInfo) GetSkmsIdOk() (*string, bool)`

GetSkmsIdOk returns a tuple with the SkmsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkmsId

`func (o *ValKeyInfo) SetSkmsId(v string)`

SetSkmsId sets SkmsId field to given value.

### HasSkmsId

`func (o *ValKeyInfo) HasSkmsId() bool`

HasSkmsId returns a boolean if a field has been set.

### GetValService

`func (o *ValKeyInfo) GetValService() string`

GetValService returns the ValService field if non-nil, zero value otherwise.

### GetValServiceOk

`func (o *ValKeyInfo) GetValServiceOk() (*string, bool)`

GetValServiceOk returns a tuple with the ValService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValService

`func (o *ValKeyInfo) SetValService(v string)`

SetValService sets ValService field to given value.


### GetValTgtUe

`func (o *ValKeyInfo) GetValTgtUe() ValTargetUe`

GetValTgtUe returns the ValTgtUe field if non-nil, zero value otherwise.

### GetValTgtUeOk

`func (o *ValKeyInfo) GetValTgtUeOk() (*ValTargetUe, bool)`

GetValTgtUeOk returns a tuple with the ValTgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValTgtUe

`func (o *ValKeyInfo) SetValTgtUe(v ValTargetUe)`

SetValTgtUe sets ValTgtUe field to given value.

### HasValTgtUe

`func (o *ValKeyInfo) HasValTgtUe() bool`

HasValTgtUe returns a boolean if a field has been set.

### GetKeyInfo

`func (o *ValKeyInfo) GetKeyInfo() string`

GetKeyInfo returns the KeyInfo field if non-nil, zero value otherwise.

### GetKeyInfoOk

`func (o *ValKeyInfo) GetKeyInfoOk() (*string, bool)`

GetKeyInfoOk returns a tuple with the KeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInfo

`func (o *ValKeyInfo) SetKeyInfo(v string)`

SetKeyInfo sets KeyInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


