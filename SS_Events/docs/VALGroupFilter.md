# VALGroupFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValSvcId** | Pointer to **string** | Identity of the VAL service | [optional] 
**ValGrpIds** | **[]string** | VAL group identifiers that event subscriber wants to know in the interested event.  | 

## Methods

### NewVALGroupFilter

`func NewVALGroupFilter(valGrpIds []string, ) *VALGroupFilter`

NewVALGroupFilter instantiates a new VALGroupFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVALGroupFilterWithDefaults

`func NewVALGroupFilterWithDefaults() *VALGroupFilter`

NewVALGroupFilterWithDefaults instantiates a new VALGroupFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValSvcId

`func (o *VALGroupFilter) GetValSvcId() string`

GetValSvcId returns the ValSvcId field if non-nil, zero value otherwise.

### GetValSvcIdOk

`func (o *VALGroupFilter) GetValSvcIdOk() (*string, bool)`

GetValSvcIdOk returns a tuple with the ValSvcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValSvcId

`func (o *VALGroupFilter) SetValSvcId(v string)`

SetValSvcId sets ValSvcId field to given value.

### HasValSvcId

`func (o *VALGroupFilter) HasValSvcId() bool`

HasValSvcId returns a boolean if a field has been set.

### GetValGrpIds

`func (o *VALGroupFilter) GetValGrpIds() []string`

GetValGrpIds returns the ValGrpIds field if non-nil, zero value otherwise.

### GetValGrpIdsOk

`func (o *VALGroupFilter) GetValGrpIdsOk() (*[]string, bool)`

GetValGrpIdsOk returns a tuple with the ValGrpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValGrpIds

`func (o *VALGroupFilter) SetValGrpIds(v []string)`

SetValGrpIds sets ValGrpIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


