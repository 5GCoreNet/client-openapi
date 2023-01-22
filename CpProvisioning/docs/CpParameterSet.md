# CpParameterSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SetId** | **string** | SCS/AS-chosen correlator provided by the SCS/AS in the request to create a resource fo CP parameter set(s). | 
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**ValidityTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**PeriodicCommunicationIndicator** | Pointer to [**CommunicationIndicator**](CommunicationIndicator.md) |  | [optional] 
**CommunicationDurationTime** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**PeriodicTime** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**ScheduledCommunicationTime** | Pointer to [**ScheduledCommunicationTime**](ScheduledCommunicationTime.md) |  | [optional] 
**ScheduledCommunicationType** | Pointer to [**ScheduledCommunicationType**](ScheduledCommunicationType.md) |  | [optional] 
**StationaryIndication** | Pointer to [**StationaryIndication**](StationaryIndication.md) |  | [optional] 
**BatteryInds** | Pointer to [**[]BatteryIndication**](BatteryIndication.md) |  | [optional] 
**TrafficProfile** | Pointer to [**TrafficProfile**](TrafficProfile.md) |  | [optional] 
**ExpectedUmts** | Pointer to [**[]UmtLocationArea5G**](UmtLocationArea5G.md) | Identifies the UE&#39;s expected geographical movement. The attribute is only applicable in 5G. | [optional] 
**ExpectedUmtDays** | Pointer to **int32** | integer between and including 1 and 7 denoting a weekday. 1 shall indicate Monday, and the subsequent weekdays shall be indicated with the next higher numbers. 7 shall indicate Sunday. | [optional] 

## Methods

### NewCpParameterSet

`func NewCpParameterSet(setId string, ) *CpParameterSet`

NewCpParameterSet instantiates a new CpParameterSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpParameterSetWithDefaults

`func NewCpParameterSetWithDefaults() *CpParameterSet`

NewCpParameterSetWithDefaults instantiates a new CpParameterSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetId

`func (o *CpParameterSet) GetSetId() string`

GetSetId returns the SetId field if non-nil, zero value otherwise.

### GetSetIdOk

`func (o *CpParameterSet) GetSetIdOk() (*string, bool)`

GetSetIdOk returns a tuple with the SetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetId

`func (o *CpParameterSet) SetSetId(v string)`

SetSetId sets SetId field to given value.


### GetSelf

`func (o *CpParameterSet) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *CpParameterSet) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *CpParameterSet) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *CpParameterSet) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetValidityTime

`func (o *CpParameterSet) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *CpParameterSet) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *CpParameterSet) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *CpParameterSet) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.

### GetPeriodicCommunicationIndicator

`func (o *CpParameterSet) GetPeriodicCommunicationIndicator() CommunicationIndicator`

GetPeriodicCommunicationIndicator returns the PeriodicCommunicationIndicator field if non-nil, zero value otherwise.

### GetPeriodicCommunicationIndicatorOk

`func (o *CpParameterSet) GetPeriodicCommunicationIndicatorOk() (*CommunicationIndicator, bool)`

GetPeriodicCommunicationIndicatorOk returns a tuple with the PeriodicCommunicationIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicCommunicationIndicator

`func (o *CpParameterSet) SetPeriodicCommunicationIndicator(v CommunicationIndicator)`

SetPeriodicCommunicationIndicator sets PeriodicCommunicationIndicator field to given value.

### HasPeriodicCommunicationIndicator

`func (o *CpParameterSet) HasPeriodicCommunicationIndicator() bool`

HasPeriodicCommunicationIndicator returns a boolean if a field has been set.

### GetCommunicationDurationTime

`func (o *CpParameterSet) GetCommunicationDurationTime() int32`

GetCommunicationDurationTime returns the CommunicationDurationTime field if non-nil, zero value otherwise.

### GetCommunicationDurationTimeOk

`func (o *CpParameterSet) GetCommunicationDurationTimeOk() (*int32, bool)`

GetCommunicationDurationTimeOk returns a tuple with the CommunicationDurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDurationTime

`func (o *CpParameterSet) SetCommunicationDurationTime(v int32)`

SetCommunicationDurationTime sets CommunicationDurationTime field to given value.

### HasCommunicationDurationTime

`func (o *CpParameterSet) HasCommunicationDurationTime() bool`

HasCommunicationDurationTime returns a boolean if a field has been set.

### GetPeriodicTime

`func (o *CpParameterSet) GetPeriodicTime() int32`

GetPeriodicTime returns the PeriodicTime field if non-nil, zero value otherwise.

### GetPeriodicTimeOk

`func (o *CpParameterSet) GetPeriodicTimeOk() (*int32, bool)`

GetPeriodicTimeOk returns a tuple with the PeriodicTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicTime

`func (o *CpParameterSet) SetPeriodicTime(v int32)`

SetPeriodicTime sets PeriodicTime field to given value.

### HasPeriodicTime

`func (o *CpParameterSet) HasPeriodicTime() bool`

HasPeriodicTime returns a boolean if a field has been set.

### GetScheduledCommunicationTime

`func (o *CpParameterSet) GetScheduledCommunicationTime() ScheduledCommunicationTime`

GetScheduledCommunicationTime returns the ScheduledCommunicationTime field if non-nil, zero value otherwise.

### GetScheduledCommunicationTimeOk

`func (o *CpParameterSet) GetScheduledCommunicationTimeOk() (*ScheduledCommunicationTime, bool)`

GetScheduledCommunicationTimeOk returns a tuple with the ScheduledCommunicationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledCommunicationTime

`func (o *CpParameterSet) SetScheduledCommunicationTime(v ScheduledCommunicationTime)`

SetScheduledCommunicationTime sets ScheduledCommunicationTime field to given value.

### HasScheduledCommunicationTime

`func (o *CpParameterSet) HasScheduledCommunicationTime() bool`

HasScheduledCommunicationTime returns a boolean if a field has been set.

### GetScheduledCommunicationType

`func (o *CpParameterSet) GetScheduledCommunicationType() ScheduledCommunicationType`

GetScheduledCommunicationType returns the ScheduledCommunicationType field if non-nil, zero value otherwise.

### GetScheduledCommunicationTypeOk

`func (o *CpParameterSet) GetScheduledCommunicationTypeOk() (*ScheduledCommunicationType, bool)`

GetScheduledCommunicationTypeOk returns a tuple with the ScheduledCommunicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledCommunicationType

`func (o *CpParameterSet) SetScheduledCommunicationType(v ScheduledCommunicationType)`

SetScheduledCommunicationType sets ScheduledCommunicationType field to given value.

### HasScheduledCommunicationType

`func (o *CpParameterSet) HasScheduledCommunicationType() bool`

HasScheduledCommunicationType returns a boolean if a field has been set.

### GetStationaryIndication

`func (o *CpParameterSet) GetStationaryIndication() StationaryIndication`

GetStationaryIndication returns the StationaryIndication field if non-nil, zero value otherwise.

### GetStationaryIndicationOk

`func (o *CpParameterSet) GetStationaryIndicationOk() (*StationaryIndication, bool)`

GetStationaryIndicationOk returns a tuple with the StationaryIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationaryIndication

`func (o *CpParameterSet) SetStationaryIndication(v StationaryIndication)`

SetStationaryIndication sets StationaryIndication field to given value.

### HasStationaryIndication

`func (o *CpParameterSet) HasStationaryIndication() bool`

HasStationaryIndication returns a boolean if a field has been set.

### GetBatteryInds

`func (o *CpParameterSet) GetBatteryInds() []BatteryIndication`

GetBatteryInds returns the BatteryInds field if non-nil, zero value otherwise.

### GetBatteryIndsOk

`func (o *CpParameterSet) GetBatteryIndsOk() (*[]BatteryIndication, bool)`

GetBatteryIndsOk returns a tuple with the BatteryInds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryInds

`func (o *CpParameterSet) SetBatteryInds(v []BatteryIndication)`

SetBatteryInds sets BatteryInds field to given value.

### HasBatteryInds

`func (o *CpParameterSet) HasBatteryInds() bool`

HasBatteryInds returns a boolean if a field has been set.

### GetTrafficProfile

`func (o *CpParameterSet) GetTrafficProfile() TrafficProfile`

GetTrafficProfile returns the TrafficProfile field if non-nil, zero value otherwise.

### GetTrafficProfileOk

`func (o *CpParameterSet) GetTrafficProfileOk() (*TrafficProfile, bool)`

GetTrafficProfileOk returns a tuple with the TrafficProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficProfile

`func (o *CpParameterSet) SetTrafficProfile(v TrafficProfile)`

SetTrafficProfile sets TrafficProfile field to given value.

### HasTrafficProfile

`func (o *CpParameterSet) HasTrafficProfile() bool`

HasTrafficProfile returns a boolean if a field has been set.

### GetExpectedUmts

`func (o *CpParameterSet) GetExpectedUmts() []UmtLocationArea5G`

GetExpectedUmts returns the ExpectedUmts field if non-nil, zero value otherwise.

### GetExpectedUmtsOk

`func (o *CpParameterSet) GetExpectedUmtsOk() (*[]UmtLocationArea5G, bool)`

GetExpectedUmtsOk returns a tuple with the ExpectedUmts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUmts

`func (o *CpParameterSet) SetExpectedUmts(v []UmtLocationArea5G)`

SetExpectedUmts sets ExpectedUmts field to given value.

### HasExpectedUmts

`func (o *CpParameterSet) HasExpectedUmts() bool`

HasExpectedUmts returns a boolean if a field has been set.

### GetExpectedUmtDays

`func (o *CpParameterSet) GetExpectedUmtDays() int32`

GetExpectedUmtDays returns the ExpectedUmtDays field if non-nil, zero value otherwise.

### GetExpectedUmtDaysOk

`func (o *CpParameterSet) GetExpectedUmtDaysOk() (*int32, bool)`

GetExpectedUmtDaysOk returns a tuple with the ExpectedUmtDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUmtDays

`func (o *CpParameterSet) SetExpectedUmtDays(v int32)`

SetExpectedUmtDays sets ExpectedUmtDays field to given value.

### HasExpectedUmtDays

`func (o *CpParameterSet) HasExpectedUmtDays() bool`

HasExpectedUmtDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


