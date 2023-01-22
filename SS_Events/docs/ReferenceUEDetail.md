# ReferenceUEDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValTgtUe** | [**ValTargetUe**](ValTargetUe.md) |  | 
**ProxRange** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 
**ProxRangeFrac** | Pointer to **float32** | string with format &#39;float&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewReferenceUEDetail

`func NewReferenceUEDetail(valTgtUe ValTargetUe, proxRange int32, ) *ReferenceUEDetail`

NewReferenceUEDetail instantiates a new ReferenceUEDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceUEDetailWithDefaults

`func NewReferenceUEDetailWithDefaults() *ReferenceUEDetail`

NewReferenceUEDetailWithDefaults instantiates a new ReferenceUEDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValTgtUe

`func (o *ReferenceUEDetail) GetValTgtUe() ValTargetUe`

GetValTgtUe returns the ValTgtUe field if non-nil, zero value otherwise.

### GetValTgtUeOk

`func (o *ReferenceUEDetail) GetValTgtUeOk() (*ValTargetUe, bool)`

GetValTgtUeOk returns a tuple with the ValTgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValTgtUe

`func (o *ReferenceUEDetail) SetValTgtUe(v ValTargetUe)`

SetValTgtUe sets ValTgtUe field to given value.


### GetProxRange

`func (o *ReferenceUEDetail) GetProxRange() int32`

GetProxRange returns the ProxRange field if non-nil, zero value otherwise.

### GetProxRangeOk

`func (o *ReferenceUEDetail) GetProxRangeOk() (*int32, bool)`

GetProxRangeOk returns a tuple with the ProxRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxRange

`func (o *ReferenceUEDetail) SetProxRange(v int32)`

SetProxRange sets ProxRange field to given value.


### GetProxRangeFrac

`func (o *ReferenceUEDetail) GetProxRangeFrac() float32`

GetProxRangeFrac returns the ProxRangeFrac field if non-nil, zero value otherwise.

### GetProxRangeFracOk

`func (o *ReferenceUEDetail) GetProxRangeFracOk() (*float32, bool)`

GetProxRangeFracOk returns a tuple with the ProxRangeFrac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxRangeFrac

`func (o *ReferenceUEDetail) SetProxRangeFrac(v float32)`

SetProxRangeFrac sets ProxRangeFrac field to given value.

### HasProxRangeFrac

`func (o *ReferenceUEDetail) HasProxRangeFrac() bool`

HasProxRangeFrac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


