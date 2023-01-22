# UpdatePduSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**HsmfUpdateData**](HsmfUpdateData.md) |  | [optional] 
**BinaryDataN1SmInfoFromUe** | Pointer to ***os.File** |  | [optional] 
**BinaryDataUnknownN1SmInfo** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN4Information** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN4InformationExt1** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN4InformationExt2** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewUpdatePduSessionRequest

`func NewUpdatePduSessionRequest() *UpdatePduSessionRequest`

NewUpdatePduSessionRequest instantiates a new UpdatePduSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePduSessionRequestWithDefaults

`func NewUpdatePduSessionRequestWithDefaults() *UpdatePduSessionRequest`

NewUpdatePduSessionRequestWithDefaults instantiates a new UpdatePduSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *UpdatePduSessionRequest) GetJsonData() HsmfUpdateData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *UpdatePduSessionRequest) GetJsonDataOk() (*HsmfUpdateData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *UpdatePduSessionRequest) SetJsonData(v HsmfUpdateData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *UpdatePduSessionRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN1SmInfoFromUe

`func (o *UpdatePduSessionRequest) GetBinaryDataN1SmInfoFromUe() *os.File`

GetBinaryDataN1SmInfoFromUe returns the BinaryDataN1SmInfoFromUe field if non-nil, zero value otherwise.

### GetBinaryDataN1SmInfoFromUeOk

`func (o *UpdatePduSessionRequest) GetBinaryDataN1SmInfoFromUeOk() (**os.File, bool)`

GetBinaryDataN1SmInfoFromUeOk returns a tuple with the BinaryDataN1SmInfoFromUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN1SmInfoFromUe

`func (o *UpdatePduSessionRequest) SetBinaryDataN1SmInfoFromUe(v *os.File)`

SetBinaryDataN1SmInfoFromUe sets BinaryDataN1SmInfoFromUe field to given value.

### HasBinaryDataN1SmInfoFromUe

`func (o *UpdatePduSessionRequest) HasBinaryDataN1SmInfoFromUe() bool`

HasBinaryDataN1SmInfoFromUe returns a boolean if a field has been set.

### GetBinaryDataUnknownN1SmInfo

`func (o *UpdatePduSessionRequest) GetBinaryDataUnknownN1SmInfo() *os.File`

GetBinaryDataUnknownN1SmInfo returns the BinaryDataUnknownN1SmInfo field if non-nil, zero value otherwise.

### GetBinaryDataUnknownN1SmInfoOk

`func (o *UpdatePduSessionRequest) GetBinaryDataUnknownN1SmInfoOk() (**os.File, bool)`

GetBinaryDataUnknownN1SmInfoOk returns a tuple with the BinaryDataUnknownN1SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataUnknownN1SmInfo

`func (o *UpdatePduSessionRequest) SetBinaryDataUnknownN1SmInfo(v *os.File)`

SetBinaryDataUnknownN1SmInfo sets BinaryDataUnknownN1SmInfo field to given value.

### HasBinaryDataUnknownN1SmInfo

`func (o *UpdatePduSessionRequest) HasBinaryDataUnknownN1SmInfo() bool`

HasBinaryDataUnknownN1SmInfo returns a boolean if a field has been set.

### GetBinaryDataN4Information

`func (o *UpdatePduSessionRequest) GetBinaryDataN4Information() *os.File`

GetBinaryDataN4Information returns the BinaryDataN4Information field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationOk

`func (o *UpdatePduSessionRequest) GetBinaryDataN4InformationOk() (**os.File, bool)`

GetBinaryDataN4InformationOk returns a tuple with the BinaryDataN4Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4Information

`func (o *UpdatePduSessionRequest) SetBinaryDataN4Information(v *os.File)`

SetBinaryDataN4Information sets BinaryDataN4Information field to given value.

### HasBinaryDataN4Information

`func (o *UpdatePduSessionRequest) HasBinaryDataN4Information() bool`

HasBinaryDataN4Information returns a boolean if a field has been set.

### GetBinaryDataN4InformationExt1

`func (o *UpdatePduSessionRequest) GetBinaryDataN4InformationExt1() *os.File`

GetBinaryDataN4InformationExt1 returns the BinaryDataN4InformationExt1 field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationExt1Ok

`func (o *UpdatePduSessionRequest) GetBinaryDataN4InformationExt1Ok() (**os.File, bool)`

GetBinaryDataN4InformationExt1Ok returns a tuple with the BinaryDataN4InformationExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4InformationExt1

`func (o *UpdatePduSessionRequest) SetBinaryDataN4InformationExt1(v *os.File)`

SetBinaryDataN4InformationExt1 sets BinaryDataN4InformationExt1 field to given value.

### HasBinaryDataN4InformationExt1

`func (o *UpdatePduSessionRequest) HasBinaryDataN4InformationExt1() bool`

HasBinaryDataN4InformationExt1 returns a boolean if a field has been set.

### GetBinaryDataN4InformationExt2

`func (o *UpdatePduSessionRequest) GetBinaryDataN4InformationExt2() *os.File`

GetBinaryDataN4InformationExt2 returns the BinaryDataN4InformationExt2 field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationExt2Ok

`func (o *UpdatePduSessionRequest) GetBinaryDataN4InformationExt2Ok() (**os.File, bool)`

GetBinaryDataN4InformationExt2Ok returns a tuple with the BinaryDataN4InformationExt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4InformationExt2

`func (o *UpdatePduSessionRequest) SetBinaryDataN4InformationExt2(v *os.File)`

SetBinaryDataN4InformationExt2 sets BinaryDataN4InformationExt2 field to given value.

### HasBinaryDataN4InformationExt2

`func (o *UpdatePduSessionRequest) HasBinaryDataN4InformationExt2() bool`

HasBinaryDataN4InformationExt2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


