# MLEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MLEvent** | [**NwdafEvent**](NwdafEvent.md) |  | 
**MLEventFilter** | [**EventFilter**](EventFilter.md) |  | 
**TgtUe** | Pointer to [**TargetUeInformation**](TargetUeInformation.md) |  | [optional] 
**MLTargetPeriod** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewMLEventSubscription

`func NewMLEventSubscription(mLEvent NwdafEvent, mLEventFilter EventFilter, ) *MLEventSubscription`

NewMLEventSubscription instantiates a new MLEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMLEventSubscriptionWithDefaults

`func NewMLEventSubscriptionWithDefaults() *MLEventSubscription`

NewMLEventSubscriptionWithDefaults instantiates a new MLEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMLEvent

`func (o *MLEventSubscription) GetMLEvent() NwdafEvent`

GetMLEvent returns the MLEvent field if non-nil, zero value otherwise.

### GetMLEventOk

`func (o *MLEventSubscription) GetMLEventOk() (*NwdafEvent, bool)`

GetMLEventOk returns a tuple with the MLEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLEvent

`func (o *MLEventSubscription) SetMLEvent(v NwdafEvent)`

SetMLEvent sets MLEvent field to given value.


### GetMLEventFilter

`func (o *MLEventSubscription) GetMLEventFilter() EventFilter`

GetMLEventFilter returns the MLEventFilter field if non-nil, zero value otherwise.

### GetMLEventFilterOk

`func (o *MLEventSubscription) GetMLEventFilterOk() (*EventFilter, bool)`

GetMLEventFilterOk returns a tuple with the MLEventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLEventFilter

`func (o *MLEventSubscription) SetMLEventFilter(v EventFilter)`

SetMLEventFilter sets MLEventFilter field to given value.


### GetTgtUe

`func (o *MLEventSubscription) GetTgtUe() TargetUeInformation`

GetTgtUe returns the TgtUe field if non-nil, zero value otherwise.

### GetTgtUeOk

`func (o *MLEventSubscription) GetTgtUeOk() (*TargetUeInformation, bool)`

GetTgtUeOk returns a tuple with the TgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUe

`func (o *MLEventSubscription) SetTgtUe(v TargetUeInformation)`

SetTgtUe sets TgtUe field to given value.

### HasTgtUe

`func (o *MLEventSubscription) HasTgtUe() bool`

HasTgtUe returns a boolean if a field has been set.

### GetMLTargetPeriod

`func (o *MLEventSubscription) GetMLTargetPeriod() TimeWindow`

GetMLTargetPeriod returns the MLTargetPeriod field if non-nil, zero value otherwise.

### GetMLTargetPeriodOk

`func (o *MLEventSubscription) GetMLTargetPeriodOk() (*TimeWindow, bool)`

GetMLTargetPeriodOk returns a tuple with the MLTargetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLTargetPeriod

`func (o *MLEventSubscription) SetMLTargetPeriod(v TimeWindow)`

SetMLTargetPeriod sets MLTargetPeriod field to given value.

### HasMLTargetPeriod

`func (o *MLEventSubscription) HasMLTargetPeriod() bool`

HasMLTargetPeriod returns a boolean if a field has been set.

### GetExpiryTime

`func (o *MLEventSubscription) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *MLEventSubscription) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *MLEventSubscription) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *MLEventSubscription) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


