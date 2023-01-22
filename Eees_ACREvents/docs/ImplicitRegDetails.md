# ImplicitRegDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegId** | **string** | Identifier of the EEC registration. | 
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 

## Methods

### NewImplicitRegDetails

`func NewImplicitRegDetails(regId string, ) *ImplicitRegDetails`

NewImplicitRegDetails instantiates a new ImplicitRegDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImplicitRegDetailsWithDefaults

`func NewImplicitRegDetailsWithDefaults() *ImplicitRegDetails`

NewImplicitRegDetailsWithDefaults instantiates a new ImplicitRegDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegId

`func (o *ImplicitRegDetails) GetRegId() string`

GetRegId returns the RegId field if non-nil, zero value otherwise.

### GetRegIdOk

`func (o *ImplicitRegDetails) GetRegIdOk() (*string, bool)`

GetRegIdOk returns a tuple with the RegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegId

`func (o *ImplicitRegDetails) SetRegId(v string)`

SetRegId sets RegId field to given value.


### GetExpTime

`func (o *ImplicitRegDetails) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *ImplicitRegDetails) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *ImplicitRegDetails) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *ImplicitRegDetails) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


