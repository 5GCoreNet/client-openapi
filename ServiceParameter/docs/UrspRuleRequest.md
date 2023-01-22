# UrspRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficDesc** | Pointer to [**TrafficDescriptorComponents**](TrafficDescriptorComponents.md) |  | [optional] 
**RelatPrecedence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**RouteSelParamSets** | Pointer to [**[]RouteSelectionParameterSet**](RouteSelectionParameterSet.md) | Sets of parameters that may be used to guide the Route Selection Descriptors of the URSP. | [optional] 

## Methods

### NewUrspRuleRequest

`func NewUrspRuleRequest() *UrspRuleRequest`

NewUrspRuleRequest instantiates a new UrspRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrspRuleRequestWithDefaults

`func NewUrspRuleRequestWithDefaults() *UrspRuleRequest`

NewUrspRuleRequestWithDefaults instantiates a new UrspRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficDesc

`func (o *UrspRuleRequest) GetTrafficDesc() TrafficDescriptorComponents`

GetTrafficDesc returns the TrafficDesc field if non-nil, zero value otherwise.

### GetTrafficDescOk

`func (o *UrspRuleRequest) GetTrafficDescOk() (*TrafficDescriptorComponents, bool)`

GetTrafficDescOk returns a tuple with the TrafficDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDesc

`func (o *UrspRuleRequest) SetTrafficDesc(v TrafficDescriptorComponents)`

SetTrafficDesc sets TrafficDesc field to given value.

### HasTrafficDesc

`func (o *UrspRuleRequest) HasTrafficDesc() bool`

HasTrafficDesc returns a boolean if a field has been set.

### GetRelatPrecedence

`func (o *UrspRuleRequest) GetRelatPrecedence() int32`

GetRelatPrecedence returns the RelatPrecedence field if non-nil, zero value otherwise.

### GetRelatPrecedenceOk

`func (o *UrspRuleRequest) GetRelatPrecedenceOk() (*int32, bool)`

GetRelatPrecedenceOk returns a tuple with the RelatPrecedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatPrecedence

`func (o *UrspRuleRequest) SetRelatPrecedence(v int32)`

SetRelatPrecedence sets RelatPrecedence field to given value.

### HasRelatPrecedence

`func (o *UrspRuleRequest) HasRelatPrecedence() bool`

HasRelatPrecedence returns a boolean if a field has been set.

### GetRouteSelParamSets

`func (o *UrspRuleRequest) GetRouteSelParamSets() []RouteSelectionParameterSet`

GetRouteSelParamSets returns the RouteSelParamSets field if non-nil, zero value otherwise.

### GetRouteSelParamSetsOk

`func (o *UrspRuleRequest) GetRouteSelParamSetsOk() (*[]RouteSelectionParameterSet, bool)`

GetRouteSelParamSetsOk returns a tuple with the RouteSelParamSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteSelParamSets

`func (o *UrspRuleRequest) SetRouteSelParamSets(v []RouteSelectionParameterSet)`

SetRouteSelParamSets sets RouteSelParamSets field to given value.

### HasRouteSelParamSets

`func (o *UrspRuleRequest) HasRouteSelParamSets() bool`

HasRouteSelParamSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


