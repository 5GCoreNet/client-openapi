# TscStreamAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamSpec** | [**StreamSpecification**](StreamSpecification.md) |  | 
**TrafficSpecs** | [**[]TrafficSpecification**](TrafficSpecification.md) |  | 

## Methods

### NewTscStreamAvailability

`func NewTscStreamAvailability(streamSpec StreamSpecification, trafficSpecs []TrafficSpecification, ) *TscStreamAvailability`

NewTscStreamAvailability instantiates a new TscStreamAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTscStreamAvailabilityWithDefaults

`func NewTscStreamAvailabilityWithDefaults() *TscStreamAvailability`

NewTscStreamAvailabilityWithDefaults instantiates a new TscStreamAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamSpec

`func (o *TscStreamAvailability) GetStreamSpec() StreamSpecification`

GetStreamSpec returns the StreamSpec field if non-nil, zero value otherwise.

### GetStreamSpecOk

`func (o *TscStreamAvailability) GetStreamSpecOk() (*StreamSpecification, bool)`

GetStreamSpecOk returns a tuple with the StreamSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamSpec

`func (o *TscStreamAvailability) SetStreamSpec(v StreamSpecification)`

SetStreamSpec sets StreamSpec field to given value.


### GetTrafficSpecs

`func (o *TscStreamAvailability) GetTrafficSpecs() []TrafficSpecification`

GetTrafficSpecs returns the TrafficSpecs field if non-nil, zero value otherwise.

### GetTrafficSpecsOk

`func (o *TscStreamAvailability) GetTrafficSpecsOk() (*[]TrafficSpecification, bool)`

GetTrafficSpecsOk returns a tuple with the TrafficSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSpecs

`func (o *TscStreamAvailability) SetTrafficSpecs(v []TrafficSpecification)`

SetTrafficSpecs sets TrafficSpecs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


