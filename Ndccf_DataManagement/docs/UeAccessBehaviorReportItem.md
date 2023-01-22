# UeAccessBehaviorReportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateTransitionType** | [**AccessStateTransitionType**](AccessStateTransitionType.md) |  | 
**Spacing** | **int32** | indicating a time in seconds. | 
**Duration** | **int32** | indicating a time in seconds. | 

## Methods

### NewUeAccessBehaviorReportItem

`func NewUeAccessBehaviorReportItem(stateTransitionType AccessStateTransitionType, spacing int32, duration int32, ) *UeAccessBehaviorReportItem`

NewUeAccessBehaviorReportItem instantiates a new UeAccessBehaviorReportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeAccessBehaviorReportItemWithDefaults

`func NewUeAccessBehaviorReportItemWithDefaults() *UeAccessBehaviorReportItem`

NewUeAccessBehaviorReportItemWithDefaults instantiates a new UeAccessBehaviorReportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStateTransitionType

`func (o *UeAccessBehaviorReportItem) GetStateTransitionType() AccessStateTransitionType`

GetStateTransitionType returns the StateTransitionType field if non-nil, zero value otherwise.

### GetStateTransitionTypeOk

`func (o *UeAccessBehaviorReportItem) GetStateTransitionTypeOk() (*AccessStateTransitionType, bool)`

GetStateTransitionTypeOk returns a tuple with the StateTransitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTransitionType

`func (o *UeAccessBehaviorReportItem) SetStateTransitionType(v AccessStateTransitionType)`

SetStateTransitionType sets StateTransitionType field to given value.


### GetSpacing

`func (o *UeAccessBehaviorReportItem) GetSpacing() int32`

GetSpacing returns the Spacing field if non-nil, zero value otherwise.

### GetSpacingOk

`func (o *UeAccessBehaviorReportItem) GetSpacingOk() (*int32, bool)`

GetSpacingOk returns a tuple with the Spacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacing

`func (o *UeAccessBehaviorReportItem) SetSpacing(v int32)`

SetSpacing sets Spacing field to given value.


### GetDuration

`func (o *UeAccessBehaviorReportItem) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UeAccessBehaviorReportItem) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UeAccessBehaviorReportItem) SetDuration(v int32)`

SetDuration sets Duration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


