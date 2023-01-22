# CalledIdentityChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalledIdentity** | Pointer to **string** |  | [optional] 
**ChangeTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewCalledIdentityChange

`func NewCalledIdentityChange() *CalledIdentityChange`

NewCalledIdentityChange instantiates a new CalledIdentityChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalledIdentityChangeWithDefaults

`func NewCalledIdentityChangeWithDefaults() *CalledIdentityChange`

NewCalledIdentityChangeWithDefaults instantiates a new CalledIdentityChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalledIdentity

`func (o *CalledIdentityChange) GetCalledIdentity() string`

GetCalledIdentity returns the CalledIdentity field if non-nil, zero value otherwise.

### GetCalledIdentityOk

`func (o *CalledIdentityChange) GetCalledIdentityOk() (*string, bool)`

GetCalledIdentityOk returns a tuple with the CalledIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalledIdentity

`func (o *CalledIdentityChange) SetCalledIdentity(v string)`

SetCalledIdentity sets CalledIdentity field to given value.

### HasCalledIdentity

`func (o *CalledIdentityChange) HasCalledIdentity() bool`

HasCalledIdentity returns a boolean if a field has been set.

### GetChangeTime

`func (o *CalledIdentityChange) GetChangeTime() time.Time`

GetChangeTime returns the ChangeTime field if non-nil, zero value otherwise.

### GetChangeTimeOk

`func (o *CalledIdentityChange) GetChangeTimeOk() (*time.Time, bool)`

GetChangeTimeOk returns a tuple with the ChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeTime

`func (o *CalledIdentityChange) SetChangeTime(v time.Time)`

SetChangeTime sets ChangeTime field to given value.

### HasChangeTime

`func (o *CalledIdentityChange) HasChangeTime() bool`

HasChangeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


