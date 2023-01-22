# StreamInfoWithReportersType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamInfo** | Pointer to [**StreamInfoType**](StreamInfoType.md) |  | [optional] 
**Reporters** | Pointer to [**[]ProducerIdType**](ProducerIdType.md) |  | [optional] 

## Methods

### NewStreamInfoWithReportersType

`func NewStreamInfoWithReportersType() *StreamInfoWithReportersType`

NewStreamInfoWithReportersType instantiates a new StreamInfoWithReportersType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamInfoWithReportersTypeWithDefaults

`func NewStreamInfoWithReportersTypeWithDefaults() *StreamInfoWithReportersType`

NewStreamInfoWithReportersTypeWithDefaults instantiates a new StreamInfoWithReportersType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamInfo

`func (o *StreamInfoWithReportersType) GetStreamInfo() StreamInfoType`

GetStreamInfo returns the StreamInfo field if non-nil, zero value otherwise.

### GetStreamInfoOk

`func (o *StreamInfoWithReportersType) GetStreamInfoOk() (*StreamInfoType, bool)`

GetStreamInfoOk returns a tuple with the StreamInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamInfo

`func (o *StreamInfoWithReportersType) SetStreamInfo(v StreamInfoType)`

SetStreamInfo sets StreamInfo field to given value.

### HasStreamInfo

`func (o *StreamInfoWithReportersType) HasStreamInfo() bool`

HasStreamInfo returns a boolean if a field has been set.

### GetReporters

`func (o *StreamInfoWithReportersType) GetReporters() []ProducerIdType`

GetReporters returns the Reporters field if non-nil, zero value otherwise.

### GetReportersOk

`func (o *StreamInfoWithReportersType) GetReportersOk() (*[]ProducerIdType, bool)`

GetReportersOk returns a tuple with the Reporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporters

`func (o *StreamInfoWithReportersType) SetReporters(v []ProducerIdType)`

SetReporters sets Reporters field to given value.

### HasReporters

`func (o *StreamInfoWithReportersType) HasReporters() bool`

HasReporters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


