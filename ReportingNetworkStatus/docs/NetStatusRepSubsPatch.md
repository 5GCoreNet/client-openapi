# NetStatusRepSubsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**LocationArea** | Pointer to [**LocationArea**](LocationArea.md) |  | [optional] 
**TimeDuration** | Pointer to **NullableTime** | string with format \&quot;date-time\&quot; as defined in OpenAPI with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**ThresholdValues** | Pointer to **[]int32** |  | [optional] 
**ThresholdTypes** | Pointer to [**[]CongestionType**](CongestionType.md) |  | [optional] 

## Methods

### NewNetStatusRepSubsPatch

`func NewNetStatusRepSubsPatch() *NetStatusRepSubsPatch`

NewNetStatusRepSubsPatch instantiates a new NetStatusRepSubsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetStatusRepSubsPatchWithDefaults

`func NewNetStatusRepSubsPatchWithDefaults() *NetStatusRepSubsPatch`

NewNetStatusRepSubsPatchWithDefaults instantiates a new NetStatusRepSubsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationDestination

`func (o *NetStatusRepSubsPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *NetStatusRepSubsPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *NetStatusRepSubsPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *NetStatusRepSubsPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetLocationArea

`func (o *NetStatusRepSubsPatch) GetLocationArea() LocationArea`

GetLocationArea returns the LocationArea field if non-nil, zero value otherwise.

### GetLocationAreaOk

`func (o *NetStatusRepSubsPatch) GetLocationAreaOk() (*LocationArea, bool)`

GetLocationAreaOk returns a tuple with the LocationArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationArea

`func (o *NetStatusRepSubsPatch) SetLocationArea(v LocationArea)`

SetLocationArea sets LocationArea field to given value.

### HasLocationArea

`func (o *NetStatusRepSubsPatch) HasLocationArea() bool`

HasLocationArea returns a boolean if a field has been set.

### GetTimeDuration

`func (o *NetStatusRepSubsPatch) GetTimeDuration() time.Time`

GetTimeDuration returns the TimeDuration field if non-nil, zero value otherwise.

### GetTimeDurationOk

`func (o *NetStatusRepSubsPatch) GetTimeDurationOk() (*time.Time, bool)`

GetTimeDurationOk returns a tuple with the TimeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeDuration

`func (o *NetStatusRepSubsPatch) SetTimeDuration(v time.Time)`

SetTimeDuration sets TimeDuration field to given value.

### HasTimeDuration

`func (o *NetStatusRepSubsPatch) HasTimeDuration() bool`

HasTimeDuration returns a boolean if a field has been set.

### SetTimeDurationNil

`func (o *NetStatusRepSubsPatch) SetTimeDurationNil(b bool)`

 SetTimeDurationNil sets the value for TimeDuration to be an explicit nil

### UnsetTimeDuration
`func (o *NetStatusRepSubsPatch) UnsetTimeDuration()`

UnsetTimeDuration ensures that no value is present for TimeDuration, not even an explicit nil
### GetThresholdValues

`func (o *NetStatusRepSubsPatch) GetThresholdValues() []int32`

GetThresholdValues returns the ThresholdValues field if non-nil, zero value otherwise.

### GetThresholdValuesOk

`func (o *NetStatusRepSubsPatch) GetThresholdValuesOk() (*[]int32, bool)`

GetThresholdValuesOk returns a tuple with the ThresholdValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValues

`func (o *NetStatusRepSubsPatch) SetThresholdValues(v []int32)`

SetThresholdValues sets ThresholdValues field to given value.

### HasThresholdValues

`func (o *NetStatusRepSubsPatch) HasThresholdValues() bool`

HasThresholdValues returns a boolean if a field has been set.

### GetThresholdTypes

`func (o *NetStatusRepSubsPatch) GetThresholdTypes() []CongestionType`

GetThresholdTypes returns the ThresholdTypes field if non-nil, zero value otherwise.

### GetThresholdTypesOk

`func (o *NetStatusRepSubsPatch) GetThresholdTypesOk() (*[]CongestionType, bool)`

GetThresholdTypesOk returns a tuple with the ThresholdTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdTypes

`func (o *NetStatusRepSubsPatch) SetThresholdTypes(v []CongestionType)`

SetThresholdTypes sets ThresholdTypes field to given value.

### HasThresholdTypes

`func (o *NetStatusRepSubsPatch) HasThresholdTypes() bool`

HasThresholdTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


