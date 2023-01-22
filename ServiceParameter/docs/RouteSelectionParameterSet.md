# RouteSelectionParameterSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Precedence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**SpatialValidityAreas** | Pointer to [**[]GeographicalArea**](GeographicalArea.md) | Indicates where the route selection parameters apply. It may correspond to a geographical area, for example using a geographic shape that is known to the AF and is configured by the operator to correspond to a list of or TAIs.  | [optional] 
**SpatialValidityTais** | Pointer to [**[]Tai**](Tai.md) | Indicates the TAIs in which the route selection parameters apply. This attribute is applicable only within the 5GC and it shall not be included in the request messages of untrusted AFs for URSP guidance. | [optional] 

## Methods

### NewRouteSelectionParameterSet

`func NewRouteSelectionParameterSet() *RouteSelectionParameterSet`

NewRouteSelectionParameterSet instantiates a new RouteSelectionParameterSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteSelectionParameterSetWithDefaults

`func NewRouteSelectionParameterSetWithDefaults() *RouteSelectionParameterSet`

NewRouteSelectionParameterSetWithDefaults instantiates a new RouteSelectionParameterSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *RouteSelectionParameterSet) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *RouteSelectionParameterSet) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *RouteSelectionParameterSet) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *RouteSelectionParameterSet) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *RouteSelectionParameterSet) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *RouteSelectionParameterSet) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *RouteSelectionParameterSet) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *RouteSelectionParameterSet) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetPrecedence

`func (o *RouteSelectionParameterSet) GetPrecedence() int32`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *RouteSelectionParameterSet) GetPrecedenceOk() (*int32, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *RouteSelectionParameterSet) SetPrecedence(v int32)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *RouteSelectionParameterSet) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetSpatialValidityAreas

`func (o *RouteSelectionParameterSet) GetSpatialValidityAreas() []GeographicalArea`

GetSpatialValidityAreas returns the SpatialValidityAreas field if non-nil, zero value otherwise.

### GetSpatialValidityAreasOk

`func (o *RouteSelectionParameterSet) GetSpatialValidityAreasOk() (*[]GeographicalArea, bool)`

GetSpatialValidityAreasOk returns a tuple with the SpatialValidityAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpatialValidityAreas

`func (o *RouteSelectionParameterSet) SetSpatialValidityAreas(v []GeographicalArea)`

SetSpatialValidityAreas sets SpatialValidityAreas field to given value.

### HasSpatialValidityAreas

`func (o *RouteSelectionParameterSet) HasSpatialValidityAreas() bool`

HasSpatialValidityAreas returns a boolean if a field has been set.

### GetSpatialValidityTais

`func (o *RouteSelectionParameterSet) GetSpatialValidityTais() []Tai`

GetSpatialValidityTais returns the SpatialValidityTais field if non-nil, zero value otherwise.

### GetSpatialValidityTaisOk

`func (o *RouteSelectionParameterSet) GetSpatialValidityTaisOk() (*[]Tai, bool)`

GetSpatialValidityTaisOk returns a tuple with the SpatialValidityTais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpatialValidityTais

`func (o *RouteSelectionParameterSet) SetSpatialValidityTais(v []Tai)`

SetSpatialValidityTais sets SpatialValidityTais field to given value.

### HasSpatialValidityTais

`func (o *RouteSelectionParameterSet) HasSpatialValidityTais() bool`

HasSpatialValidityTais returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


