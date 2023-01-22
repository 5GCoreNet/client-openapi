# AmEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**AmEvent**](AmEvent.md) |  | 
**AppliedCov** | Pointer to [**ServiceAreaCoverageInfo**](ServiceAreaCoverageInfo.md) |  | [optional] 
**PduidInfo** | Pointer to [**PduidInformation**](PduidInformation.md) |  | [optional] 

## Methods

### NewAmEventNotification

`func NewAmEventNotification(event AmEvent, ) *AmEventNotification`

NewAmEventNotification instantiates a new AmEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmEventNotificationWithDefaults

`func NewAmEventNotificationWithDefaults() *AmEventNotification`

NewAmEventNotificationWithDefaults instantiates a new AmEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *AmEventNotification) GetEvent() AmEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AmEventNotification) GetEventOk() (*AmEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AmEventNotification) SetEvent(v AmEvent)`

SetEvent sets Event field to given value.


### GetAppliedCov

`func (o *AmEventNotification) GetAppliedCov() ServiceAreaCoverageInfo`

GetAppliedCov returns the AppliedCov field if non-nil, zero value otherwise.

### GetAppliedCovOk

`func (o *AmEventNotification) GetAppliedCovOk() (*ServiceAreaCoverageInfo, bool)`

GetAppliedCovOk returns a tuple with the AppliedCov field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCov

`func (o *AmEventNotification) SetAppliedCov(v ServiceAreaCoverageInfo)`

SetAppliedCov sets AppliedCov field to given value.

### HasAppliedCov

`func (o *AmEventNotification) HasAppliedCov() bool`

HasAppliedCov returns a boolean if a field has been set.

### GetPduidInfo

`func (o *AmEventNotification) GetPduidInfo() PduidInformation`

GetPduidInfo returns the PduidInfo field if non-nil, zero value otherwise.

### GetPduidInfoOk

`func (o *AmEventNotification) GetPduidInfoOk() (*PduidInformation, bool)`

GetPduidInfoOk returns a tuple with the PduidInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduidInfo

`func (o *AmEventNotification) SetPduidInfo(v PduidInformation)`

SetPduidInfo sets PduidInfo field to given value.

### HasPduidInfo

`func (o *AmEventNotification) HasPduidInfo() bool`

HasPduidInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


