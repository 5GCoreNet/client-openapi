# AlarmListSingleAllOfAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministrativeState** | Pointer to [**AdministrativeState**](AdministrativeState.md) |  | [optional] 
**OperationalState** | Pointer to [**OperationalState**](OperationalState.md) |  | [optional] 
**NumOfAlarmRecords** | Pointer to **int32** |  | [optional] 
**LastModification** | Pointer to **time.Time** |  | [optional] 
**AlarmRecords** | Pointer to [**map[string]AlarmRecord**](AlarmRecord.md) | This resource represents a map of alarm records. The alarmIds are used as keys in the map. | [optional] 

## Methods

### NewAlarmListSingleAllOfAttributes

`func NewAlarmListSingleAllOfAttributes() *AlarmListSingleAllOfAttributes`

NewAlarmListSingleAllOfAttributes instantiates a new AlarmListSingleAllOfAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmListSingleAllOfAttributesWithDefaults

`func NewAlarmListSingleAllOfAttributesWithDefaults() *AlarmListSingleAllOfAttributes`

NewAlarmListSingleAllOfAttributesWithDefaults instantiates a new AlarmListSingleAllOfAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministrativeState

`func (o *AlarmListSingleAllOfAttributes) GetAdministrativeState() AdministrativeState`

GetAdministrativeState returns the AdministrativeState field if non-nil, zero value otherwise.

### GetAdministrativeStateOk

`func (o *AlarmListSingleAllOfAttributes) GetAdministrativeStateOk() (*AdministrativeState, bool)`

GetAdministrativeStateOk returns a tuple with the AdministrativeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeState

`func (o *AlarmListSingleAllOfAttributes) SetAdministrativeState(v AdministrativeState)`

SetAdministrativeState sets AdministrativeState field to given value.

### HasAdministrativeState

`func (o *AlarmListSingleAllOfAttributes) HasAdministrativeState() bool`

HasAdministrativeState returns a boolean if a field has been set.

### GetOperationalState

`func (o *AlarmListSingleAllOfAttributes) GetOperationalState() OperationalState`

GetOperationalState returns the OperationalState field if non-nil, zero value otherwise.

### GetOperationalStateOk

`func (o *AlarmListSingleAllOfAttributes) GetOperationalStateOk() (*OperationalState, bool)`

GetOperationalStateOk returns a tuple with the OperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalState

`func (o *AlarmListSingleAllOfAttributes) SetOperationalState(v OperationalState)`

SetOperationalState sets OperationalState field to given value.

### HasOperationalState

`func (o *AlarmListSingleAllOfAttributes) HasOperationalState() bool`

HasOperationalState returns a boolean if a field has been set.

### GetNumOfAlarmRecords

`func (o *AlarmListSingleAllOfAttributes) GetNumOfAlarmRecords() int32`

GetNumOfAlarmRecords returns the NumOfAlarmRecords field if non-nil, zero value otherwise.

### GetNumOfAlarmRecordsOk

`func (o *AlarmListSingleAllOfAttributes) GetNumOfAlarmRecordsOk() (*int32, bool)`

GetNumOfAlarmRecordsOk returns a tuple with the NumOfAlarmRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfAlarmRecords

`func (o *AlarmListSingleAllOfAttributes) SetNumOfAlarmRecords(v int32)`

SetNumOfAlarmRecords sets NumOfAlarmRecords field to given value.

### HasNumOfAlarmRecords

`func (o *AlarmListSingleAllOfAttributes) HasNumOfAlarmRecords() bool`

HasNumOfAlarmRecords returns a boolean if a field has been set.

### GetLastModification

`func (o *AlarmListSingleAllOfAttributes) GetLastModification() time.Time`

GetLastModification returns the LastModification field if non-nil, zero value otherwise.

### GetLastModificationOk

`func (o *AlarmListSingleAllOfAttributes) GetLastModificationOk() (*time.Time, bool)`

GetLastModificationOk returns a tuple with the LastModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModification

`func (o *AlarmListSingleAllOfAttributes) SetLastModification(v time.Time)`

SetLastModification sets LastModification field to given value.

### HasLastModification

`func (o *AlarmListSingleAllOfAttributes) HasLastModification() bool`

HasLastModification returns a boolean if a field has been set.

### GetAlarmRecords

`func (o *AlarmListSingleAllOfAttributes) GetAlarmRecords() map[string]AlarmRecord`

GetAlarmRecords returns the AlarmRecords field if non-nil, zero value otherwise.

### GetAlarmRecordsOk

`func (o *AlarmListSingleAllOfAttributes) GetAlarmRecordsOk() (*map[string]AlarmRecord, bool)`

GetAlarmRecordsOk returns a tuple with the AlarmRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmRecords

`func (o *AlarmListSingleAllOfAttributes) SetAlarmRecords(v map[string]AlarmRecord)`

SetAlarmRecords sets AlarmRecords field to given value.

### HasAlarmRecords

`func (o *AlarmListSingleAllOfAttributes) HasAlarmRecords() bool`

HasAlarmRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


