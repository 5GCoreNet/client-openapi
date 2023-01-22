# RepositoryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SequenceNumber** | **int32** | Unsigned integer containing the sequence number associated to the current version of Repository Data  | 
**ServiceData** | **string** |  | 

## Methods

### NewRepositoryData

`func NewRepositoryData(sequenceNumber int32, serviceData string, ) *RepositoryData`

NewRepositoryData instantiates a new RepositoryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryDataWithDefaults

`func NewRepositoryDataWithDefaults() *RepositoryData`

NewRepositoryDataWithDefaults instantiates a new RepositoryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequenceNumber

`func (o *RepositoryData) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *RepositoryData) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *RepositoryData) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetServiceData

`func (o *RepositoryData) GetServiceData() string`

GetServiceData returns the ServiceData field if non-nil, zero value otherwise.

### GetServiceDataOk

`func (o *RepositoryData) GetServiceDataOk() (*string, bool)`

GetServiceDataOk returns a tuple with the ServiceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceData

`func (o *RepositoryData) SetServiceData(v string)`

SetServiceData sets ServiceData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


