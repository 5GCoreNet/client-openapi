# MonitorFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Idnts** | Pointer to [**[]ValTargetUe**](ValTargetUe.md) | List of VAL User or UE IDs whose events monitoring is requested. | [optional] 
**ValSvcId** | Pointer to **string** | Identity of the VAL service. | [optional] 
**ValGrpId** | Pointer to **string** | Identity of the group of the target UEs. | [optional] 
**ProfId** | Pointer to **string** | The monitoring profile ID identifying a list of monitoring, analytics events. | [optional] 
**ValCnds** | Pointer to [**[]ValidityConditions**](ValidityConditions.md) | The temporal,spatial conditions for the events to be considered valid. | [optional] 
**EvntDets** | Pointer to [**[]MonitorEvents**](MonitorEvents.md) | List of monitoring, analytics events to be monitored. | [optional] 

## Methods

### NewMonitorFilter

`func NewMonitorFilter() *MonitorFilter`

NewMonitorFilter instantiates a new MonitorFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorFilterWithDefaults

`func NewMonitorFilterWithDefaults() *MonitorFilter`

NewMonitorFilterWithDefaults instantiates a new MonitorFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdnts

`func (o *MonitorFilter) GetIdnts() []ValTargetUe`

GetIdnts returns the Idnts field if non-nil, zero value otherwise.

### GetIdntsOk

`func (o *MonitorFilter) GetIdntsOk() (*[]ValTargetUe, bool)`

GetIdntsOk returns a tuple with the Idnts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdnts

`func (o *MonitorFilter) SetIdnts(v []ValTargetUe)`

SetIdnts sets Idnts field to given value.

### HasIdnts

`func (o *MonitorFilter) HasIdnts() bool`

HasIdnts returns a boolean if a field has been set.

### GetValSvcId

`func (o *MonitorFilter) GetValSvcId() string`

GetValSvcId returns the ValSvcId field if non-nil, zero value otherwise.

### GetValSvcIdOk

`func (o *MonitorFilter) GetValSvcIdOk() (*string, bool)`

GetValSvcIdOk returns a tuple with the ValSvcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValSvcId

`func (o *MonitorFilter) SetValSvcId(v string)`

SetValSvcId sets ValSvcId field to given value.

### HasValSvcId

`func (o *MonitorFilter) HasValSvcId() bool`

HasValSvcId returns a boolean if a field has been set.

### GetValGrpId

`func (o *MonitorFilter) GetValGrpId() string`

GetValGrpId returns the ValGrpId field if non-nil, zero value otherwise.

### GetValGrpIdOk

`func (o *MonitorFilter) GetValGrpIdOk() (*string, bool)`

GetValGrpIdOk returns a tuple with the ValGrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValGrpId

`func (o *MonitorFilter) SetValGrpId(v string)`

SetValGrpId sets ValGrpId field to given value.

### HasValGrpId

`func (o *MonitorFilter) HasValGrpId() bool`

HasValGrpId returns a boolean if a field has been set.

### GetProfId

`func (o *MonitorFilter) GetProfId() string`

GetProfId returns the ProfId field if non-nil, zero value otherwise.

### GetProfIdOk

`func (o *MonitorFilter) GetProfIdOk() (*string, bool)`

GetProfIdOk returns a tuple with the ProfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfId

`func (o *MonitorFilter) SetProfId(v string)`

SetProfId sets ProfId field to given value.

### HasProfId

`func (o *MonitorFilter) HasProfId() bool`

HasProfId returns a boolean if a field has been set.

### GetValCnds

`func (o *MonitorFilter) GetValCnds() []ValidityConditions`

GetValCnds returns the ValCnds field if non-nil, zero value otherwise.

### GetValCndsOk

`func (o *MonitorFilter) GetValCndsOk() (*[]ValidityConditions, bool)`

GetValCndsOk returns a tuple with the ValCnds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValCnds

`func (o *MonitorFilter) SetValCnds(v []ValidityConditions)`

SetValCnds sets ValCnds field to given value.

### HasValCnds

`func (o *MonitorFilter) HasValCnds() bool`

HasValCnds returns a boolean if a field has been set.

### GetEvntDets

`func (o *MonitorFilter) GetEvntDets() []MonitorEvents`

GetEvntDets returns the EvntDets field if non-nil, zero value otherwise.

### GetEvntDetsOk

`func (o *MonitorFilter) GetEvntDetsOk() (*[]MonitorEvents, bool)`

GetEvntDetsOk returns a tuple with the EvntDets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvntDets

`func (o *MonitorFilter) SetEvntDets(v []MonitorEvents)`

SetEvntDets sets EvntDets field to given value.

### HasEvntDets

`func (o *MonitorFilter) HasEvntDets() bool`

HasEvntDets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


