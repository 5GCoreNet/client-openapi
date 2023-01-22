# EDNConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdnConInfo** | [**EDNConInfo**](EDNConInfo.md) |  | 
**Eess** | [**[]EESInfo**](EESInfo.md) | Contains the list of EESs of the EDN. | 
**LifeTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 

## Methods

### NewEDNConfigInfo

`func NewEDNConfigInfo(ednConInfo EDNConInfo, eess []EESInfo, ) *EDNConfigInfo`

NewEDNConfigInfo instantiates a new EDNConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEDNConfigInfoWithDefaults

`func NewEDNConfigInfoWithDefaults() *EDNConfigInfo`

NewEDNConfigInfoWithDefaults instantiates a new EDNConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdnConInfo

`func (o *EDNConfigInfo) GetEdnConInfo() EDNConInfo`

GetEdnConInfo returns the EdnConInfo field if non-nil, zero value otherwise.

### GetEdnConInfoOk

`func (o *EDNConfigInfo) GetEdnConInfoOk() (*EDNConInfo, bool)`

GetEdnConInfoOk returns a tuple with the EdnConInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnConInfo

`func (o *EDNConfigInfo) SetEdnConInfo(v EDNConInfo)`

SetEdnConInfo sets EdnConInfo field to given value.


### GetEess

`func (o *EDNConfigInfo) GetEess() []EESInfo`

GetEess returns the Eess field if non-nil, zero value otherwise.

### GetEessOk

`func (o *EDNConfigInfo) GetEessOk() (*[]EESInfo, bool)`

GetEessOk returns a tuple with the Eess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEess

`func (o *EDNConfigInfo) SetEess(v []EESInfo)`

SetEess sets Eess field to given value.


### GetLifeTime

`func (o *EDNConfigInfo) GetLifeTime() time.Time`

GetLifeTime returns the LifeTime field if non-nil, zero value otherwise.

### GetLifeTimeOk

`func (o *EDNConfigInfo) GetLifeTimeOk() (*time.Time, bool)`

GetLifeTimeOk returns a tuple with the LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeTime

`func (o *EDNConfigInfo) SetLifeTime(v time.Time)`

SetLifeTime sets LifeTime field to given value.

### HasLifeTime

`func (o *EDNConfigInfo) HasLifeTime() bool`

HasLifeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


