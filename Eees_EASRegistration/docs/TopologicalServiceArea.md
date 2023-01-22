# TopologicalServiceArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ecgis** | Pointer to [**[]Ecgi**](Ecgi.md) | A list of E-UTRA cell identities. | [optional] 
**Ncgis** | Pointer to [**[]Ncgi**](Ncgi.md) | A list of NR cell identities. | [optional] 
**Tais** | Pointer to [**[]Tai**](Tai.md) | A list of tracking area identities. | [optional] 
**PlmnIds** | Pointer to [**[]PlmnId1**](PlmnId1.md) | A list of PLMN identities. | [optional] 

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

### GetEcgis

`func (o *TopologicalServiceArea) GetEcgis() []Ecgi`

GetEcgis returns the Ecgis field if non-nil, zero value otherwise.

### GetEcgisOk

`func (o *TopologicalServiceArea) GetEcgisOk() (*[]Ecgi, bool)`

GetEcgisOk returns a tuple with the Ecgis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgis

`func (o *TopologicalServiceArea) SetEcgis(v []Ecgi)`

SetEcgis sets Ecgis field to given value.

### HasEcgis

`func (o *TopologicalServiceArea) HasEcgis() bool`

HasEcgis returns a boolean if a field has been set.

### GetNcgis

`func (o *TopologicalServiceArea) GetNcgis() []Ncgi`

GetNcgis returns the Ncgis field if non-nil, zero value otherwise.

### GetNcgisOk

`func (o *TopologicalServiceArea) GetNcgisOk() (*[]Ncgi, bool)`

GetNcgisOk returns a tuple with the Ncgis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgis

`func (o *TopologicalServiceArea) SetNcgis(v []Ncgi)`

SetNcgis sets Ncgis field to given value.

### HasNcgis

`func (o *TopologicalServiceArea) HasNcgis() bool`

HasNcgis returns a boolean if a field has been set.

### GetTais

`func (o *TopologicalServiceArea) GetTais() []Tai`

GetTais returns the Tais field if non-nil, zero value otherwise.

### GetTaisOk

`func (o *TopologicalServiceArea) GetTaisOk() (*[]Tai, bool)`

GetTaisOk returns a tuple with the Tais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTais

`func (o *TopologicalServiceArea) SetTais(v []Tai)`

SetTais sets Tais field to given value.

### HasTais

`func (o *TopologicalServiceArea) HasTais() bool`

HasTais returns a boolean if a field has been set.

### GetPlmnIds

`func (o *TopologicalServiceArea) GetPlmnIds() []PlmnId1`

GetPlmnIds returns the PlmnIds field if non-nil, zero value otherwise.

### GetPlmnIdsOk

`func (o *TopologicalServiceArea) GetPlmnIdsOk() (*[]PlmnId1, bool)`

GetPlmnIdsOk returns a tuple with the PlmnIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnIds

`func (o *TopologicalServiceArea) SetPlmnIds(v []PlmnId1)`

SetPlmnIds sets PlmnIds field to given value.

### HasPlmnIds

`func (o *TopologicalServiceArea) HasPlmnIds() bool`

HasPlmnIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


