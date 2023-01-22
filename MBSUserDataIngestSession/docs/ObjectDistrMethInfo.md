# ObjectDistrMethInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatingMode** | Pointer to [**ObjDistributionOperatingMode**](ObjDistributionOperatingMode.md) |  | [optional] 
**ObjAcqMethod** | Pointer to [**ObjAcquisitionMethod**](ObjAcquisitionMethod.md) |  | [optional] 
**ObjAcqIds** | **[]string** |  | 
**ObjIngUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ObjDistrUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ObjRepairUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

## Methods

### NewObjectDistrMethInfo

`func NewObjectDistrMethInfo(objAcqIds []string, ) *ObjectDistrMethInfo`

NewObjectDistrMethInfo instantiates a new ObjectDistrMethInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectDistrMethInfoWithDefaults

`func NewObjectDistrMethInfoWithDefaults() *ObjectDistrMethInfo`

NewObjectDistrMethInfoWithDefaults instantiates a new ObjectDistrMethInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatingMode

`func (o *ObjectDistrMethInfo) GetOperatingMode() ObjDistributionOperatingMode`

GetOperatingMode returns the OperatingMode field if non-nil, zero value otherwise.

### GetOperatingModeOk

`func (o *ObjectDistrMethInfo) GetOperatingModeOk() (*ObjDistributionOperatingMode, bool)`

GetOperatingModeOk returns a tuple with the OperatingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingMode

`func (o *ObjectDistrMethInfo) SetOperatingMode(v ObjDistributionOperatingMode)`

SetOperatingMode sets OperatingMode field to given value.

### HasOperatingMode

`func (o *ObjectDistrMethInfo) HasOperatingMode() bool`

HasOperatingMode returns a boolean if a field has been set.

### GetObjAcqMethod

`func (o *ObjectDistrMethInfo) GetObjAcqMethod() ObjAcquisitionMethod`

GetObjAcqMethod returns the ObjAcqMethod field if non-nil, zero value otherwise.

### GetObjAcqMethodOk

`func (o *ObjectDistrMethInfo) GetObjAcqMethodOk() (*ObjAcquisitionMethod, bool)`

GetObjAcqMethodOk returns a tuple with the ObjAcqMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAcqMethod

`func (o *ObjectDistrMethInfo) SetObjAcqMethod(v ObjAcquisitionMethod)`

SetObjAcqMethod sets ObjAcqMethod field to given value.

### HasObjAcqMethod

`func (o *ObjectDistrMethInfo) HasObjAcqMethod() bool`

HasObjAcqMethod returns a boolean if a field has been set.

### GetObjAcqIds

`func (o *ObjectDistrMethInfo) GetObjAcqIds() []string`

GetObjAcqIds returns the ObjAcqIds field if non-nil, zero value otherwise.

### GetObjAcqIdsOk

`func (o *ObjectDistrMethInfo) GetObjAcqIdsOk() (*[]string, bool)`

GetObjAcqIdsOk returns a tuple with the ObjAcqIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAcqIds

`func (o *ObjectDistrMethInfo) SetObjAcqIds(v []string)`

SetObjAcqIds sets ObjAcqIds field to given value.


### GetObjIngUri

`func (o *ObjectDistrMethInfo) GetObjIngUri() string`

GetObjIngUri returns the ObjIngUri field if non-nil, zero value otherwise.

### GetObjIngUriOk

`func (o *ObjectDistrMethInfo) GetObjIngUriOk() (*string, bool)`

GetObjIngUriOk returns a tuple with the ObjIngUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjIngUri

`func (o *ObjectDistrMethInfo) SetObjIngUri(v string)`

SetObjIngUri sets ObjIngUri field to given value.

### HasObjIngUri

`func (o *ObjectDistrMethInfo) HasObjIngUri() bool`

HasObjIngUri returns a boolean if a field has been set.

### GetObjDistrUri

`func (o *ObjectDistrMethInfo) GetObjDistrUri() string`

GetObjDistrUri returns the ObjDistrUri field if non-nil, zero value otherwise.

### GetObjDistrUriOk

`func (o *ObjectDistrMethInfo) GetObjDistrUriOk() (*string, bool)`

GetObjDistrUriOk returns a tuple with the ObjDistrUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDistrUri

`func (o *ObjectDistrMethInfo) SetObjDistrUri(v string)`

SetObjDistrUri sets ObjDistrUri field to given value.

### HasObjDistrUri

`func (o *ObjectDistrMethInfo) HasObjDistrUri() bool`

HasObjDistrUri returns a boolean if a field has been set.

### GetObjRepairUri

`func (o *ObjectDistrMethInfo) GetObjRepairUri() string`

GetObjRepairUri returns the ObjRepairUri field if non-nil, zero value otherwise.

### GetObjRepairUriOk

`func (o *ObjectDistrMethInfo) GetObjRepairUriOk() (*string, bool)`

GetObjRepairUriOk returns a tuple with the ObjRepairUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjRepairUri

`func (o *ObjectDistrMethInfo) SetObjRepairUri(v string)`

SetObjRepairUri sets ObjRepairUri field to given value.

### HasObjRepairUri

`func (o *ObjectDistrMethInfo) HasObjRepairUri() bool`

HasObjRepairUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


