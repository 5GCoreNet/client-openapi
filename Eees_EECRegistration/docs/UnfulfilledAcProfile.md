# UnfulfilledAcProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcId** | Pointer to **string** | The AC ID of a AC profile. | [optional] 
**Reason** | Pointer to [**UnfulfillACProfRsn**](UnfulfillACProfRsn.md) |  | [optional] 

## Methods

### NewUnfulfilledAcProfile

`func NewUnfulfilledAcProfile() *UnfulfilledAcProfile`

NewUnfulfilledAcProfile instantiates a new UnfulfilledAcProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnfulfilledAcProfileWithDefaults

`func NewUnfulfilledAcProfileWithDefaults() *UnfulfilledAcProfile`

NewUnfulfilledAcProfileWithDefaults instantiates a new UnfulfilledAcProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcId

`func (o *UnfulfilledAcProfile) GetAcId() string`

GetAcId returns the AcId field if non-nil, zero value otherwise.

### GetAcIdOk

`func (o *UnfulfilledAcProfile) GetAcIdOk() (*string, bool)`

GetAcIdOk returns a tuple with the AcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcId

`func (o *UnfulfilledAcProfile) SetAcId(v string)`

SetAcId sets AcId field to given value.

### HasAcId

`func (o *UnfulfilledAcProfile) HasAcId() bool`

HasAcId returns a boolean if a field has been set.

### GetReason

`func (o *UnfulfilledAcProfile) GetReason() UnfulfillACProfRsn`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UnfulfilledAcProfile) GetReasonOk() (*UnfulfillACProfRsn, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UnfulfilledAcProfile) SetReason(v UnfulfillACProfRsn)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UnfulfilledAcProfile) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


