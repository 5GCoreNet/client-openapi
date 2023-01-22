# TargetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrgetEASInfo** | Pointer to [**DiscoveredEas**](DiscoveredEas.md) |  | [optional] 
**TrgetEESInfo** | Pointer to [**EDNConfigInfo**](EDNConfigInfo.md) |  | [optional] 

## Methods

### NewTargetInfo

`func NewTargetInfo() *TargetInfo`

NewTargetInfo instantiates a new TargetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetInfoWithDefaults

`func NewTargetInfoWithDefaults() *TargetInfo`

NewTargetInfoWithDefaults instantiates a new TargetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrgetEASInfo

`func (o *TargetInfo) GetTrgetEASInfo() DiscoveredEas`

GetTrgetEASInfo returns the TrgetEASInfo field if non-nil, zero value otherwise.

### GetTrgetEASInfoOk

`func (o *TargetInfo) GetTrgetEASInfoOk() (*DiscoveredEas, bool)`

GetTrgetEASInfoOk returns a tuple with the TrgetEASInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrgetEASInfo

`func (o *TargetInfo) SetTrgetEASInfo(v DiscoveredEas)`

SetTrgetEASInfo sets TrgetEASInfo field to given value.

### HasTrgetEASInfo

`func (o *TargetInfo) HasTrgetEASInfo() bool`

HasTrgetEASInfo returns a boolean if a field has been set.

### GetTrgetEESInfo

`func (o *TargetInfo) GetTrgetEESInfo() EDNConfigInfo`

GetTrgetEESInfo returns the TrgetEESInfo field if non-nil, zero value otherwise.

### GetTrgetEESInfoOk

`func (o *TargetInfo) GetTrgetEESInfoOk() (*EDNConfigInfo, bool)`

GetTrgetEESInfoOk returns a tuple with the TrgetEESInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrgetEESInfo

`func (o *TargetInfo) SetTrgetEESInfo(v EDNConfigInfo)`

SetTrgetEESInfo sets TrgetEESInfo field to given value.

### HasTrgetEESInfo

`func (o *TargetInfo) HasTrgetEESInfo() bool`

HasTrgetEESInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


