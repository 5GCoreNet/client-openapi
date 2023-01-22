# TopologicalServiceArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellIdList** | Pointer to **[]int32** |  | [optional] 
**TrackingAreaIdList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**ServingPLMN** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 

## Methods

### NewTopologicalServiceArea

`func NewTopologicalServiceArea() *TopologicalServiceArea`

NewTopologicalServiceArea instantiates a new TopologicalServiceArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologicalServiceAreaWithDefaults

`func NewTopologicalServiceAreaWithDefaults() *TopologicalServiceArea`

NewTopologicalServiceAreaWithDefaults instantiates a new TopologicalServiceArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellIdList

`func (o *TopologicalServiceArea) GetCellIdList() []int32`

GetCellIdList returns the CellIdList field if non-nil, zero value otherwise.

### GetCellIdListOk

`func (o *TopologicalServiceArea) GetCellIdListOk() (*[]int32, bool)`

GetCellIdListOk returns a tuple with the CellIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellIdList

`func (o *TopologicalServiceArea) SetCellIdList(v []int32)`

SetCellIdList sets CellIdList field to given value.

### HasCellIdList

`func (o *TopologicalServiceArea) HasCellIdList() bool`

HasCellIdList returns a boolean if a field has been set.

### GetTrackingAreaIdList

`func (o *TopologicalServiceArea) GetTrackingAreaIdList() []Tai`

GetTrackingAreaIdList returns the TrackingAreaIdList field if non-nil, zero value otherwise.

### GetTrackingAreaIdListOk

`func (o *TopologicalServiceArea) GetTrackingAreaIdListOk() (*[]Tai, bool)`

GetTrackingAreaIdListOk returns a tuple with the TrackingAreaIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingAreaIdList

`func (o *TopologicalServiceArea) SetTrackingAreaIdList(v []Tai)`

SetTrackingAreaIdList sets TrackingAreaIdList field to given value.

### HasTrackingAreaIdList

`func (o *TopologicalServiceArea) HasTrackingAreaIdList() bool`

HasTrackingAreaIdList returns a boolean if a field has been set.

### GetServingPLMN

`func (o *TopologicalServiceArea) GetServingPLMN() PlmnId`

GetServingPLMN returns the ServingPLMN field if non-nil, zero value otherwise.

### GetServingPLMNOk

`func (o *TopologicalServiceArea) GetServingPLMNOk() (*PlmnId, bool)`

GetServingPLMNOk returns a tuple with the ServingPLMN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingPLMN

`func (o *TopologicalServiceArea) SetServingPLMN(v PlmnId)`

SetServingPLMN sets ServingPLMN field to given value.

### HasServingPLMN

`func (o *TopologicalServiceArea) HasServingPLMN() bool`

HasServingPLMN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


