# PerUeAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeDest** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**Route** | Pointer to **string** |  | [optional] 
**AvgSpeed** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**TimeOfArrival** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewPerUeAttribute

`func NewPerUeAttribute() *PerUeAttribute`

NewPerUeAttribute instantiates a new PerUeAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerUeAttributeWithDefaults

`func NewPerUeAttributeWithDefaults() *PerUeAttribute`

NewPerUeAttributeWithDefaults instantiates a new PerUeAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeDest

`func (o *PerUeAttribute) GetUeDest() LocationArea5G`

GetUeDest returns the UeDest field if non-nil, zero value otherwise.

### GetUeDestOk

`func (o *PerUeAttribute) GetUeDestOk() (*LocationArea5G, bool)`

GetUeDestOk returns a tuple with the UeDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeDest

`func (o *PerUeAttribute) SetUeDest(v LocationArea5G)`

SetUeDest sets UeDest field to given value.

### HasUeDest

`func (o *PerUeAttribute) HasUeDest() bool`

HasUeDest returns a boolean if a field has been set.

### GetRoute

`func (o *PerUeAttribute) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *PerUeAttribute) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *PerUeAttribute) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *PerUeAttribute) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetAvgSpeed

`func (o *PerUeAttribute) GetAvgSpeed() string`

GetAvgSpeed returns the AvgSpeed field if non-nil, zero value otherwise.

### GetAvgSpeedOk

`func (o *PerUeAttribute) GetAvgSpeedOk() (*string, bool)`

GetAvgSpeedOk returns a tuple with the AvgSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSpeed

`func (o *PerUeAttribute) SetAvgSpeed(v string)`

SetAvgSpeed sets AvgSpeed field to given value.

### HasAvgSpeed

`func (o *PerUeAttribute) HasAvgSpeed() bool`

HasAvgSpeed returns a boolean if a field has been set.

### GetTimeOfArrival

`func (o *PerUeAttribute) GetTimeOfArrival() time.Time`

GetTimeOfArrival returns the TimeOfArrival field if non-nil, zero value otherwise.

### GetTimeOfArrivalOk

`func (o *PerUeAttribute) GetTimeOfArrivalOk() (*time.Time, bool)`

GetTimeOfArrivalOk returns a tuple with the TimeOfArrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfArrival

`func (o *PerUeAttribute) SetTimeOfArrival(v time.Time)`

SetTimeOfArrival sets TimeOfArrival field to given value.

### HasTimeOfArrival

`func (o *PerUeAttribute) HasTimeOfArrival() bool`

HasTimeOfArrival returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


