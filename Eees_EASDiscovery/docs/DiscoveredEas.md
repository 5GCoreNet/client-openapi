# DiscoveredEas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eas** | [**EASProfile**](EASProfile.md) |  | 
**LifeTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 

## Methods

### NewDiscoveredEas

`func NewDiscoveredEas(eas EASProfile, ) *DiscoveredEas`

NewDiscoveredEas instantiates a new DiscoveredEas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredEasWithDefaults

`func NewDiscoveredEasWithDefaults() *DiscoveredEas`

NewDiscoveredEasWithDefaults instantiates a new DiscoveredEas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEas

`func (o *DiscoveredEas) GetEas() EASProfile`

GetEas returns the Eas field if non-nil, zero value otherwise.

### GetEasOk

`func (o *DiscoveredEas) GetEasOk() (*EASProfile, bool)`

GetEasOk returns a tuple with the Eas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEas

`func (o *DiscoveredEas) SetEas(v EASProfile)`

SetEas sets Eas field to given value.


### GetLifeTime

`func (o *DiscoveredEas) GetLifeTime() time.Time`

GetLifeTime returns the LifeTime field if non-nil, zero value otherwise.

### GetLifeTimeOk

`func (o *DiscoveredEas) GetLifeTimeOk() (*time.Time, bool)`

GetLifeTimeOk returns a tuple with the LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeTime

`func (o *DiscoveredEas) SetLifeTime(v time.Time)`

SetLifeTime sets LifeTime field to given value.

### HasLifeTime

`func (o *DiscoveredEas) HasLifeTime() bool`

HasLifeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


