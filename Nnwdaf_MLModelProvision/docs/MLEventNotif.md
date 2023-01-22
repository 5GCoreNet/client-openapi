# MLEventNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**NwdafEvent**](NwdafEvent.md) |  | 
**NotifCorreId** | Pointer to **string** |  | [optional] 
**MLFileAddr** | [**MLModelAddr**](MLModelAddr.md) |  | 
**ValidityPeriod** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**SpatialValidity** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 

## Methods

### NewMLEventNotif

`func NewMLEventNotif(event NwdafEvent, mLFileAddr MLModelAddr, ) *MLEventNotif`

NewMLEventNotif instantiates a new MLEventNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMLEventNotifWithDefaults

`func NewMLEventNotifWithDefaults() *MLEventNotif`

NewMLEventNotifWithDefaults instantiates a new MLEventNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *MLEventNotif) GetEvent() NwdafEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *MLEventNotif) GetEventOk() (*NwdafEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *MLEventNotif) SetEvent(v NwdafEvent)`

SetEvent sets Event field to given value.


### GetNotifCorreId

`func (o *MLEventNotif) GetNotifCorreId() string`

GetNotifCorreId returns the NotifCorreId field if non-nil, zero value otherwise.

### GetNotifCorreIdOk

`func (o *MLEventNotif) GetNotifCorreIdOk() (*string, bool)`

GetNotifCorreIdOk returns a tuple with the NotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorreId

`func (o *MLEventNotif) SetNotifCorreId(v string)`

SetNotifCorreId sets NotifCorreId field to given value.

### HasNotifCorreId

`func (o *MLEventNotif) HasNotifCorreId() bool`

HasNotifCorreId returns a boolean if a field has been set.

### GetMLFileAddr

`func (o *MLEventNotif) GetMLFileAddr() MLModelAddr`

GetMLFileAddr returns the MLFileAddr field if non-nil, zero value otherwise.

### GetMLFileAddrOk

`func (o *MLEventNotif) GetMLFileAddrOk() (*MLModelAddr, bool)`

GetMLFileAddrOk returns a tuple with the MLFileAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLFileAddr

`func (o *MLEventNotif) SetMLFileAddr(v MLModelAddr)`

SetMLFileAddr sets MLFileAddr field to given value.


### GetValidityPeriod

`func (o *MLEventNotif) GetValidityPeriod() TimeWindow`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *MLEventNotif) GetValidityPeriodOk() (*TimeWindow, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *MLEventNotif) SetValidityPeriod(v TimeWindow)`

SetValidityPeriod sets ValidityPeriod field to given value.

### HasValidityPeriod

`func (o *MLEventNotif) HasValidityPeriod() bool`

HasValidityPeriod returns a boolean if a field has been set.

### GetSpatialValidity

`func (o *MLEventNotif) GetSpatialValidity() NetworkAreaInfo`

GetSpatialValidity returns the SpatialValidity field if non-nil, zero value otherwise.

### GetSpatialValidityOk

`func (o *MLEventNotif) GetSpatialValidityOk() (*NetworkAreaInfo, bool)`

GetSpatialValidityOk returns a tuple with the SpatialValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpatialValidity

`func (o *MLEventNotif) SetSpatialValidity(v NetworkAreaInfo)`

SetSpatialValidity sets SpatialValidity field to given value.

### HasSpatialValidity

`func (o *MLEventNotif) HasSpatialValidity() bool`

HasSpatialValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


