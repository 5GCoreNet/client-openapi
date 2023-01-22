# ComplexQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CnfUnits** | [**[]CnfUnit**](CnfUnit.md) |  | 
**DnfUnits** | [**[]DnfUnit**](DnfUnit.md) |  | 

## Methods

### NewComplexQuery

`func NewComplexQuery(cnfUnits []CnfUnit, dnfUnits []DnfUnit, ) *ComplexQuery`

NewComplexQuery instantiates a new ComplexQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplexQueryWithDefaults

`func NewComplexQueryWithDefaults() *ComplexQuery`

NewComplexQueryWithDefaults instantiates a new ComplexQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCnfUnits

`func (o *ComplexQuery) GetCnfUnits() []CnfUnit`

GetCnfUnits returns the CnfUnits field if non-nil, zero value otherwise.

### GetCnfUnitsOk

`func (o *ComplexQuery) GetCnfUnitsOk() (*[]CnfUnit, bool)`

GetCnfUnitsOk returns a tuple with the CnfUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnfUnits

`func (o *ComplexQuery) SetCnfUnits(v []CnfUnit)`

SetCnfUnits sets CnfUnits field to given value.


### GetDnfUnits

`func (o *ComplexQuery) GetDnfUnits() []DnfUnit`

GetDnfUnits returns the DnfUnits field if non-nil, zero value otherwise.

### GetDnfUnitsOk

`func (o *ComplexQuery) GetDnfUnitsOk() (*[]DnfUnit, bool)`

GetDnfUnitsOk returns a tuple with the DnfUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnfUnits

`func (o *ComplexQuery) SetDnfUnits(v []DnfUnit)`

SetDnfUnits sets DnfUnits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


