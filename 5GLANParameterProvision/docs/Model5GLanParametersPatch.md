# Model5GLanParametersPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsis** | Pointer to **map[string]string** | Contains the list of 5G VN Group members, each member is identified by GPSI. Any string value can be used as a key of the map.  | [optional] 
**AppDesps** | Pointer to [**map[string]AppDescriptorRm**](AppDescriptorRm.md) | Describes the operation systems and the corresponding applications for each operation systems. The key of map is osId.  | [optional] 

## Methods

### NewModel5GLanParametersPatch

`func NewModel5GLanParametersPatch() *Model5GLanParametersPatch`

NewModel5GLanParametersPatch instantiates a new Model5GLanParametersPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel5GLanParametersPatchWithDefaults

`func NewModel5GLanParametersPatchWithDefaults() *Model5GLanParametersPatch`

NewModel5GLanParametersPatchWithDefaults instantiates a new Model5GLanParametersPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsis

`func (o *Model5GLanParametersPatch) GetGpsis() map[string]string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *Model5GLanParametersPatch) GetGpsisOk() (*map[string]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *Model5GLanParametersPatch) SetGpsis(v map[string]string)`

SetGpsis sets Gpsis field to given value.

### HasGpsis

`func (o *Model5GLanParametersPatch) HasGpsis() bool`

HasGpsis returns a boolean if a field has been set.

### GetAppDesps

`func (o *Model5GLanParametersPatch) GetAppDesps() map[string]AppDescriptorRm`

GetAppDesps returns the AppDesps field if non-nil, zero value otherwise.

### GetAppDespsOk

`func (o *Model5GLanParametersPatch) GetAppDespsOk() (*map[string]AppDescriptorRm, bool)`

GetAppDespsOk returns a tuple with the AppDesps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDesps

`func (o *Model5GLanParametersPatch) SetAppDesps(v map[string]AppDescriptorRm)`

SetAppDesps sets AppDesps field to given value.

### HasAppDesps

`func (o *Model5GLanParametersPatch) HasAppDesps() bool`

HasAppDesps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


