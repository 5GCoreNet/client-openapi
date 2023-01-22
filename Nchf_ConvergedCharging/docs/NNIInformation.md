# NNIInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionDirection** | Pointer to [**NNISessionDirection**](NNISessionDirection.md) |  | [optional] 
**NNIType** | Pointer to [**NNIType**](NNIType.md) |  | [optional] 
**RelationshipMode** | Pointer to [**NNIRelationshipMode**](NNIRelationshipMode.md) |  | [optional] 
**NeighbourNodeAddress** | Pointer to [**IMSAddress**](IMSAddress.md) |  | [optional] 

## Methods

### NewNNIInformation

`func NewNNIInformation() *NNIInformation`

NewNNIInformation instantiates a new NNIInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNNIInformationWithDefaults

`func NewNNIInformationWithDefaults() *NNIInformation`

NewNNIInformationWithDefaults instantiates a new NNIInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionDirection

`func (o *NNIInformation) GetSessionDirection() NNISessionDirection`

GetSessionDirection returns the SessionDirection field if non-nil, zero value otherwise.

### GetSessionDirectionOk

`func (o *NNIInformation) GetSessionDirectionOk() (*NNISessionDirection, bool)`

GetSessionDirectionOk returns a tuple with the SessionDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDirection

`func (o *NNIInformation) SetSessionDirection(v NNISessionDirection)`

SetSessionDirection sets SessionDirection field to given value.

### HasSessionDirection

`func (o *NNIInformation) HasSessionDirection() bool`

HasSessionDirection returns a boolean if a field has been set.

### GetNNIType

`func (o *NNIInformation) GetNNIType() NNIType`

GetNNIType returns the NNIType field if non-nil, zero value otherwise.

### GetNNITypeOk

`func (o *NNIInformation) GetNNITypeOk() (*NNIType, bool)`

GetNNITypeOk returns a tuple with the NNIType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNNIType

`func (o *NNIInformation) SetNNIType(v NNIType)`

SetNNIType sets NNIType field to given value.

### HasNNIType

`func (o *NNIInformation) HasNNIType() bool`

HasNNIType returns a boolean if a field has been set.

### GetRelationshipMode

`func (o *NNIInformation) GetRelationshipMode() NNIRelationshipMode`

GetRelationshipMode returns the RelationshipMode field if non-nil, zero value otherwise.

### GetRelationshipModeOk

`func (o *NNIInformation) GetRelationshipModeOk() (*NNIRelationshipMode, bool)`

GetRelationshipModeOk returns a tuple with the RelationshipMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipMode

`func (o *NNIInformation) SetRelationshipMode(v NNIRelationshipMode)`

SetRelationshipMode sets RelationshipMode field to given value.

### HasRelationshipMode

`func (o *NNIInformation) HasRelationshipMode() bool`

HasRelationshipMode returns a boolean if a field has been set.

### GetNeighbourNodeAddress

`func (o *NNIInformation) GetNeighbourNodeAddress() IMSAddress`

GetNeighbourNodeAddress returns the NeighbourNodeAddress field if non-nil, zero value otherwise.

### GetNeighbourNodeAddressOk

`func (o *NNIInformation) GetNeighbourNodeAddressOk() (*IMSAddress, bool)`

GetNeighbourNodeAddressOk returns a tuple with the NeighbourNodeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbourNodeAddress

`func (o *NNIInformation) SetNeighbourNodeAddress(v IMSAddress)`

SetNeighbourNodeAddress sets NeighbourNodeAddress field to given value.

### HasNeighbourNodeAddress

`func (o *NNIInformation) HasNeighbourNodeAddress() bool`

HasNeighbourNodeAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


