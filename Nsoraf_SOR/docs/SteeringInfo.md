# SteeringInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**SnpnId** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**Gin** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**AccessTechList** | Pointer to [**[]AccessTech**](AccessTech.md) |  | [optional] 

## Methods

### NewSteeringInfo

`func NewSteeringInfo() *SteeringInfo`

NewSteeringInfo instantiates a new SteeringInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSteeringInfoWithDefaults

`func NewSteeringInfoWithDefaults() *SteeringInfo`

NewSteeringInfoWithDefaults instantiates a new SteeringInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *SteeringInfo) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *SteeringInfo) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *SteeringInfo) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *SteeringInfo) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetSnpnId

`func (o *SteeringInfo) GetSnpnId() PlmnIdNid`

GetSnpnId returns the SnpnId field if non-nil, zero value otherwise.

### GetSnpnIdOk

`func (o *SteeringInfo) GetSnpnIdOk() (*PlmnIdNid, bool)`

GetSnpnIdOk returns a tuple with the SnpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnpnId

`func (o *SteeringInfo) SetSnpnId(v PlmnIdNid)`

SetSnpnId sets SnpnId field to given value.

### HasSnpnId

`func (o *SteeringInfo) HasSnpnId() bool`

HasSnpnId returns a boolean if a field has been set.

### GetGin

`func (o *SteeringInfo) GetGin() PlmnIdNid`

GetGin returns the Gin field if non-nil, zero value otherwise.

### GetGinOk

`func (o *SteeringInfo) GetGinOk() (*PlmnIdNid, bool)`

GetGinOk returns a tuple with the Gin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGin

`func (o *SteeringInfo) SetGin(v PlmnIdNid)`

SetGin sets Gin field to given value.

### HasGin

`func (o *SteeringInfo) HasGin() bool`

HasGin returns a boolean if a field has been set.

### GetAccessTechList

`func (o *SteeringInfo) GetAccessTechList() []AccessTech`

GetAccessTechList returns the AccessTechList field if non-nil, zero value otherwise.

### GetAccessTechListOk

`func (o *SteeringInfo) GetAccessTechListOk() (*[]AccessTech, bool)`

GetAccessTechListOk returns a tuple with the AccessTechList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTechList

`func (o *SteeringInfo) SetAccessTechList(v []AccessTech)`

SetAccessTechList sets AccessTechList field to given value.

### HasAccessTechList

`func (o *SteeringInfo) HasAccessTechList() bool`

HasAccessTechList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


