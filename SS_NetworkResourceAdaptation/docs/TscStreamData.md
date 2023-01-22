# TscStreamData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamSpec** | [**StreamSpecification**](StreamSpecification.md) |  | 
**TrafficSpecInfo** | [**TrafficSpecInformation**](TrafficSpecInformation.md) |  | 

## Methods

### NewTscStreamData

`func NewTscStreamData(streamSpec StreamSpecification, trafficSpecInfo TrafficSpecInformation, ) *TscStreamData`

NewTscStreamData instantiates a new TscStreamData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTscStreamDataWithDefaults

`func NewTscStreamDataWithDefaults() *TscStreamData`

NewTscStreamDataWithDefaults instantiates a new TscStreamData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamSpec

`func (o *TscStreamData) GetStreamSpec() StreamSpecification`

GetStreamSpec returns the StreamSpec field if non-nil, zero value otherwise.

### GetStreamSpecOk

`func (o *TscStreamData) GetStreamSpecOk() (*StreamSpecification, bool)`

GetStreamSpecOk returns a tuple with the StreamSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamSpec

`func (o *TscStreamData) SetStreamSpec(v StreamSpecification)`

SetStreamSpec sets StreamSpec field to given value.


### GetTrafficSpecInfo

`func (o *TscStreamData) GetTrafficSpecInfo() TrafficSpecInformation`

GetTrafficSpecInfo returns the TrafficSpecInfo field if non-nil, zero value otherwise.

### GetTrafficSpecInfoOk

`func (o *TscStreamData) GetTrafficSpecInfoOk() (*TrafficSpecInformation, bool)`

GetTrafficSpecInfoOk returns a tuple with the TrafficSpecInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSpecInfo

`func (o *TscStreamData) SetTrafficSpecInfo(v TrafficSpecInformation)`

SetTrafficSpecInfo sets TrafficSpecInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


