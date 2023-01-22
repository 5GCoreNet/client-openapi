# UiccConfigurationParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingId** | Pointer to **string** | Represents a routing indicator. | [optional] 
**SteeringContainer** | Pointer to [**[]SteeringInfo**](SteeringInfo.md) |  | [optional] 
**ExtendedSteeringContainer** | Pointer to [**ExtendedSteeringContainer**](ExtendedSteeringContainer.md) |  | [optional] 

## Methods

### NewUiccConfigurationParameter

`func NewUiccConfigurationParameter() *UiccConfigurationParameter`

NewUiccConfigurationParameter instantiates a new UiccConfigurationParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiccConfigurationParameterWithDefaults

`func NewUiccConfigurationParameterWithDefaults() *UiccConfigurationParameter`

NewUiccConfigurationParameterWithDefaults instantiates a new UiccConfigurationParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutingId

`func (o *UiccConfigurationParameter) GetRoutingId() string`

GetRoutingId returns the RoutingId field if non-nil, zero value otherwise.

### GetRoutingIdOk

`func (o *UiccConfigurationParameter) GetRoutingIdOk() (*string, bool)`

GetRoutingIdOk returns a tuple with the RoutingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingId

`func (o *UiccConfigurationParameter) SetRoutingId(v string)`

SetRoutingId sets RoutingId field to given value.

### HasRoutingId

`func (o *UiccConfigurationParameter) HasRoutingId() bool`

HasRoutingId returns a boolean if a field has been set.

### GetSteeringContainer

`func (o *UiccConfigurationParameter) GetSteeringContainer() []SteeringInfo`

GetSteeringContainer returns the SteeringContainer field if non-nil, zero value otherwise.

### GetSteeringContainerOk

`func (o *UiccConfigurationParameter) GetSteeringContainerOk() (*[]SteeringInfo, bool)`

GetSteeringContainerOk returns a tuple with the SteeringContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteeringContainer

`func (o *UiccConfigurationParameter) SetSteeringContainer(v []SteeringInfo)`

SetSteeringContainer sets SteeringContainer field to given value.

### HasSteeringContainer

`func (o *UiccConfigurationParameter) HasSteeringContainer() bool`

HasSteeringContainer returns a boolean if a field has been set.

### GetExtendedSteeringContainer

`func (o *UiccConfigurationParameter) GetExtendedSteeringContainer() ExtendedSteeringContainer`

GetExtendedSteeringContainer returns the ExtendedSteeringContainer field if non-nil, zero value otherwise.

### GetExtendedSteeringContainerOk

`func (o *UiccConfigurationParameter) GetExtendedSteeringContainerOk() (*ExtendedSteeringContainer, bool)`

GetExtendedSteeringContainerOk returns a tuple with the ExtendedSteeringContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedSteeringContainer

`func (o *UiccConfigurationParameter) SetExtendedSteeringContainer(v ExtendedSteeringContainer)`

SetExtendedSteeringContainer sets ExtendedSteeringContainer field to given value.

### HasExtendedSteeringContainer

`func (o *UiccConfigurationParameter) HasExtendedSteeringContainer() bool`

HasExtendedSteeringContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


