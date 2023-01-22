# SrvccData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StnSr** | **string** | String representing the STN-SR as defined in clause 18.6 of 3GPP TS 23.003. | 
**UeSrvccCapabilities** | Pointer to [**[]SrvccCapability**](SrvccCapability.md) |  | [optional] 

## Methods

### NewSrvccData

`func NewSrvccData(stnSr string, ) *SrvccData`

NewSrvccData instantiates a new SrvccData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrvccDataWithDefaults

`func NewSrvccDataWithDefaults() *SrvccData`

NewSrvccDataWithDefaults instantiates a new SrvccData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStnSr

`func (o *SrvccData) GetStnSr() string`

GetStnSr returns the StnSr field if non-nil, zero value otherwise.

### GetStnSrOk

`func (o *SrvccData) GetStnSrOk() (*string, bool)`

GetStnSrOk returns a tuple with the StnSr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStnSr

`func (o *SrvccData) SetStnSr(v string)`

SetStnSr sets StnSr field to given value.


### GetUeSrvccCapabilities

`func (o *SrvccData) GetUeSrvccCapabilities() []SrvccCapability`

GetUeSrvccCapabilities returns the UeSrvccCapabilities field if non-nil, zero value otherwise.

### GetUeSrvccCapabilitiesOk

`func (o *SrvccData) GetUeSrvccCapabilitiesOk() (*[]SrvccCapability, bool)`

GetUeSrvccCapabilitiesOk returns a tuple with the UeSrvccCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeSrvccCapabilities

`func (o *SrvccData) SetUeSrvccCapabilities(v []SrvccCapability)`

SetUeSrvccCapabilities sets UeSrvccCapabilities field to given value.

### HasUeSrvccCapabilities

`func (o *SrvccData) HasUeSrvccCapabilities() bool`

HasUeSrvccCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


