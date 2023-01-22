# N4Information

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N4MessageType** | [**N4MessageType**](N4MessageType.md) |  | 
**N4MessagePayload** | [**RefToBinaryData**](RefToBinaryData.md) |  | 
**N4DnaiInfo** | Pointer to [**DnaiInformation**](DnaiInformation.md) |  | [optional] 
**PsaUpfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**UlClBpId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**N9UlPdrIdList** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewN4Information

`func NewN4Information(n4MessageType N4MessageType, n4MessagePayload RefToBinaryData, ) *N4Information`

NewN4Information instantiates a new N4Information object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN4InformationWithDefaults

`func NewN4InformationWithDefaults() *N4Information`

NewN4InformationWithDefaults instantiates a new N4Information object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN4MessageType

`func (o *N4Information) GetN4MessageType() N4MessageType`

GetN4MessageType returns the N4MessageType field if non-nil, zero value otherwise.

### GetN4MessageTypeOk

`func (o *N4Information) GetN4MessageTypeOk() (*N4MessageType, bool)`

GetN4MessageTypeOk returns a tuple with the N4MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4MessageType

`func (o *N4Information) SetN4MessageType(v N4MessageType)`

SetN4MessageType sets N4MessageType field to given value.


### GetN4MessagePayload

`func (o *N4Information) GetN4MessagePayload() RefToBinaryData`

GetN4MessagePayload returns the N4MessagePayload field if non-nil, zero value otherwise.

### GetN4MessagePayloadOk

`func (o *N4Information) GetN4MessagePayloadOk() (*RefToBinaryData, bool)`

GetN4MessagePayloadOk returns a tuple with the N4MessagePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4MessagePayload

`func (o *N4Information) SetN4MessagePayload(v RefToBinaryData)`

SetN4MessagePayload sets N4MessagePayload field to given value.


### GetN4DnaiInfo

`func (o *N4Information) GetN4DnaiInfo() DnaiInformation`

GetN4DnaiInfo returns the N4DnaiInfo field if non-nil, zero value otherwise.

### GetN4DnaiInfoOk

`func (o *N4Information) GetN4DnaiInfoOk() (*DnaiInformation, bool)`

GetN4DnaiInfoOk returns a tuple with the N4DnaiInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4DnaiInfo

`func (o *N4Information) SetN4DnaiInfo(v DnaiInformation)`

SetN4DnaiInfo sets N4DnaiInfo field to given value.

### HasN4DnaiInfo

`func (o *N4Information) HasN4DnaiInfo() bool`

HasN4DnaiInfo returns a boolean if a field has been set.

### GetPsaUpfId

`func (o *N4Information) GetPsaUpfId() string`

GetPsaUpfId returns the PsaUpfId field if non-nil, zero value otherwise.

### GetPsaUpfIdOk

`func (o *N4Information) GetPsaUpfIdOk() (*string, bool)`

GetPsaUpfIdOk returns a tuple with the PsaUpfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsaUpfId

`func (o *N4Information) SetPsaUpfId(v string)`

SetPsaUpfId sets PsaUpfId field to given value.

### HasPsaUpfId

`func (o *N4Information) HasPsaUpfId() bool`

HasPsaUpfId returns a boolean if a field has been set.

### GetUlClBpId

`func (o *N4Information) GetUlClBpId() string`

GetUlClBpId returns the UlClBpId field if non-nil, zero value otherwise.

### GetUlClBpIdOk

`func (o *N4Information) GetUlClBpIdOk() (*string, bool)`

GetUlClBpIdOk returns a tuple with the UlClBpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlClBpId

`func (o *N4Information) SetUlClBpId(v string)`

SetUlClBpId sets UlClBpId field to given value.

### HasUlClBpId

`func (o *N4Information) HasUlClBpId() bool`

HasUlClBpId returns a boolean if a field has been set.

### GetN9UlPdrIdList

`func (o *N4Information) GetN9UlPdrIdList() []int32`

GetN9UlPdrIdList returns the N9UlPdrIdList field if non-nil, zero value otherwise.

### GetN9UlPdrIdListOk

`func (o *N4Information) GetN9UlPdrIdListOk() (*[]int32, bool)`

GetN9UlPdrIdListOk returns a tuple with the N9UlPdrIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN9UlPdrIdList

`func (o *N4Information) SetN9UlPdrIdList(v []int32)`

SetN9UlPdrIdList sets N9UlPdrIdList field to given value.

### HasN9UlPdrIdList

`func (o *N4Information) HasN9UlPdrIdList() bool`

HasN9UlPdrIdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


