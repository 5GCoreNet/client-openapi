# SACEventStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReachedNumUes** | Pointer to [**SACInfo**](SACInfo.md) |  | [optional] 
**ReachedNumPduSess** | Pointer to [**SACInfo**](SACInfo.md) |  | [optional] 

## Methods

### NewSACEventStatus

`func NewSACEventStatus() *SACEventStatus`

NewSACEventStatus instantiates a new SACEventStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSACEventStatusWithDefaults

`func NewSACEventStatusWithDefaults() *SACEventStatus`

NewSACEventStatusWithDefaults instantiates a new SACEventStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReachedNumUes

`func (o *SACEventStatus) GetReachedNumUes() SACInfo`

GetReachedNumUes returns the ReachedNumUes field if non-nil, zero value otherwise.

### GetReachedNumUesOk

`func (o *SACEventStatus) GetReachedNumUesOk() (*SACInfo, bool)`

GetReachedNumUesOk returns a tuple with the ReachedNumUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachedNumUes

`func (o *SACEventStatus) SetReachedNumUes(v SACInfo)`

SetReachedNumUes sets ReachedNumUes field to given value.

### HasReachedNumUes

`func (o *SACEventStatus) HasReachedNumUes() bool`

HasReachedNumUes returns a boolean if a field has been set.

### GetReachedNumPduSess

`func (o *SACEventStatus) GetReachedNumPduSess() SACInfo`

GetReachedNumPduSess returns the ReachedNumPduSess field if non-nil, zero value otherwise.

### GetReachedNumPduSessOk

`func (o *SACEventStatus) GetReachedNumPduSessOk() (*SACInfo, bool)`

GetReachedNumPduSessOk returns a tuple with the ReachedNumPduSess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachedNumPduSess

`func (o *SACEventStatus) SetReachedNumPduSess(v SACInfo)`

SetReachedNumPduSess sets ReachedNumPduSess field to given value.

### HasReachedNumPduSess

`func (o *SACEventStatus) HasReachedNumPduSess() bool`

HasReachedNumPduSess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


