# TrafficSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficClass** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | 
**E2eMaxLatency** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 

## Methods

### NewTrafficSpecification

`func NewTrafficSpecification(trafficClass int32, e2eMaxLatency int32, ) *TrafficSpecification`

NewTrafficSpecification instantiates a new TrafficSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficSpecificationWithDefaults

`func NewTrafficSpecificationWithDefaults() *TrafficSpecification`

NewTrafficSpecificationWithDefaults instantiates a new TrafficSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficClass

`func (o *TrafficSpecification) GetTrafficClass() int32`

GetTrafficClass returns the TrafficClass field if non-nil, zero value otherwise.

### GetTrafficClassOk

`func (o *TrafficSpecification) GetTrafficClassOk() (*int32, bool)`

GetTrafficClassOk returns a tuple with the TrafficClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficClass

`func (o *TrafficSpecification) SetTrafficClass(v int32)`

SetTrafficClass sets TrafficClass field to given value.


### GetE2eMaxLatency

`func (o *TrafficSpecification) GetE2eMaxLatency() int32`

GetE2eMaxLatency returns the E2eMaxLatency field if non-nil, zero value otherwise.

### GetE2eMaxLatencyOk

`func (o *TrafficSpecification) GetE2eMaxLatencyOk() (*int32, bool)`

GetE2eMaxLatencyOk returns a tuple with the E2eMaxLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE2eMaxLatency

`func (o *TrafficSpecification) SetE2eMaxLatency(v int32)`

SetE2eMaxLatency sets E2eMaxLatency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


