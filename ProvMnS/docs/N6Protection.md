# N6Protection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServAttrCom** | Pointer to [**ServAttrCom**](ServAttrCom.md) |  | [optional] 
**SecFuncList** | Pointer to [**[]SecFunc**](SecFunc.md) |  | [optional] 

## Methods

### NewN6Protection

`func NewN6Protection() *N6Protection`

NewN6Protection instantiates a new N6Protection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN6ProtectionWithDefaults

`func NewN6ProtectionWithDefaults() *N6Protection`

NewN6ProtectionWithDefaults instantiates a new N6Protection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServAttrCom

`func (o *N6Protection) GetServAttrCom() ServAttrCom`

GetServAttrCom returns the ServAttrCom field if non-nil, zero value otherwise.

### GetServAttrComOk

`func (o *N6Protection) GetServAttrComOk() (*ServAttrCom, bool)`

GetServAttrComOk returns a tuple with the ServAttrCom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServAttrCom

`func (o *N6Protection) SetServAttrCom(v ServAttrCom)`

SetServAttrCom sets ServAttrCom field to given value.

### HasServAttrCom

`func (o *N6Protection) HasServAttrCom() bool`

HasServAttrCom returns a boolean if a field has been set.

### GetSecFuncList

`func (o *N6Protection) GetSecFuncList() []SecFunc`

GetSecFuncList returns the SecFuncList field if non-nil, zero value otherwise.

### GetSecFuncListOk

`func (o *N6Protection) GetSecFuncListOk() (*[]SecFunc, bool)`

GetSecFuncListOk returns a tuple with the SecFuncList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecFuncList

`func (o *N6Protection) SetSecFuncList(v []SecFunc)`

SetSecFuncList sets SecFuncList field to given value.

### HasSecFuncList

`func (o *N6Protection) HasSecFuncList() bool`

HasSecFuncList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


